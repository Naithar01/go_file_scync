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
	Duplication   bool      `json:"duplication"`
	Latest        int       `json:"latest"`
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

type AutoConnect struct {
	Type    string             `json:"type"`
	Content AutoConnectContent `json:"content"`
}

type AutoConnectContent struct {
	IP   string `json:"ip"`
	PORT int    `json:"port"`
}

type ReadFile struct {
	FileName string `json:"filename"`
	FileData []byte `json:"filedata"`
}

type FileData struct {
	Type    string   `json:"type"`
	Content ReadFile `json:"content"`
}
