package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var (
	normalText = fyne.TextStyle{}
	boldText   = fyne.TextStyle{Bold: true}
	italicText = fyne.TextStyle{Italic: true}
)

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
