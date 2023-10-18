package main

import (
	"context"
	"go_file_sync/src/file"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context

	Files []file.File
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) applicationMenu() *menu.Menu {
	return menu.NewMenuFromItems(
		menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Separator(),
			menu.Text("종료", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
				runtime.Quit(a.ctx)
			}),
		)),
	)
}

type ResponseFileStruct struct {
	Root_path string                 `json:"root_path"`
	Files     map[string][]file.File `json:"files"`
	File      file.File              `json:"file"`
}

// root path, files
func (a *App) OpenDirectory() ResponseFileStruct {
	directory_path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           "",
		DefaultFilename:            "",
		Title:                      "Select Directory",
		Filters:                    nil,
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: true,
	})
	if err != nil {
		return ResponseFileStruct{
			Root_path: "",
			Files:     nil,
		}
	}

	files, err := file.NewFiles(directory_path, 0)
	if err != nil {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Error",
			Message:       "Can't Find Directory",
			Buttons:       nil,
			DefaultButton: "",
			CancelButton:  "",
		})
		return ResponseFileStruct{
			Root_path: "",
			Files:     nil,
		}
	}

	a.Files = files
	fileDir := file.ParseDirectoryFiles(a.Files)
	return ResponseFileStruct{
		Root_path: directory_path,
		Files:     fileDir,
	}
}

// Custom Error Dialog
func (a *App) CustomErrorDialog(errorMessage string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.ErrorDialog,
		Title:         "Error",
		Message:       errorMessage,
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
	})
}

// Custom Info Dialog
func (a *App) CustomInfoDialog(InfoMessage string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         "Info",
		Message:       InfoMessage,
		Buttons:       nil,
		DefaultButton: "",
		CancelButton:  "",
	})
}
