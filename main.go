package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	App        fyne.App
	MainWindow fyne.Window
}

func main() {
	var myApp App

	myApp.App = app.NewWithID("com.goalodoro.aa")
	myApp.MainWindow = myApp.App.NewWindow("Goalodoro")
	myApp.MainWindow.Resize(fyne.NewSize(600, 500))

	myApp.makeUI()

	myApp.MainWindow.ShowAndRun()
}
