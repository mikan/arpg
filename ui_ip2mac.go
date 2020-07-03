package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func newIP2MACTab(w fyne.Window, adapters []adapter) fyne.CanvasObject {
	// adapter select
	adapterNames := adapterNames(adapters)
	adapterEntry := widget.NewSelectEntry(adapterNames)
	if len(adapterNames) > 0 {
		adapterEntry.SetText(adapterNames[0])
	}

	// result box
	macEntry := widget.NewEntry()
	var macCopyButton *widget.Button
	macCopyButton = widget.NewButton("Copy to clipboard", func() {
		w.Clipboard().SetContent(macEntry.Text)
		macCopyButton.Disable()
		macCopyButton.SetText("Copied!")
	})
	macResult := widget.NewVBox(
		widget.NewLabel("MAC address:"),
		macEntry,
		macCopyButton,
	)
	macResult.Hide()

	// address input and submit
	ipEntry := newEnterEntry()
	ipEntry.SetPlaceHolder("ex. 192.168.1.1")
	var resolveButton *widget.Button
	resolveButton = widget.NewButton("Resolve", func() {
		resolveButton.Disable()
		resolveButton.SetText("Resolving...")
		defer func() {
			resolveButton.Enable()
			resolveButton.SetText("Resolve")
		}()
		mac, err := ip2mac(ipEntry.Text, findAdapter(adapters, adapterEntry.Text))
		if err != nil {
			macEntry.SetText("ERROR: " + err.Error())
		} else {
			macEntry.SetText(mac)
		}
		macCopyButton.SetText("Copy to clipboard")
		macCopyButton.Enable()
		macResult.Show()
	})
	ipEntry.setOnEnter(resolveButton.OnTapped)
	resolveButton.Disable()
	ipEntry.OnChanged = func(s string) {
		if len(s) > 0 && ipPattern.MatchString(s) {
			resolveButton.Enable()
		} else {
			resolveButton.Disable()
		}
	}

	// layout
	return widget.NewVBox(
		widget.NewLabel("Target IP address:"),
		ipEntry,
		widget.NewLabel("Network adapter:"),
		adapterEntry,
		resolveButton,
		macResult,
	)
}
