package global

import (
	"os"
)

func SetRootPath(root_path string) error {
	err := os.Setenv("RootPath", root_path)
	if err != nil {
		return err
	}

	return nil
}

func GetRootPath() string {
	root_path_string := os.Getenv("RootPath")
	return root_path_string
}
