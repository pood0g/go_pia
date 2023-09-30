package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getTransmissionSettings() TransmissionSettings {
	var config TransmissionSettings
	curConfig, err := os.ReadFile(T_CONF_FILE)
	logFatal(err, true)
	json.Unmarshal(curConfig, &config)
	return config
}

func modifyUID(config *goPiaConfig) error {
	err := runShellCommand("lusermod",
		[]string{
			"-u", config.LinuxUID,
			"transmission",
		})
	return err
}

func modifyGID(config *goPiaConfig) error {
	err := runShellCommand("lgroupmod",
		[]string{
			"-g", config.LinuxUID,
			"transmission",
		})
	return err
}

func chownFiles() error {
	err := runShellCommand("chown",
		[]string{
			"-R",
			"transmission:transmission",
			"/config", "/downloads",
		})
	return err
}

func startTransmission() error {
	err := runShellCommand("su",
		[]string{
			"transmission",
			"-s", "/bin/ash",
			"-c", fmt.Sprintf("transmission-daemon --config-dir %s", T_CONF_DIR),
		})
	return err
}

func startStunnel() error {
	err := runShellCommand("su",
		[]string{
			"stunnel",
			"-s", "/bin/ash",
			"-c", "stunnel",
		})
	return err
}

func writeTransmissionSettings(settings TransmissionSettings) error {
	config, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = os.WriteFile(T_CONF_FILE, config, 0666)
	return err
}
