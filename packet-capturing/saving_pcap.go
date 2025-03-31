package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	deviceName        = "en0"
	snapshotLen int32 = 1024
	promiscuous       = false
	err         error
	timeout     = -1 * time.Second
)

func main() {
	// Open output file and write header
	file, _ := os.Create("test.pcap")
	writer := pcapgo.NewWriter(file)
	writer.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)

	// Open the device for capturing
	handle, err := pcap.OpenLive(deviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatalf("error opening device %s: %v", deviceName, err)
	}
	defer handle.Close()

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	var packetCount int

	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
		writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		// Only capture 100 packets and stop
		if packetCount > 100 {
			break
		}
	}

}
