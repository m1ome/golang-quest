package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	fmt.Println("Quest: Given an IPv4 adress and subnet mask, compute the network, broadcast and first/last addresses")
	fmt.Println("Usage: [ip] [mask]")
	fmt.Println("Example: 0.0.0.0 255.255.255.255")

	if len(args) != 3 {
		fmt.Println("Error: Too few arguments")
		os.Exit(0)
	}

	ipv4 := args[1]
	netmask := args[2]

	// Simple verification regexp
	verifyRegexp := regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}$`)

	// Checking IPv4 Address
	if !verifyRegexp.MatchString(ipv4) {
		fmt.Println("Error: ip in wrong")
		os.Exit(0)
	}

	// Checking bitmask
	if !verifyRegexp.MatchString(netmask) {
		fmt.Println("Error: mask in wrong")
		os.Exit(0)
	}

	ipBytes := readAddress(ipv4)
	netmaskBytes := readAddress(netmask)

	fmt.Println()

	fmt.Printf("IP address: %s", stringify(ipBytes))
	fmt.Println()

	fmt.Printf("Netmask: %s", stringify(netmaskBytes))
	fmt.Println()

	// Calculate Network Address
	var networkAddressBytes [4]byte
	for i := 0; i < 4; i++ {
		networkAddressBytes[i] = ipBytes[i] & netmaskBytes[i]
	}

	fmt.Printf("Network address: %s", stringify(networkAddressBytes))
	fmt.Println()

	// Calculate Broadcast
	var broadcastAddressBytes [4]byte
	for i := 0; i < 4; i++ {
		broadcastAddressBytes[i] = ipBytes[i] | ^netmaskBytes[i]
	}
	fmt.Printf("Broadcast address: %s", stringify(broadcastAddressBytes))
	fmt.Println()

}

func readAddress(address string) [4]byte {
	var bytes [4]byte

	for idx, octet := range strings.Split(address, ".") {
		octet, err := strconv.ParseInt(octet, 10, 32)

		if err != nil {
			fmt.Printf("Error: reading %s (%s) at %i octet", address, err, idx+1)
			fmt.Println()
			os.Exit(0)
		}

		if octet > 255 || octet < 0 {
			fmt.Printf("Error: reading %s at %i octet must be at 0...255 range", address, idx+1)
			fmt.Println()
			os.Exit(0)
		}

		bytes[idx] = byte(octet)
	}

	return bytes
}

func stringify(address [4]byte) string {
	return strconv.FormatInt(int64(address[0]), 10) + "." +
		strconv.FormatInt(int64(address[1]), 10) + "." +
		strconv.FormatInt(int64(address[2]), 10) + "." +
		strconv.FormatInt(int64(address[3]), 10)
}
