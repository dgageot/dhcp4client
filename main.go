package main

import (
	"github.com/d2g/dhcp4client/dhcp4client"
	"log"
	"fmt"
)

func main() {
	c, err := dhcp4client.NewInetSock()
	if err != nil {
		log.Fatal("Client Connection Generation", err)
	}

	client, err := dhcp4client.New(dhcp4client.Connection(c))
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