package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var goals []Goal
var goalFile = "goals.json"

func loadGoals() {
	file, err := os.ReadFile(goalFile)
	if err == nil {
		_ = json.Unmarshal(file, &goals)
	}
}

func saveGoals() {
	data, _ := json.MarshalIndent(goals, "", "  ")
	_ = os.WriteFile(goalFile, data, 0644)
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func (app *App) makeUI() {
	loadGoals()

	selectedGoalIndex := -1
	var ticker *time.Ticker
	var stopChan chan struct{}

	goalList := widget.NewList(
		func() int { return len(goals) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Goal")
		},
		func(i int, o fyne.CanvasObject) {
			g := goals[i]
			spentStr := formatDuration(g.Spent)
			target := time.Duration(g.TargetHours * float64(time.Hour))
			targetStr := formatDuration(target)
			o.(*widget.Label).SetText(fmt.Sprintf("%s - %s / %s", g.Name, spentStr, targetStr))
		},
	)

	goalList.OnSelected = func(id widget.ListItemID) {
		selectedGoalIndex = id
	}

	addGoalBtn := widget.NewButton("➕ Add Goal", func() {
		name := widget.NewEntry()
		hours := widget.NewEntry()
		hours.SetPlaceHolder("e.g. 100")
		dlg := dialog.NewForm("Add Goal", "Add", "Cancel", []*widget.FormItem{
			{Text: "Goal Name", Widget: name},
			{Text: "Target Hours", Widget: hours},
		}, func(ok bool) {
			if ok {
				var h float64
				fmt.Sscanf(hours.Text, "%f", &h)
				goals = append(goals, Goal{Name: name.Text, TargetHours: h})
				saveGoals()
				goalList.Refresh()
			}
		}, app.MainWindow)
		dlg.Resize(fyne.NewSize(300, 200))
		dlg.Show()
	})

	startBtn := widget.NewButton("▶️ Start", func() {
		if selectedGoalIndex == -1 {
			dialog.ShowInformation("No Goal Selected", "Please select a goal to start.", app.MainWindow)
			return
		}
		if ticker != nil {
			return
		}
		stopChan = make(chan struct{})
		ticker = time.NewTicker(time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					goals[selectedGoalIndex].Spent += time.Second
					saveGoals()
					goalList.Refresh()
				case <-stopChan:
					ticker.Stop()
					ticker = nil
					return
				}
			}
		}()
	})

	stopBtn := widget.NewButton("⏹ Stop", func() {
		if ticker != nil && stopChan != nil {
			close(stopChan)
			goalList.Refresh()
		}
	})

	deleteBtn := widget.NewButton("❌ Delete", func() {
		if selectedGoalIndex == -1 {
			dialog.ShowInformation("No Goal Selected", "Please select a goal to delete.", app.MainWindow)
			return
		}
		goals = append(goals[:selectedGoalIndex], goals[selectedGoalIndex+1:]...)
		saveGoals()
		goalList.Refresh()
	})

	icon, _ := fyne.LoadResourceFromPath("Icon.png")
	logo := widget.NewIcon(icon)
	logo.Resize(fyne.NewSize(60, 60))

	controls := container.NewHBox(logo, addGoalBtn, startBtn, stopBtn, deleteBtn)
	app.MainWindow.SetContent(container.NewBorder(controls, nil, nil, nil, goalList))
}
