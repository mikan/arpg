package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func newMainMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	homepage, err := url.Parse("https://github.com/mikan/arpg")
	if err != nil {
		panic(fmt.Sprintf("oops! failed to parse homepage URL: %v", err))
	}
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Save log...", func() {
				dialog.NewFileSave(func(file fyne.URIWriteCloser, err error) {
					if file != nil {
						if _, err = file.Write([]byte(logContent)); err != nil {
							dialog.ShowError(err, w)
						}
						if err := file.Close(); err != nil {
							dialog.ShowError(err, w)
						}
					}
				}, w).Show()
			}),
			fyne.NewMenuItem("Quit", func() { a.Quit() }),
		),
		fyne.NewMenu("Help", fyne.NewMenuItem("About...", func() {
			dialog.NewCustom("About", "OK", container.NewVBox(
				widget.NewHyperlinkWithStyle("ARPG", homepage, fyne.TextAlignCenter, boldText),
				widget.NewLabelWithStyle("A simple ARP support tool", fyne.TextAlignCenter, italicText),
				widget.NewLabelWithStyle("Licensed under the", fyne.TextAlignCenter, normalText),
				widget.NewLabelWithStyle("BSD 3-Clause", fyne.TextAlignCenter, normalText),
			), w).Show()
		})),
	)
}
