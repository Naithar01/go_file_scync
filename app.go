package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
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
			menu.Text("Open Directory", keys.CmdOrCtrl("o"), func(cbdata *menu.CallbackData) {
				directory_path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
					DefaultDirectory:           "",
					DefaultFilename:            "",
					Title:                      "Selct Directory",
					Filters:                    nil,
					ShowHiddenFiles:            false,
					CanCreateDirectories:       false,
					ResolvesAliases:            false,
					TreatPackagesAsDirectories: false,
				})
				if err != nil {
					return
				}

				// 폴더 선택을 하지 않았을 때, 그냥 Dialog를 닫았을 때 에러 메시지 출력하게 수정, 밑 코드 실행하지 않게

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

				runtime.BrowserOpenURL(a.ctx, directory_path)
			}),
			menu.Separator(), // <br />
			menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
				runtime.Quit(a.ctx)
			}),
		)),
	)
}
