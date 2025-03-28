package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// By default, the pure Go resolver is used. The resolver can be overridden
// by setting the netdns value of the GODEBUG environment variable
//
// export GODEBUG=netdns=go # force pure Go resolver (Default)
// export GODEBUG=netdns=cgo # force cgo resolver”

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No IP address argument provided.")
	}

	arg := os.Args[1]

	// Parse the IP for validation
	ip := net.ParseIP(arg)
	if ip == nil {
		log.Fatalf("Valid IP not detected. Value provided: %v", arg)
	}

	fmt.Printf("Looking up hostnames for IP address: %v\n", arg)
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostname := range hostnames {
		fmt.Println(hostname)
	}
}
