package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"sync"
)

var TLSClient = getTLSClient()
var waitGroup sync.WaitGroup
var config goPiaConfig
var region Region
var serverData RegionData

func main() {

	// Perform runtime checks
	if runtime.GOOS != "linux" {
		logFatal("This app currently only supports linux OS")
	}

	if cur_user, _ := user.Current(); cur_user.Uid != "0" {
		logFatal("Please run this program as root")
	}

	// Fetch PIA server data
	logInfo("Requesting Server and Region Data")
	err := getPIAServerData()
	if err != nil {
		logFatal(err.Error())
	}

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

	// Get configured region from serverData.Regions
	region = func() Region {
		var regRet Region
		for _, reg := range serverData.Regions {
			if config.PiaRegion == reg.ID {
				regRet = reg
				break
			}
		}
		return regRet
	}()

	// Begin connect to PIA
	piaConfig, auth := connectToPIA(&config, &region, &serverData)

	// Get Port Forwarding Auth
	payloadAndSignature, portNo, err := getPFSignature(
		piaConfig.ServerVIP,
		"19999",
		auth.Token,
	)
	if err != nil {
		log.Printf("Port Forwarding failed - %s", err)
	}

	logInfo(fmt.Sprintf("Got Signature and Payload, requesting port bind for port %d", portNo))

	waitGroup.Add(1)

	// begin forever go routine for port forwarding anonymous function
	go refreshPortForward(payloadAndSignature, &piaConfig)

	// update settings.json
	tConfig := getTransmissionSettings()
	tConfig.BindAddressIpv4 = piaConfig.PeerIP
	tConfig.PeerPort = portNo
	tConfig.RPCUsername = config.TransUser
	tConfig.RPCPassword = config.TransPass
	err = writeTransmissionSettings(tConfig)
	if err != nil {
		logFatal(err.Error())
	}

	// Start transmission-daemon + stunnel TLS proxy
	logInfo("Starting stunnel")
	err = startStunnel()
	if err != nil {
		logFatal(err.Error())
	}
	logInfo("Starting transmission-daemon")
	err = startTransmission()
	if err != nil {
		logFatal(err.Error())
	}
	waitGroup.Wait()

}
