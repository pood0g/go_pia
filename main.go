package main

import (
	"fmt"
	"os"
	// "sync"
)

// var waitGroup sync.WaitGroup

func main() {
	// waitGroup.Add(1)
	regions := getRegionData()
	
	for _, p := range regions.Regions {
		fmt.Printf("%s\n", p.Name)
		for _, ip := range p.Servers.Wg {
			fmt.Printf("\t%s\n", ip.IP)
		}
	}

	username := os.Getenv("PIA_USER")
	password := os.Getenv("PIA_PASS")
	auth := getToken(username, password)

	fmt.Printf("%s\n", auth.Token)

	keyPair := genKeyPair()
	fmt.Println(keyPair.prvKey)
	fmt.Println(keyPair.pubKey)
	// waitGroup.Wait()
}