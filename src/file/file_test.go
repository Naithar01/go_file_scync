package file

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestNewFiles(t *testing.T) {
	testDir := "testdata"
	rootDepth := 1

	expectedFiles := []File{
		{
			DirectoryPath: testDir,
			FileName:      "file1.txt",
			FileSize:      0,
			FileModTime:   time.Time{},
			Depth:         rootDepth,
		},
		{
			DirectoryPath: testDir,
			FileName:      "file2.txt",
			FileSize:      0,
			FileModTime:   time.Time{},
			Depth:         rootDepth,
		},
	}

	files, err := NewFiles(testDir, rootDepth)
	if err != nil {
		t.Errorf("Error creating files: %v", err)
	}

	if !reflect.DeepEqual(files, expectedFiles) {
		t.Errorf("Expected %+v, but got %+v", expectedFiles, files)
	}
}

func TestParseDirectoryFiles(t *testing.T) {
	files := []File{
		{
			DirectoryPath: "dir1",
			FileName:      "file1.txt",
			FileSize:      100,
			FileModTime:   time.Now(),
			Depth:         1,
		},
		{
			DirectoryPath: "dir1",
			FileName:      "file2.txt",
			FileSize:      200,
			FileModTime:   time.Now(),
			Depth:         1,
		},
		{
			DirectoryPath: "dir2",
			FileName:      "file3.txt",
			FileSize:      150,
			FileModTime:   time.Now(),
			Depth:         1,
		},
	}

	expectedFileDir := map[string][]File{
		"dir1": {
			files[0],
			files[1],
		},
		"dir2": {
			files[2],
		},
	}

	result := ParseDirectoryFiles(files)

	if !reflect.DeepEqual(result, expectedFileDir) {
		t.Errorf("Expected %+v, but got %+v", expectedFileDir, result)
	}
}

func TestMain(m *testing.M) {
	testDir := "testdata"
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		os.Mkdir(testDir, os.ModePerm)
		os.Create(filepath.Join(testDir, "file1.txt"))
		os.Create(filepath.Join(testDir, "file2.txt"))
	}

	exitCode := m.Run()

	os.RemoveAll(testDir)

	os.Exit(exitCode)
}
