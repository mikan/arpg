package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func newMainMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	homepage, err := url.Parse("https://github.com/mikan/arping-gui")
	if err != nil {
		panic(fmt.Sprintf("oops! failed to parse homepage URL: %v", err))
	}
	return fyne.NewMainMenu(
		fyne.NewMenu("File", fyne.NewMenuItem("Quit", func() { a.Quit() })),
		fyne.NewMenu("Help", fyne.NewMenuItem("About...", func() {
			dialog.NewCustom("About", "OK", widget.NewVBox(
				widget.NewHyperlinkWithStyle("arping-gui", homepage, fyne.TextAlignCenter, boldText),
				widget.NewLabelWithStyle("A simple arping-link tool", fyne.TextAlignCenter, italicText),
				widget.NewLabelWithStyle("Licensed under the", fyne.TextAlignCenter, normalText),
				widget.NewLabelWithStyle("BSD 3-Clause", fyne.TextAlignCenter, normalText),
			), w)
		})),
	)
}
