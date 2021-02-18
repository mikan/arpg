package main

import (
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

var (
	macPattern = regexp.MustCompile("([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})")
	ipPattern  = regexp.MustCompile("(?:[0-9]{1,3}\\.){3}[0-9]{1,3}")
)

func main() {
	overwriteFyneFont()
	a := app.New()
	w := a.NewWindow("ARPG")
	w.SetMainMenu(newMainMenu(a, w))
	w.Resize(fyne.NewSize(350, 500))
	adapters, err := adapters()
	if err != nil {
		dialog.NewError(err, w).Show()
	}
	w.SetContent(
		container.NewAppTabs(
			container.NewTabItem("IP to MAC", newIP2MACTab(w, adapters)),
			container.NewTabItem("MAC to IP", newMAC2IPTab(w, adapters)),
		),
	)
	w.ShowAndRun()
}
