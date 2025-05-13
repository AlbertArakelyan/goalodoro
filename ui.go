package main

import (
	"github.com/AlbertArakelyan/goalodoro/components/layouts"
	"github.com/AlbertArakelyan/goalodoro/pages/home"
)

func (app *App) makeUI() {
	mainLayout := layouts.NewMainLayout(app.MainWindow)

	homePage := home.Home(app.MainWindow)

	app.MainWindow.SetContent(mainLayout(homePage))
}
