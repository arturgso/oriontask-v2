package main

import (
	"embed"
	"oriontask-v2/database"
	"oriontask-v2/dharmas"
	"oriontask-v2/tasks"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	db, err := database.Connect("oriontask")
	if err != nil {
		panic(err)
	}

	dharmaRepo := dharmas.NewDharmaRepository(db)
	dharma := dharmas.NewDharmaService(dharmaRepo)

	tasksRepo := tasks.NewTaskRepository(db)
	taskService := tasks.NewTaskService(tasksRepo)

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "Oriontask",
		Width:     1024,
		Height:    768,
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			dharma,
			taskService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
