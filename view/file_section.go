package view

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/leo-alvarenga/gfile/ng"
	"github.com/leo-alvarenga/gfile/ng/dirtree"
)

func getFileSection() *container.Scroll {
	section := container.NewGridWrap(item_size)
	scroll := container.NewScroll(section)
	populateWithFiles(section, scroll)

	return scroll
}

func forceRefreshOnFileSection(wrapper *fyne.Container, scroll *container.Scroll) {
	wrapper.RemoveAll()
	populateWithFiles(wrapper, scroll)

	wrapper.Refresh()

	scroll.ScrollToTop()
	scroll.Refresh()
}

func populateWithFiles(wrapper *fyne.Container, scroll *container.Scroll) {
	nodes := ng.LOCATION_CURSOR.GetInfo().GetChildren()

	if !ng.LOCATION_CURSOR.IsThisRoot() {
		wrapper.Add(
			widget.NewButtonWithIcon(
				"Go back",
				theme.NavigateBackIcon(),
				func() {
					ng.LOCATION_CURSOR.GoBack()

					forceRefreshOnFileSection(wrapper, scroll)
				},
			),
		)
	}

	for _, node := range nodes {
		if show_hidden_files || node.GetName()[0] != '.' {
			wrapper.Add(newDirButton(node, wrapper, scroll))
		}
	}
}

func newDirButton(node *dirtree.DirNode, wrapper *fyne.Container, scroll *container.Scroll) *widget.Button {
	icon := theme.FileIcon()

	if !node.IsDir() {
		parts := strings.Split(node.GetName(), ".")

		if len(parts) > 1 {
			switch strings.ToLower(parts[1]) {
			case "jpg":
				icon = theme.FileImageIcon()
			case "jpge":
				icon = theme.FileImageIcon()
			case "png":
				icon = theme.FileImageIcon()
			case "gif":
				icon = theme.FileImageIcon()

			case "mp4":
				icon = theme.FileVideoIcon()
			case "avi":
				icon = theme.FileVideoIcon()

			case "wav":
				icon = theme.FileAudioIcon()
			case "mp3":
				icon = theme.FileAudioIcon()
			case "aac":
				icon = theme.FileAudioIcon()

			default:
				icon = theme.FileTextIcon()
			}
		}
	} else {
		icon = theme.FolderIcon()
	}

	name := node.GetName()
	if len(name)-1 > max_filename_len {
		name = name[:max_filename_len] + "..."
	}

	return widget.NewButtonWithIcon(
		name,
		icon,
		func() {
			ng.LOCATION_CURSOR.EnterChild(node.GetName())

			forceRefreshOnFileSection(wrapper, scroll)
		},
	)
}
