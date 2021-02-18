package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newMAC2IPTab(w fyne.Window, adapters []adapter) fyne.CanvasObject {
	// adapter select
	adapterNames := adapterNames(adapters)
	adapterEntry := widget.NewSelectEntry(adapterNames)
	if len(adapterNames) > 0 {
		adapterEntry.SetText(adapterNames[0])
	}

	// log pane
	logPaneLabel.Hide()
	logPane.Disable()
	logPane.Hide()

	// result box
	ipEntry := widget.NewEntry()
	var ipCopyButton *widget.Button
	ipCopyButton = widget.NewButton("Copy to clipboard", func() {
		w.Clipboard().SetContent(ipEntry.Text)
		ipCopyButton.Disable()
		ipCopyButton.SetText("Copied!")
		logContent = fmt.Sprintf("[%s]\nCP: %s\n\n%s", time.Now().Format(logTimeFormat), ipEntry.Text, logContent)
		logPane.SetText(logContent)
	})
	ipResult := container.NewVBox(
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
		mac := macEntry.Text
		ip, err := mac2ip(mac, findAdapter(adapters, adapterEntry.Text))
		if err != nil {
			ipEntry.SetText("ERROR: " + err.Error())
			logContent = fmt.Sprintf("[%s]\n%s > ERROR\n\n%s", time.Now().Format(logTimeFormat), mac, logContent)
		} else {
			ipEntry.SetText(ip)
			logContent = fmt.Sprintf("[%s]\n%s > %s\n\n%s", time.Now().Format(logTimeFormat), mac, ip, logContent)
		}
		logPane.SetText(logContent)
		if logPane.Hidden {
			logPaneLabel.Show()
			logPane.Show()
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
	return container.NewVBox(
		widget.NewLabel("Target MAC address:"),
		macEntry,
		widget.NewLabel("Network adapter:"),
		adapterEntry,
		resolveButton,
		ipResult,
		logPaneLabel,
		logPane,
	)
}
