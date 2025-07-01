package main

import (
	"embed"
	"context"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"flacer/cmd/libmanager"
	//"flacer/cmd/player"
)

//go:embed all:frontend/dist
var assets embed.FS
func main() {
	app := NewApp()
	service := libmanager.NewService()

	err := wails.Run(&options.App{
		Title:  "flacer",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},

		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			service.SetContext(ctx)
		},

		Bind: []interface{}{
			app, service,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

