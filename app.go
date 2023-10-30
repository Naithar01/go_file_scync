package main

import (
	"context"
	"go_file_sync/src/file"
	"go_file_sync/src/global"
	"go_file_sync/src/logs"
	"go_file_sync/src/models"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context

	Files []models.File
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "프로그램 종료",
		Message: "프로그램을 종료하시겠습니까?",
	})

	if err != nil {
		return false
	}

	if dialog == "Yes" {
		logs.PrintMsgLog("프로그램 종료")
		return false
	}

	return true
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

// Parse Directory
func (a *App) ParseRootPath(directory_path string) models.ResponseFileStruct {
	files, err := file.NewFiles(directory_path, 0)
	if err != nil {
		logs.CustomErrorDialog(a.ctx, err.Error())
		return models.ResponseFileStruct{
			Root_path: "",
			Files:     nil,
		}
	}

	// 파일 정보와, 선택한 폴더의 경로를 저장
	a.Files = files
	if err := global.SetRootPath(directory_path); err != nil {
		logs.CustomErrorDialog(a.ctx, err.Error())
		return models.ResponseFileStruct{
			Root_path: "",
			Files:     nil,
		}
	}

	fileDir := file.ParseDirectoryFiles(a.Files)
	return models.ResponseFileStruct{
		Root_path: directory_path,
		Files:     fileDir,
	}
}

// root path, files
func (a *App) OpenDirectory() models.ResponseFileStruct {
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
		return models.ResponseFileStruct{
			Root_path: "",
			Files:     nil,
		}
	}

	return a.ParseRootPath(directory_path)
}

// Custom Error Dialog
func (a *App) CustomErrorDialog(errorMessage string) {
	logs.CustomErrorDialog(a.ctx, errorMessage)
}

// Custom Info Dialog
func (a *App) CustomInfoDialog(InfoMessage string) {
	logs.CustomInfoDialog(a.ctx, InfoMessage)
}

// Get environment variable Data
func (a *App) GetRootPath() string {
	return global.GetRootPath()
}
