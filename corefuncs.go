package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func runShellCommand(command string, args []string) error {
	cmd, err := exec.Command(command, args...).CombinedOutput()
	fmt.Printf("%s", cmd)
	return err
}

func makeConfiguration(config *goPiaConfig, serverData *RegionData) {
	fmt.Println("The following username and password will be used to connect")
	fmt.Println("to the Private Internet Access VPN")
	fmt.Printf("Enter PIA Username: ")
	fmt.Scanln(&config.PiaUser)
	fmt.Printf("Enter PIA password: ")
	fmt.Scanln(&config.PiaPass)
	fmt.Println("The following username and password will be used to access the")
	fmt.Println("Transmission daemon web interface")
	fmt.Printf("Enter Transmission Username: ")
	fmt.Scanln(&config.TransUser)
	fmt.Printf("Enter Transmission password: ")
	fmt.Scanln(&config.TransPass)
	fmt.Println("The following UID/GID should be set to the same values as")
	fmt.Println("the owner of the directory where downloads are to be stored.")
	fmt.Printf("Enter linux UID: ")
	fmt.Scanln(&config.LinuxUID)
	fmt.Printf("Enter linux GID: ")
	fmt.Scanln(&config.LinuxGID)
	// fmt.Println("Please enter any subnets to exclude from the wireguard VPN eg. 192.168.1.1/24")
	// fmt.Println("")

	config.PiaRegion = pickRegion(serverData).ID

	if len(config.PiaUser) == 0 || len(config.PiaPass) == 0 || len(config.TransUser) == 0 {
		logWarn(LOGERROR + "Configuration items cannot be blank.")
		os.Exit(1)
	}

	jsonData, _ := json.MarshalIndent(config, "", "    ")

	os.WriteFile(CONFIG_FILE, jsonData, 0600)
}

func killServices() error {
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
	return err
}
