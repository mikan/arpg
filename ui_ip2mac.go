package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newIP2MACTab(w fyne.Window, adapters []adapter) fyne.CanvasObject {
	// adapter select
	adapterNames := adapterNames(adapters)
	adapterEntry := widget.NewSelectEntry(adapterNames)
	if len(adapterNames) > 0 {
		adapterEntry.SetText(adapterNames[0])
	}
	adapterEntry.Disable()
	adapterAutoCheck := widget.NewCheck("Auto", func(checked bool) {
		if checked {
			adapterEntry.Disable()
		} else {
			adapterEntry.Enable()
		}
	})
	adapterAutoCheck.SetChecked(true)

	// log pane
	logPaneLabel.Hide()
	logPane.Disable()
	logPane.Hide()

	// result box
	macEntry := widget.NewEntry()
	var macCopyButton *widget.Button
	macCopyButton = widget.NewButton("Copy to clipboard", func() {
		w.Clipboard().SetContent(macEntry.Text)
		macCopyButton.Disable()
		macCopyButton.SetText("Copied!")
		logContent = fmt.Sprintf("[%s]\nCP: %s\n\n%s", time.Now().Format(logTimeFormat), macEntry.Text, logContent)
		logPane.SetText(logContent)
	})
	macResult := container.NewVBox(
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
		ip := ipEntry.Text
		var mac string
		var err error
		if adapterAutoCheck.Checked {
			mac, err = ip2mac(ip)
		} else {
			mac, err = ip2macWithAdapter(ip, findAdapter(adapters, adapterEntry.Text))
		}
		if err != nil {
			macEntry.SetText("ERROR: " + err.Error())
			logContent = fmt.Sprintf("[%s]\n%s > ERROR\n\n%s", time.Now().Format(logTimeFormat), ip, logContent)
		} else {
			macEntry.SetText(mac)
			logContent = fmt.Sprintf("[%s]\n%s > %s\n\n%s", time.Now().Format(logTimeFormat), ip, mac, logContent)
		}
		logPane.SetText(logContent)
		if logPane.Hidden {
			logPaneLabel.Show()
			logPane.Show()
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
	return container.NewVBox(
		widget.NewLabel("Target IP address:"),
		ipEntry,
		widget.NewLabel("Network adapter:"),
		adapterAutoCheck,
		adapterEntry,
		resolveButton,
		macResult,
		logPaneLabel,
		logPane,
	)
}
