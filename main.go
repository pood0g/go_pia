package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/user"
	"runtime"
	// "sync"
)

// var waitGroup sync.WaitGroup

func main() {

	// future use for running shell commands
	// waitGroup.Add(1)

	var choice uint8
	username := os.Getenv("PIA_USER")
	password := os.Getenv("PIA_PASS")

	// Check environment variables are set
	if len(username) == 0 || len(password) == 0 {
		log.Fatalf("%sERROR:%s PIA_USER or PIA_PASS environment variables not set!", Red, Reset)
	}

	if runtime.GOOS != "linux" {
		log.Fatalf("%sERROR:%s This app currently only supports linux OS", Red, Reset)
	}

	if cur_user, _ := user.Current(); cur_user.Uid != "0" {
		log.Fatalf("%sERROR:%s Please run this program as root", Red, Reset)
	}

	serverData, err := getPIAServerData()
	logFatal(err)
	keyPair := genKeyPair()

	fmt.Printf("Available regions:\n\n")
	for i, p := range serverData.Regions {
		fmt.Printf("\t %s[%d]%s %s\n", Green, i, Reset, p.Name)

	}

	fmt.Printf("\nPick a Region: ")
	fmt.Scanln(&choice)

	rand_server := rand.Intn(len(serverData.Regions[choice].Servers.Wg))
	ip := serverData.Regions[choice].Servers.Wg[rand_server].IP

	log.Printf("Connecting to %s - %s\n", serverData.Regions[choice].Name, ip)
	auth, err := getToken(username, password)
	logFatal(err)
	log.Printf("Got auth token.\n")

	piaConfig, err := getPIAConfig(
		ip,
		fmt.Sprintf("%d", serverData.Groups.Wg[0].Ports[0]),
		auth.Token,
		keyPair.pubKey,
	)
	logFatal(err)

	log.Printf("Server status %s", piaConfig.Status)

	if piaConfig.Status == "OK" {
		log.Printf("Got server public key.\n\n")
		configFile := genWgConfigFile(piaConfig, keyPair)
		writeFile("/etc/wireguard/pia.conf", configFile)
	} else {
		log.Fatalln("failed")
	}

	// waitGroup.Wait()
}
