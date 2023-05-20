package main

import (
	"fmt"
	// "sync"
)

// var waitGroup sync.WaitGroup

func main() {
	// waitGroup.Add(1)
	regions := getRegionData()
	
	for _, p := range regions.Regions {
		fmt.Printf("%s\n", p.Name)
		fmt.Printf("\t%s\n", p.Servers.Wg[0].IP)
	}

	keyPair := genKeyPair()
	fmt.Println(keyPair.prvKey)
	fmt.Println(keyPair.pubKey)
	// waitGroup.Wait()
}