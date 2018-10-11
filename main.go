package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var pcapPath = flag.String("pcap", "", "path of the configuration directory")

func main() {
	flag.Parse()
	handle, err := pcap.OpenOffline(*pcapPath)
	if err != nil {
		log.Printf("failed opening pcap: %v", err)
		os.Exit(1)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()

	packetsBytes := []string{}
	for packet := range packets {
		packetsBytes = append(packetsBytes, byteArrayToString(packet.Data()))
	}

	fmt.Printf("packetBytes := [][]byte{\n%v\n}", strings.Join(packetsBytes, "\n"))
}

func byteArrayToString(bytes []byte) string {
	hex := hex.EncodeToString(bytes)

	hexArray := []string{}
	char := 2
	for i := 0; i < len(hex); i += char {
		hexArray = append(hexArray, fmt.Sprintf("0x%v", hex[i:i+2]))
	}

	return fmt.Sprintf("[]byte{%v},", strings.Join(hexArray, ", "))
}
