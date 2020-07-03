package main

import (
	"errors"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var macPattern = regexp.MustCompile("([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})")

type enterEntry struct {
	widget.Entry
	tapped func()
}

func (e *enterEntry) setOnEnter(tapped func()) {
	e.tapped = tapped
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn:
		e.tapped()
	default:
		e.Entry.KeyDown(key)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("arping-gui")
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
		mac, err := resolve(ipEntry.Text)
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
		if len(s) > 0 {
			resolveButton.Enable()
		} else {
			resolveButton.Disable()
		}
	}
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Target IP address:"),
		ipEntry,
		resolveButton,
		macResult,
	))
	w.ShowAndRun()
}

func resolve(ip string) (string, error) {
	var pingCmd, arpCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		pingCmd = exec.Command("ping", "-n", "1", ip)
		arpCmd = exec.Command("arp", "-a", ip)
	} else {
		pingCmd = exec.Command("ping", "-c", "1", ip)
		arpCmd = exec.Command("arp", ip)
	}
	if err := pingCmd.Run(); err != nil {
		return "", err
	}
	mac, err := arpCmd.Output()
	if err != nil {
		return "", err
	}
	matches := macPattern.FindAll(mac, -1)
	if len(matches) == 0 {
		return "", errors.New("no data")
	}
	return strings.ToLower(strings.ReplaceAll(string(matches[0]), "-", ":")), nil
}
