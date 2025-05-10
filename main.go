package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Goal struct {
	Name        string        `json:"name"`
	TargetHours float64       `json:"target_hours"`
	Spent       time.Duration `json:"spent"`
}

var goals []Goal
var goalFile = "goals.json"
var timerRunning = false
var startTime time.Time
var selectedGoalIndex = -1

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

func main() {
	a := app.New()
	w := a.NewWindow("Goalodoro")
	w.Resize(fyne.NewSize(600, 500))

	loadGoals()

	goalList := widget.NewList(
		func() int { return len(goals) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Goal")
		},
		func(i int, o fyne.CanvasObject) {
			g := goals[i]
			percent := float64(g.Spent.Hours()) / g.TargetHours * 100
			o.(*widget.Label).SetText(fmt.Sprintf("%s - %.1f%% (%.1fh/%.1fh)", g.Name, percent, g.Spent.Hours(), g.TargetHours))
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
		}, w)
		dlg.Show()
	})

	startBtn := widget.NewButton("▶ Start", func() {
		if selectedGoalIndex == -1 {
			dialog.ShowInformation("No Goal Selected", "Please select a goal to start timing.", w)
			return
		}
		if !timerRunning {
			startTime = time.Now()
			timerRunning = true
			dialog.ShowInformation("Started", fmt.Sprintf("Started tracking time for '%s'", goals[selectedGoalIndex].Name), w)
		}
	})

	stopBtn := widget.NewButton("■ Stop", func() {
		if timerRunning {
			elapsed := time.Since(startTime)
			goals[selectedGoalIndex].Spent += elapsed
			saveGoals()
			goalList.Refresh()
			timerRunning = false
			dialog.ShowInformation("Stopped", fmt.Sprintf("Stopped. %v added to '%s'", elapsed.Truncate(time.Second), goals[selectedGoalIndex].Name), w)
		} else {
			dialog.ShowInformation("Not Running", "No timer is currently running.", w)
		}
	})

	w.SetContent(container.NewBorder(
		container.NewHBox(addGoalBtn, startBtn, stopBtn),
		nil, nil, nil,
		goalList,
	))

	w.ShowAndRun()
}
