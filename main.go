package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/user"
	"runtime"
	"sync"

	
)

var TLSClient = getTLSClient()
var waitGroup sync.WaitGroup

func main() {

	var config goPiaConfig

	// begin runtime checks
	if runtime.GOOS != "linux" {
		log.Fatalf("%s This app currently only supports linux OS", LOGERROR)
	}

	if cur_user, _ := user.Current(); cur_user.Uid != "0" {
		log.Fatalf("%s Please run this program as root", LOGERROR)
	}
	// end runtime checks

	// refresh PIA server data
	logInfo("Requesting Server and Region Data")
	serverData, err := getPIAServerData()
	logFatal(err, false)

	// Check configuration file exists and load, else run configuration tool
	configFile, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		logInfo(CONFIG_FILE + " not found, running initial configuration.")
		makeConfiguration(&config, &serverData)
		modifyGID(&config)
		modifyUID(&config)
		chownFiles()
		os.Exit(0)
	} else {
		json.Unmarshal(configFile, &config)
	}

	// Ask user to select region from list
	region := func() Region {
		var regRet Region
		for _, reg := range serverData.Regions {
			if config.PiaRegion == reg.ID {
				regRet = reg
				break
			}
		}
		return regRet
	}()

	rand_server := rand.Intn(len(region.Servers.Wg))
	ip := region.Servers.Wg[rand_server].IP
	// End ask user for region

	logInfo("Creating WireGuard Key Pair")
	keyPair := genKeyPair()

	// Begin connect to PIA
	logInfo(fmt.Sprintf("Connecting to %s - %s\n", region.Name, ip))
	auth, err := getToken(config.PiaUser, config.PiaPass)
	logFatal(err, false)
	logInfo("Got auth token.")

	piaConfig, err := getPIAConfig(
		ip,
		fmt.Sprintf("%d", serverData.Groups.Wg[0].Ports[0]),
		auth.Token,
		keyPair.pubKey,
	)
	logFatal(err, false)

	logInfo(fmt.Sprintf("Server status %s", piaConfig.Status))

	if piaConfig.Status == "OK" {
		logInfo("Got server config successfully.")
		configFile := genWgConfigFile(piaConfig, keyPair)
		writeFile("/etc/wireguard/pia.conf", configFile)
		logInfo("Bringing up wg interface")
		err := runShellCommand("wg-quick", []string{"up", "pia"})
		if err != nil {
			logFatal(err, false)
		}
		logInfo("WireGuard connection established")
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

	logInfo(fmt.Sprintf("Got Signature and Payload, requesting port bind for port %d", payload.Port))

	waitGroup.Add(1)

	// begin forever go routine for port forwarding anonymous function
	go refreshPortForward(payloadAndSignature, &piaConfig)

	// update settings.json
	tConfig := getTransmissionSettings()
	tConfig.BindAddressIpv4 = piaConfig.PeerIP
	tConfig.PeerPort = payload.Port
	tConfig.RPCUsername = config.TransUser
	tConfig.RPCPassword = config.TransPass
	err = writeTransmissionSettings(tConfig)
	logFatal(err, true)
	
	// Start transmission-daemon + stunnel TLS proxy
	logInfo("Starting stunnel")
	err = startStunnel()
	logFatal(err, true)
	logInfo("Starting transmission-daemon")
	err = startTransmission()
	logFatal(err, true)

	waitGroup.Wait()

}

