// +build !windows

package main

import "os/exec"

func prepareBackgroundCommand(_ *exec.Cmd) {
	// no-op
}
