package logs

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 23.10.20 Dev Mode

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

// Custom Error Dialog
func CustomErrorDialog(c context.Context, errorMessage string) {
	runtime.MessageDialog(c, runtime.MessageDialogOptions{
		Type:          runtime.ErrorDialog,
		Title:         "Error",
		Message:       errorMessage,
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
	})
}

// Custom Info Dialog
func CustomInfoDialog(c context.Context, InfoMessage string) {
	runtime.MessageDialog(c, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Info",
		Message:       InfoMessage,
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
	})
}
