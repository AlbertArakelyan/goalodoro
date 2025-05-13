package main

import "github.com/AlbertArakelyan/goalodoro/pages/home"

func (app *App) makeUI() {
	homePage := home.Home(app.MainWindow)

	app.MainWindow.SetContent(homePage)
}
