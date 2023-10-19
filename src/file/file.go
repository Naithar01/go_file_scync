package file

import (
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

// 선택한 폴더에 있는 파일, 하위 폴더에 존재하는 파일들의 목록을 수집
func NewFiles(directoryPath string, rootDepth int) ([]File, error) {
	var files []File
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
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
		} else {
			dirPath := path
			depth := strings.Count(dirPath, string(filepath.Separator)) - strings.Count(directoryPath, string(filepath.Separator))
			files = append(files, File{
				DirectoryPath: dirPath,
				FileName:      "",
				FileSize:      0,
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

// 폴더 이름별로 파일을 변수화
func ParseDirectoryFiles(filesInfo []File) map[string][]File {
	var fileDir = map[string][]File{}

	for _, v := range filesInfo {
		fileDir[v.DirectoryPath] = append(fileDir[v.DirectoryPath], v)
	}

	return fileDir
}
