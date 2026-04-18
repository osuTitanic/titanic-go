package stern

import (
	"embed"
	"io/fs"
	"path"
	"strings"
)

//go:embed template
var Templates embed.FS

//go:embed static
var Static embed.FS

func StaticFS(assetPath string) (fs.FS, error) {
	cleaned := strings.Trim(assetPath, "/")
	if cleaned == "" {
		return fs.Sub(Static, "static")
	}

	target := path.Join("static", cleaned)
	info, err := fs.Stat(Static, target)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return fs.Sub(Static, target)
	}

	return fs.Sub(Static, "static")
}
