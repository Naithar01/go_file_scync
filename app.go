package main

import (
	"context"
	"fmt"
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
			menu.Text("Setting Directory", keys.CmdOrCtrl("o"), func(cbdata *menu.CallbackData) {
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
					return
				}

				files, err := file.NewFiles(directory_path)
				if err != nil {
					runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
						Type:          runtime.ErrorDialog,
						Title:         "Error",
						Message:       "Can't Find Directory",
						Buttons:       nil,
						DefaultButton: "",
						CancelButton:  "",
					})
					return
				}

				a.Files = files

				_, err = runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
					Type:          "info",
					Title:         "Selected Directory",
					Message:       directory_path,
					Buttons:       nil,
					DefaultButton: "",
					CancelButton:  "",
				})
				if err != nil {
					return
				}

				// runtime.BrowserOpenURL(a.ctx, directory_path)
			}),
			menu.Separator(), // <br />
			menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
				runtime.Quit(a.ctx)
			}),
		)),
	)
}

type ResponseFileStruct struct {
	Root_path string      `json:"root_path"`
	Files     []file.File `json:"files"`
}

// root path, files
func (a *App) ResponseFileData() ResponseFileStruct {
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

	files, err := file.NewFiles(directory_path)
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

	fmt.Println(files)

	for _, v := range files {
		fmt.Println(v.DirectoryPath, ">", v.FileName)
	}

	a.Files = files
	return ResponseFileStruct{
		Root_path: directory_path,
		Files:     a.Files,
	}
}
