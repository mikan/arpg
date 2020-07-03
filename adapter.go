package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strings"
)

type adapter struct {
	name      string
	ip        string
	broadcast string
}

func adapters() ([]adapter, error) {
	var adapters []adapter
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to collect network interface: %w", err)
	}
	for _, ifEntry := range interfaces {
		if ifEntry.HardwareAddr == nil || !strings.Contains(ifEntry.Flags.String(), "up") {
			continue // ignore odd or down adapters
		}
		addresses, err := ifEntry.Addrs()
		if err != nil || len(addresses) == 0 {
			continue // ignore unassigned adapters
		}
		for _, address := range addresses {
			if !strings.Contains(address.String(), ".") {
				continue
			}
			ip, n, err := net.ParseCIDR(address.String())
			if err != nil {
				continue
			}
			broadcast := broadcast(n)
			adapters = append(adapters, adapter{ifEntry.Name, ip.String(), broadcast})
		}
	}
	if len(adapters) == 0 {
		return nil, errors.New("no adapters available")
	}
	return adapters, nil
}

func broadcast(n *net.IPNet) string {
	bc := make(net.IP, len(n.IP.To4()))
	binary.BigEndian.PutUint32(bc, binary.BigEndian.Uint32(n.IP.To4())|^binary.BigEndian.Uint32(net.IP(n.Mask).To4()))
	return bc.String()
}

func adapterNames(adapters []adapter) []string {
	s := make([]string, len(adapters))
	for i, a := range adapters {
		s[i] = fmt.Sprintf("%s [%s]", a.name, a.ip)
	}
	return s
}

func findAdapter(adapters []adapter, name string) adapter {
	for _, a := range adapters {
		if strings.HasPrefix(name, a.name) {
			return a
		}
	}
	return adapter{}
}
