package main

import (
	"fmt"
	"os"
	"runtime"

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

func overwriteFyneFont() {
	if runtime.GOOS == "windows" {
		winDir := os.Getenv("WINDIR")
		if len(winDir) == 0 {
			return
		}
		fontPath := determineWindowsFont(winDir + "\\Fonts")
		if err := os.Setenv("FYNE_FONT", fontPath); err != nil {
			fmt.Printf("WARNING: failed to set FYNE_FONT=%s: %v\n", fontPath, err)
		}
	}
}

func determineWindowsFont(fontsDir string) string {
	font := "YuGothM.ttc"
	if _, err := os.Stat(fontsDir + "\\" + font); err == nil {
		return fontsDir + "\\" + font
	}
	font = "meiryo.ttc"
	if _, err := os.Stat(fontsDir + "\\" + font); err == nil {
		return fontsDir + "\\" + font
	}
	font = "msgothic.ttc"
	if _, err := os.Stat(fontsDir + "\\" + font); err == nil {
		return fontsDir + "\\" + font
	}
	font = "segoeui.ttf"
	if _, err := os.Stat(fontsDir + "\\" + font); err == nil {
		return fontsDir + "\\" + font
	}
	return ""
}
