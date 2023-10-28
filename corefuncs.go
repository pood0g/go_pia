package main

import (
	"fmt"
	"os"
	"os/exec"
	"encoding/json"

	"golang.org/x/term"
)

func runShellCommand(command string, args []string) error {
	cmd, err := exec.Command(command, args...).CombinedOutput()
	fmt.Printf("%s", cmd)
	return err
}

func makeConfiguration(config *goPiaConfig, serverData *RegionData) {
	fmt.Printf("Enter PIA Username: ")
	fmt.Scanln(&config.PiaUser)
	fmt.Printf("Enter PIA password: ")
	passBytes, err := term.ReadPassword(0)
	if err != nil {
		logFatal(err.Error())
	}
	fmt.Println()
	config.PiaPass = string(passBytes)

	fmt.Printf("Enter Transmission Username: ")
	fmt.Scanln(&config.TransUser)
	fmt.Printf("Enter Transmission password: ")
	tPassBytes, err := term.ReadPassword(0)
	if err != nil {
		logFatal(err.Error())
	}
	fmt.Println()

	fmt.Printf("Enter linux UID: ")
	fmt.Scanln(&config.LinuxUID)
	fmt.Printf("Enter linux GID: ")
	fmt.Scanln(&config.LinuxGID)

	config.TransPass = string(tPassBytes)

	config.PiaRegion = pickRegion(serverData).ID

	if len(config.PiaUser) == 0 || len(config.PiaPass) == 0 {
		logWarn(LOGERROR + "Configuration items cannot be blank.")
		os.Exit(1)
	}
	

	jsonData, _ := json.Marshal(config)

	os.WriteFile(CONFIG_FILE, jsonData, 0600)
}

func restartServices() error {
	logWarn("Connection Interupted, restarting!")
	
	logInfo("Bringing down wg interface")
	err := runShellCommand("wg-quick", []string{"down", "pia"})
	if err != nil {
		return err
	}
	logInfo("Terminating Transmission Daemon")
	err = runShellCommand("pkill", []string{"-9", "transmission-daemon"})
	if err != nil {
		return err
	}

	

	logInfo("Bringing up wg interface")
	err = runShellCommand("wg-quick", []string{"down", "pia"})
	if err != nil {
		return err
	}
	logInfo("Restarting transmission-daemon")
	err = startTransmission()
	if err != nil {
		logFatal(err.Error())
	}
	return err
}