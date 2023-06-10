package main

import (
	"math/rand"
	"fmt"
	"os"
	"log"
	// "sync"
)

// var waitGroup sync.WaitGroup

func main() {
	var choice uint8
	username := os.Getenv("PIA_USER")
	password := os.Getenv("PIA_PASS")
	// waitGroup.Add(1)
	serverData, err := getPIAServerData()
	handleFatal(err)
	keyPair := genKeyPair()

	fmt.Printf("Available regions:\n\n")
	for i, p := range serverData.Regions {
		fmt.Printf("\t [%d] %s\n", i, p.Name)

	}

	fmt.Printf("\nPick a Region: ")
	fmt.Scanln(&choice)

	rand_server := rand.Intn(len(serverData.Regions[choice].Servers.Wg))
	ip := serverData.Regions[choice].Servers.Wg[rand_server].IP

	log.Printf("Connecting to %s - %s\n", serverData.Regions[choice].Name, ip)
	auth, err := getToken(username, password)
	handleFatal(err)
	log.Printf("Got auth token.\n")

	piaConfig, err := getPIAConfig(
		ip,
		fmt.Sprintf("%d", serverData.Groups.Wg[0].Ports[0]),
		auth.Token,
		keyPair.pubKey,
	)
	handleFatal(err)

	log.Printf("Server status %s", piaConfig.Status)

	if piaConfig.Status == "OK" {
		log.Printf("Got server public key.\n")
	} else {
		log.Fatalln("failed")
	}
	
	// waitGroup.Wait()
}
