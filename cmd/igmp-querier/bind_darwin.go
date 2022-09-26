//go:build darwin

package main

import (
	"fmt"
	"net"
	"syscall"
)

func BindToDevice(socket int, interfaceName string) error {
	iFaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	index := -1
	for _, i := range iFaces {
		if i.Name == interfaceName {
			index = i.Index
		}
	}
	if index == -1 {
		return fmt.Errorf("no interface %s", interfaceName)
	}

	return syscall.SetsockoptInt(socket, syscall.IPPROTO_IP, syscall.IP_BOUND_IF, index)
}
