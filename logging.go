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

// func logWarning(err error) {
// 	if err != nil {
// 		log.Printf("%s", err)
// 	}
// }

func logInfo(info string) {
	log.Printf("%s", info)
}

func terminateProgram() error {
	log.Printf("Bringing down wg interface")
	err := runShellCommand("wg-quick", []string{"down", "pia"})
	return err
}