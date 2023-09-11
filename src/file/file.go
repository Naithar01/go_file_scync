package file

import (
	"os"
	"path/filepath"
	"time"
)

type File struct {
	DirectoryPath string    `json:"directorypath"` // 추가된 필드
	FileName      string    `json:"filename"`
	FileSize      int64     `json:"filesize"`
	FileModTime   time.Time `json:"filemodtime"`
}

func NewFiles(directoryPath string) ([]File, error) {
	var files []File
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			dirPath := filepath.Dir(path)
			files = append(files, File{
				DirectoryPath: dirPath,
				FileName:      info.Name(),
				FileSize:      info.Size(),
				FileModTime:   info.ModTime(),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
