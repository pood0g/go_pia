package main

import (
	"fmt"
	"os"
	// "sync"
)

// var waitGroup sync.WaitGroup

func main() {
	// waitGroup.Add(1)
	serverData := getPIAServerData()
	keyPair := genKeyPair()

	for i, p := range serverData.Regions {
		fmt.Printf("[%d] %s\n", i, p.Name)
		for _, ip := range p.Servers.Wg {
			fmt.Printf("\t%s\n", ip.IP)
		}
	}

	username := os.Getenv("PIA_USER")
	password := os.Getenv("PIA_PASS")
	auth, err := getToken(username, password)
	handleFatal(err)
	fmt.Printf("%s\n", auth.Token)

	piaConfig := getPIAConfig(
		"154.6.147.75",
		fmt.Sprintf("%d", serverData.Groups.Wg[0].Ports[0]),
		auth.Token,
		keyPair.pubKey,
	)

	fmt.Println(piaConfig.ServerKey, piaConfig.Status)
	
	// waitGroup.Wait()
}
