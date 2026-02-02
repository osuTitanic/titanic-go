package storage

import "io"

// Storage defines the interface for a storage backend
type Storage interface {
	Setup() error
	Save(key string, bucket string, data []byte) error
	SaveStream(key string, bucket string, stream io.Reader) error
	Read(key string, bucket string) ([]byte, error)
	ReadStream(key string, bucket string) (io.ReadSeekCloser, error)
	Remove(key string, bucket string) error
	Exists(key string, bucket string) bool
}

var RequiredDirectories = []string{
	"audio",
	"avatars",
	"beatmaps",
	"osz",
	"osz2",
	"release",
	"replays",
	"screenshots",
	"thumbnails",
	"logs",
}
