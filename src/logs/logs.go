package logs

import (
	"io"
	"log"
	"os"
)

var logFile *os.File

func LoadLogFile() {
	loadLogFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(loadLogFile, os.Stdout)
	log.SetOutput(multiWriter)

	logFile = loadLogFile
}

func CloseLogFile() {
	logFile.Close()
}

func PrintMsgLog(msg string) {
	log.Print(msg)
}
