package main

import (
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

func ip2mac(ip string, adapter adapter) (string, error) {
	var pingCmd, arpCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		pingCmd = exec.Command("ping", "-n", "1", ip)
		arpCmd = exec.Command("arp", "-a", ip, "-N", adapter.ip)
	case "darwin":
		pingCmd = exec.Command("ping", "-b", adapter.name, "-c", "1", ip)
		arpCmd = exec.Command("arp", "-i", adapter.name, ip)
	default:
		pingCmd = exec.Command("ping", "-I", adapter.name, "-c", "1", ip)
		arpCmd = exec.Command("arp", "-i", adapter.name, ip)
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
	return strings.ReplaceAll(strings.ToLower(string(matches[0])), "-", ":"), nil
}

func mac2ip(mac string, adapter adapter) (string, error) {
	var pingCmd, arpCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		mac = strings.ReplaceAll(strings.ToLower(mac), ":", "-")
		pingCmd = exec.Command("ping", "-n", "1", adapter.broadcast)
		arpCmd = exec.Command("arp", "-a", "-N", adapter.ip)
	case "darwin":
		mac = strings.ReplaceAll(strings.ToLower(mac), "-", ":")
		pingCmd = exec.Command("ping", "-b", adapter.name, "-c", "1", adapter.broadcast)
		arpCmd = exec.Command("arp", "-a", "-i", adapter.name)
	default:
		mac = strings.ReplaceAll(strings.ToLower(mac), "-", ":")
		pingCmd = exec.Command("ping", "-I", adapter.name, "-c", "1", adapter.broadcast, "-b")
		arpCmd = exec.Command("arp", "-a", "-i", adapter.name)
	}
	if err := pingCmd.Run(); err != nil {
		return "", err
	}
	arpOut, err := arpCmd.Output()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(bytes.NewReader(arpOut))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, mac) {
			continue
		}
		matches := ipPattern.FindAllString(line, -1)
		if len(matches) == 0 {
			return "", errors.New("no data")
		}
		return matches[0], nil
	}
	return "", errors.New("no data")
}
