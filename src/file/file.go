package file

import (
	"errors"
	"go_file_sync/src/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 선택한 폴더에 있는 파일, 하위 폴더에 존재하는 파일들의 목록을 수집
func NewFiles(directoryPath string, rootDepth int) ([]models.File, error) {
	var files []models.File
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			dirPath := filepath.Dir(path)
			depth := strings.Count(dirPath, string(filepath.Separator)) - strings.Count(directoryPath, string(filepath.Separator))
			files = append(files, models.File{
				DirectoryPath: dirPath,
				FileName:      info.Name(),
				FileSize:      info.Size(),
				FileModTime:   info.ModTime(),
				Depth:         depth + rootDepth,
			})
		} else {
			dirPath := path
			depth := strings.Count(dirPath, string(filepath.Separator)) - strings.Count(directoryPath, string(filepath.Separator))
			files = append(files, models.File{
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
func ParseDirectoryFiles(filesInfo []models.File) map[string][]models.File {
	var fileDir = map[string][]models.File{}

	for _, v := range filesInfo {
		fileDir[v.DirectoryPath] = append(fileDir[v.DirectoryPath], v)
	}

	return fileDir
}

// 경로에 있는 파일을 읽고 문자열로 반환
func ReadFile(file_path string) ([]byte, error) {
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// 새로운 파일 생성 및 내용 작성
func WriteNewFile(filePath string, content string) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return errors.New(err.Error())
	}

	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
