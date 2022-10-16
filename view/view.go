package view

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func RunUi() {
	app := app.New()
	win1 := app.NewWindow(fmt.Sprintf("gfile - %s", gfile_version))
	app.SetIcon(theme.FolderIcon())

	getFileSection()

	win1.SetContent(
		container.NewBorder(
			getToolbar(),
			nil,
			nil,
			nil,
			getFileSection(),
		),
	)

	win1.RequestFocus()
	win1.Resize(window_size)
	win1.Show()

	app.Run()
}
