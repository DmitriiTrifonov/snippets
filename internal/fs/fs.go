package fs

import (
	"log"
	"net/http"
	"path/filepath"

	"go.uber.org/zap"
)

// NeuteredFileSystem is used to provide static content to service
type NeuteredFileSystem struct {
	fs     http.FileSystem
	logger *zap.Logger
}

// NewNeuteredFS inits NeuteredFileSystem
func NewNeuteredFS(fs http.FileSystem, logger *zap.Logger) NeuteredFileSystem {
	return NeuteredFileSystem{
		fs:     fs,
		logger: logger,
	}
}

// Open implements http.FileSystem interface
func (nfs NeuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err = nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
