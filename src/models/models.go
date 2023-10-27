package models

import (
	"time"
)

type ResponseFileStruct struct {
	Root_path string            `json:"root_path"`
	Files     map[string][]File `json:"files"`
	File      File              `json:"file"`
}

type File struct {
	DirectoryPath string    `json:"directorypath"`
	FileName      string    `json:"filename"`
	FileSize      int64     `json:"filesize"`
	FileModTime   time.Time `json:"filemodtime"`
	Depth         int       `json:"depth"`
}

type Message struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

type ServerShutdownMessage struct {
	Type    string `json:"type"`
	Content error  `json:"content"`
}

type DirectoryContent struct {
	Type    string             `json:"type"`
	Content ResponseFileStruct `json:"content"`
}

type TCPAutoConnect struct {
	Type    string `json:"type"`
	Content int    `json:"content"`
}
