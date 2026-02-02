package storage

import (
	"os"
	"testing"
)

func TestFileStorageSetup(t *testing.T) {
	dataPath := ".data"

	_ = os.RemoveAll(dataPath)
	t.Cleanup(func() {
		_ = os.RemoveAll(dataPath)
	})

	storage := NewFileStorage(dataPath)
	if err := storage.Setup(); err != nil {
		t.Fatalf("Setup() failed: %v", err)
	}
}
