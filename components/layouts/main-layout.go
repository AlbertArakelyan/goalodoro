package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainLayout(page *fyne.Container) *fyne.Container {
	icon, _ := fyne.LoadResourceFromPath("Icon.png")
	logo := widget.NewIcon(icon)

	sideBarContent := container.NewVBox(
		logo,
	)

	sidebar := container.NewPadded(sideBarContent)

	return container.NewBorder(nil, nil, sidebar, nil, page)
}
