package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/user"
	"runtime"
)

func main() {

	var choice uint8
	username := os.Getenv("PIA_USER")
	password := os.Getenv("PIA_PASS")

	// Check environment variables are set
	if len(username) == 0 || len(password) == 0 {
		log.Fatalf("%s PIA_USER or PIA_PASS environment variables not set!", logError)
	}

	if runtime.GOOS != "linux" {
		log.Fatalf("%s This app currently only supports linux OS", logError)
	}

	if cur_user, _ := user.Current(); cur_user.Uid != "0" {
		log.Fatalf("%s Please run this program as root", logError)
	}

	serverData, err := getPIAServerData()
	logFatal(err, false)
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
	logFatal(err, false)
	log.Printf("Got auth token.\n")

	piaConfig, err := getPIAConfig(
		ip,
		fmt.Sprintf("%d", serverData.Groups.Wg[0].Ports[0]),
		auth.Token,
		keyPair.pubKey,
	)
	logFatal(err, false)

	log.Printf("Server status %s", piaConfig.Status)

	if piaConfig.Status == "OK" {
		log.Printf("Got server config successfully.")
		configFile := genWgConfigFile(piaConfig, keyPair)
		writeFile("/etc/wireguard/pia.conf", configFile)
		log.Printf("Bringing up wg interface")
		err := runShellCommand("wg-quick", []string{"up", "pia"})
		if err != nil {
			logFatal(err, false)
		}
		log.Printf("WireGuard connection established")
	} else {
		log.Fatalln("failed")
	}

	payloadAndSignature, payload, err := getPFSignature(
		piaConfig.ServerVIP,
		"19999",
		auth.Token,
	)
	if err != nil {
		log.Printf("Port Forwarding failed - %s", err)
	}

	log.Printf("Got Signature and Payload, requesting port bind for port %d", payload.Port)

	pfStatus, err := requestBindPort(
		piaConfig.ServerVIP,
		"19999",
		payloadAndSignature,
	)
	logFatal(err, true)

	fmt.Printf("%s", pfStatus)
}
