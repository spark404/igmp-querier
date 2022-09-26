//go:build linux

package main

import "syscall"

func BindToDevice(socket int, interfaceName string) error {
	return syscall.BindToDevice(socket, interfaceName)
}
