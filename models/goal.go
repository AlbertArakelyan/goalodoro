package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Goal struct {
	Name        string        `json:"name"`
	TargetHours float64       `json:"target_hours"`
	Spent       time.Duration `json:"spent"`
}

var Goals []Goal
var goalFile = "goals.json"

func LoadGoals() {
	file, err := os.ReadFile(goalFile)
	if err == nil {
		_ = json.Unmarshal(file, &Goals)
	}
}

func SaveGoals() {
	data, _ := json.MarshalIndent(Goals, "", "  ")
	_ = os.WriteFile(goalFile, data, 0644)
}

func ExportGoalsToJSON(mainWondow fyne.Window) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, err error) {
		if uri == nil {
			fmt.Println("Save operation was canceled.")
			return
		}

		if err != nil {
			dialog.ShowError(err, mainWondow)
			return
		}

		if uri.URI().Path() == "" {
			dialog.ShowError(err, mainWondow)
			return
		}

		data, _ := json.MarshalIndent(Goals, "", "  ")
		_, err = uri.Write(data)

		if err != nil {
			dialog.ShowError(err, mainWondow)
			return
		}

		_ = uri.Close()
	}, mainWondow)
}
