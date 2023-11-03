package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	var config goPiaConfig
	
	// Perform runtime checks
	if runtime.GOOS != "linux" {
		logFatal("This app currently only supports linux OS")
	}

	if cur_user, _ := user.Current(); cur_user.Uid != "0" {
		logFatal("Please run this program as root")
	}

	// Fetch PIA server data
	logInfo("Requesting Server and Region Data")
	serverData, err := getPIAServerData()
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

	// Get configured region from serverData.Regions anonymous function.
	region := func() Region {
		var ret Region
		
		for _, elem := range serverData.Regions {
			if config.PiaRegion == elem.ID {
				ret = elem
				break
			}
		}
		return ret
	}()

	// Begin connect to PIA
	piaConfig, auth := connectToPIA(&config, &region, &serverData)

	// Get Port Forwarding Auth
	payloadAndSignature, portNo, err := getPortForwardSignature(
		piaConfig.ServerVIP,
		"19999",
		auth.Token,
	)
	if err != nil {
		logWarn(fmt.Sprintf("Port Forwarding failed - %s", err))
	}

	logInfo(fmt.Sprintf("Got Signature and Payload, requesting port bind for port %d", portNo))

	waitGroup.Add(1)

	// begin forever go routine for port forwarding anonymous function
	go refreshPortForward(&payloadAndSignature, &piaConfig)

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

	// Start transmission-daemon
	logInfo("Starting transmission-daemon")
	err = startTransmission()
	if err != nil {
		logFatal(err.Error())
	}
	waitGroup.Wait()

}
