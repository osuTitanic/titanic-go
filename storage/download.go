package storage

import (
	"io"
	"net/http"
)

var DefaultAssetUrls = map[string]string{
	// Default avatars
	"avatars/unknown": "https://github.com/osuTitanic/titanic/blob/main/.github/images/avatars/unknown.jpg?raw=true",
	"avatars/1":       "https://github.com/osuTitanic/titanic/blob/main/.github/images/avatars/banchobot.jpg?raw=true",
}

func GetDownloadStream(key string) (io.ReadCloser, error) {
	url, exists := DefaultAssetUrls[key]
	if !exists {
		return nil, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
