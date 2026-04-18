package templates

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet/v6"
	"github.com/osuTitanic/titanic-go/internal/config"
)

type Engine struct {
	Set *jet.Set
}

func NewEngine(cfg *config.Config) (*Engine, error) {
	root, err := resolveTemplateRoot()
	if err != nil {
		return nil, err
	}

	set := jet.NewSet(
		jet.NewOSFileSystemLoader(root),
		jet.DevelopmentMode(cfg.Reload),
	)
	return &Engine{Set: set}, nil
}

func (e *Engine) Render(name string, data any) ([]byte, error) {
	view, err := e.Set.GetTemplate(name)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := view.Execute(&buf, nil, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TODO: Use embedded filesystems, this is way to hacky
func resolveTemplateRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("templates: get working directory: %w", err)
	}

	// Find folder with go.mod & append template path to it
	for current := wd; ; current = filepath.Dir(current) {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			root := filepath.Join(current, "services", "stern", "web", "template")

			if _, err := os.Stat(root); err != nil {
				return "", fmt.Errorf("templates: template root %q: %w", root, err)
			}
			return root, nil
		}

		// If we reach the root of the filesystem, stop searching
		if filepath.Dir(current) == current {
			break
		}
	}

	return "", fmt.Errorf("templates: could not locate repo root from %q", wd)
}
