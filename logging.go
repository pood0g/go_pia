package main

import (
	"log"
)

func logFatal(err error, wgUp bool) {
	if err != nil {
		if wgUp{
			log.Printf("Bringing down wg interface")
			err = runShellCommand("wg-quick", []string{"down", "pia"})
		}
		log.Fatalf("%s", err)
	}
}

// func logWarning(err error) {
// 	if err != nil {
// 		log.Printf("%s", err)
// 	}
// }