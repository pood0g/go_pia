package main

import (
	"log"
)

func logFatal(err error, wgUp bool) {
	if err != nil {
		if wgUp{
			err := terminateProgram()
			log.Printf("%s", err)
		}
		log.Fatalf("%s", err)
	}
}

func logWarn(warning string) {
	log.Printf("%s %s",LOGWARN, warning)
}

func logInfo(info string) {
	log.Printf("%s %s", LOGINFO, info)
}

func terminateProgram() error {
	log.Printf("Bringing down wg interface")
	err := runShellCommand("wg-quick", []string{"down", "pia"})
	if err != nil {
		return err
	}
	log.Printf("Terminating Transmission Daemon")
	err = runShellCommand("pkill", []string{"-9", "transmission-daemon"})
	return err
}