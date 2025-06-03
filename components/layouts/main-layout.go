package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainLayout(mainWindow fyne.Window) func(page *fyne.Container) *fyne.Container {
	return func(page *fyne.Container) *fyne.Container {
		// icon, _ := fyne.LoadResourceFromPath("Icon.png")
		// logo := widget.NewIcon(icon)

		// booksBtn := widget.NewButton("📚", func() {
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
		// 	dialog.ShowInformation("Growth Books", msg, mainWindow)
		// })

		// exportGoalsToJSONBtn := widget.NewButton("📝", func() {
		// 	models.ExportGoalsToJSON(mainWindow)
		// })

		// sideBarContent := container.NewVBox(
		// 	container.NewPadded(logo),
		// 	booksBtn,
		// 	exportGoalsToJSONBtn,
		// )

		// sidebar := sideBarContent

		return container.NewBorder(nil, nil, nil, nil, page)
	}
}
