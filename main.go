package main

import (
	"regexp"

	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

var (
	macPattern = regexp.MustCompile("([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})")
	ipPattern  = regexp.MustCompile("(?:[0-9]{1,3}\\.){3}[0-9]{1,3}")
)

func main() {
	a := app.New()
	w := a.NewWindow("arping-gui")
	w.SetMainMenu(newMainMenu(a, w))
	adapters, err := adapters()
	if err != nil {
		dialog.NewError(err, w)
	}
	w.SetContent(
		widget.NewTabContainer(
			widget.NewTabItem("IP to MAC", newIP2MACTab(w, adapters)),
			widget.NewTabItem("MAC to IP", newMAC2IPTab(w, adapters)),
		),
	)
	w.ShowAndRun()
}
