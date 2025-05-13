package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewMainLayout(mainWindow fyne.Window) func(page *fyne.Container) *fyne.Container {
	return func(page *fyne.Container) *fyne.Container {
		icon, _ := fyne.LoadResourceFromPath("Icon.png")
		logo := widget.NewIcon(icon)

		sideBarContent := container.NewVBox(
			logo,
			widget.NewButtonWithIcon("Book", theme.CalendarIcon(), nil),
			widget.NewButtonWithIcon("Timer", theme.ComputerIcon(), nil),
			widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), nil),
		)

		sidebar := container.NewPadded(sideBarContent)

		return container.NewBorder(nil, nil, sidebar, nil, page)
	}
}
