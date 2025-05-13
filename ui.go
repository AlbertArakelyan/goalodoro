package main

import (
	"github.com/AlbertArakelyan/goalodoro/components/layouts"
	"github.com/AlbertArakelyan/goalodoro/pages/home"
)

func (app *App) makeUI() {
	homePage := home.Home(app.MainWindow)

	app.MainWindow.SetContent(layouts.MainLayout(homePage))
}
