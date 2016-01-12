package main

import (
	dhcp "github.com/d2g/dhcp4client/dhcp4client"
	"log"
	"fmt"
	"net"
)

func main() {
//	m, err := net.ParseMAC("DC-53-60-D8-5E-D8")
	m, err := net.ParseMAC("DC-53-60-D8-5E-D7")
	if err != nil {
		log.Printf("MAC Error:%v\n", err)
	}

	c, err := dhcp.NewInetSock()
	if err != nil {
		log.Fatal("Client Connection Generation", err)
	}

	client, err := dhcp.New(dhcp.HardwareAddr(m), dhcp.Connection(c))
	if err != nil {
		log.Fatal("Unable to create client", err)
	}

	defer client.Close()

	ok, packet, err := client.Request()
	if err != nil {
		log.Fatal("Unable to request", err)
	}

	fmt.Println(ok)
	fmt.Println(string(packet.Options()))
}