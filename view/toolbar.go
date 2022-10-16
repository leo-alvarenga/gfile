package view

import (
	"log"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func getToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), func() {}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
}
