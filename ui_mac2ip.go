package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func newMAC2IPTab(w fyne.Window, adapters []adapter) fyne.CanvasObject {
	// adapter select
	adapterNames := adapterNames(adapters)
	adapterEntry := widget.NewSelectEntry(adapterNames)
	if len(adapterNames) > 0 {
		adapterEntry.SetText(adapterNames[0])
	}

	// result box
	ipEntry := widget.NewEntry()
	var ipCopyButton *widget.Button
	ipCopyButton = widget.NewButton("Copy to clipboard", func() {
		w.Clipboard().SetContent(ipEntry.Text)
		ipCopyButton.Disable()
		ipCopyButton.SetText("Copied!")
	})
	ipResult := widget.NewVBox(
		widget.NewLabel("IP address:"),
		ipEntry,
		ipCopyButton,
	)
	ipResult.Hide()

	// address input and submit
	macEntry := newEnterEntry()
	macEntry.SetPlaceHolder("ex. 12:34:56:78:90:ab")
	var resolveButton *widget.Button
	resolveButton = widget.NewButton("Resolve", func() {
		resolveButton.Disable()
		resolveButton.SetText("Resolving...")
		defer func() {
			resolveButton.Enable()
			resolveButton.SetText("Resolve")
		}()
		ip, err := mac2ip(macEntry.Text, findAdapter(adapters, adapterEntry.Text))
		if err != nil {
			ipEntry.SetText("ERROR: " + err.Error())
		} else {
			ipEntry.SetText(ip)
		}
		ipCopyButton.SetText("Copy to clipboard")
		ipCopyButton.Enable()
		ipResult.Show()
	})
	macEntry.setOnEnter(resolveButton.OnTapped)
	resolveButton.Disable()
	macEntry.OnChanged = func(s string) {
		if len(s) > 0 && macPattern.MatchString(s) {
			resolveButton.Enable()
		} else {
			resolveButton.Disable()
		}
	}

	// layout
	return widget.NewVBox(
		widget.NewLabel("Target MAC address:"),
		macEntry,
		widget.NewLabel("Network adapter:"),
		adapterEntry,
		resolveButton,
		ipResult,
	)
}
