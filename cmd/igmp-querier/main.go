package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		usage()
	}

	version := args[0]
	interfaceName := args[1]

	queryVersion, err := strconv.Atoi(version)
	if err != nil || (queryVersion != 2 && queryVersion != 3) {
		log.Fatalf("%s is not a valid igmp version, expected 2 or 3", version)
	}

	addr := syscall.SockaddrInet4{
		Port: 0,
		Addr: [4]byte{224, 0, 0, 1},
	}

	igmpQuery := []byte{
		0x11,
		0x64,
		0xee,
		0x9b,
		0x00,
		0x00,
		0x00,
		0x00,
	}

	if queryVersion == 3 {
		igmpQuery = append(igmpQuery, []byte{
			0x0,
			0x0,
			0x0,
			0x0,
		}...)
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_IGMP)
	if err != nil {
		log.Fatal("Socket create failed: ", err)
	}

	err = BindToDevice(fd, interfaceName)
	if err != nil {
		syscall.Close(fd)
		log.Fatal("Set IP_BOUND_IF sockopt failed: ", err)
	}

	log.Printf("Sending IGMPv%d Group membership query on interface %s", queryVersion, interfaceName)
	err = syscall.Sendto(fd, igmpQuery, 0, &addr)
	if err != nil {
		log.Fatal("Sendto: ", err)
	}
}

func usage() {
	fmt.Printf("Usage:\n")
	fmt.Printf(" %s <2|3> <interfacename>", os.Args[0])
	fmt.Printf("\n")
	os.Exit(1)
}
