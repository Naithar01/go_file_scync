package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	DirectoryPath string    `json:"directorypath"`
	FileName      string    `json:"filename"`
	FileSize      int64     `json:"filesize"`
	FileModTime   time.Time `json:"filemodtime"`
	Depth         int       `json:"depth"`
}

func NewFiles(directoryPath string, rootDepth int) ([]File, error) {
	var files []File
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			dirPath := filepath.Dir(path)
			depth := strings.Count(dirPath, string(filepath.Separator)) - strings.Count(directoryPath, string(filepath.Separator))
			files = append(files, File{
				DirectoryPath: dirPath,
				FileName:      info.Name(),
				FileSize:      info.Size(),
				FileModTime:   info.ModTime(),
				Depth:         depth + rootDepth,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ParseDirectoryFiles(filesInfo []File) map[string][]File {
	var fileDir = map[string][]File{}

	for _, v := range filesInfo {
		fileDir[v.DirectoryPath] = append(fileDir[v.DirectoryPath], v)
	}

	return fileDir
}
