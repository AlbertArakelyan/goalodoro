package home

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/AlbertArakelyan/goalodoro/models"
)

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func Home(mainWondow fyne.Window) *fyne.Container {
	models.LoadGoals()

	selectedGoalIndex := -1
	var ticker *time.Ticker
	var stopChan chan struct{}

	goalList := widget.NewList(
		func() int { return len(models.Goals) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Goal")
		},
		func(i int, o fyne.CanvasObject) {
			g := models.Goals[i]
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
				models.Goals = append(models.Goals, models.Goal{Name: name.Text, TargetHours: h})
				models.SaveGoals()
				goalList.Refresh()
			}
		}, mainWondow)
		dlg.Resize(fyne.NewSize(300, 200))
		dlg.Show()
	})

	startBtn := widget.NewButton("▶️ Start", func() {
		if selectedGoalIndex == -1 {
			dialog.ShowInformation("No Goal Selected", "Please select a goal to start.", mainWondow)
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
					models.Goals[selectedGoalIndex].Spent += time.Second
					models.SaveGoals()
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
			dialog.ShowInformation("No Goal Selected", "Please select a goal to delete.", mainWondow)
			return
		}
		models.Goals = append(models.Goals[:selectedGoalIndex], models.Goals[selectedGoalIndex+1:]...)
		// TODO add error handling when deleting a non existing index item
		models.SaveGoals()
		goalList.Refresh()
	})

	// exportGoalsToJSONBtn := widget.NewButton("📝 Export Goals", func() {
	// 	exportGoalsToJSON(mainWondow)
	// })

	// booksBtn := widget.NewButton("📚 Growth Reads", func() {
	// 	books := []string{
	// 		// "The Power of Now – Eckhart Tolle",
	// 		// "Can’t Hurt Me – David Goggins",
	// 		"📘 Atomic Habits — James Clear",
	// 		"📕 Deep Work — Cal Newport",
	// 		"📗 The One Thing — Gary Keller",
	// 		"📙 Grit — Angela Duckworth",
	// 		"📔 Make Time — Jake Knapp",
	// 	}
	// 	msg := "Recommended Reads:\n\n" + fmt.Sprint("- "+books[0])
	// 	for _, b := range books[1:] {
	// 		msg += "\n- " + b
	// 	}
	// 	dialog.ShowInformation("Growth Books", msg, mainWondow)
	// })

	controls := container.NewHBox(
		addGoalBtn,
		startBtn,
		stopBtn,
		deleteBtn,
	)

	homePageContent := container.NewBorder(controls, nil, nil, nil, goalList)

	return homePageContent
}
