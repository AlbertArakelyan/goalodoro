package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	App        fyne.App
	MainWindow fyne.Window
}

type Goal struct {
	Name        string        `json:"name"`
	TargetHours float64       `json:"target_hours"`
	Spent       time.Duration `json:"spent"`
}

func main() {
	var myApp App

	myApp.App = app.New()
	myApp.MainWindow = myApp.App.NewWindow("Goalodoro")
	myApp.MainWindow.Resize(fyne.NewSize(600, 500))

	myApp.makeUI()

	myApp.MainWindow.ShowAndRun()
}
