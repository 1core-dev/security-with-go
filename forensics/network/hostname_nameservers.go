package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No domain name argument provided.")
	}

	arg := os.Args[1]

	fmt.Printf("Looking up nameservers for: %v\n", arg)

	nameservers, err := net.LookupNS(arg)
	if err != nil {
		log.Fatal(err)
	}
	for _, nameserver := range nameservers {
		fmt.Println(nameserver.Host)
	}
}
