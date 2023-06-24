package main

import (
	"log"
)

func logFatal(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

// func logWarning(err error) {
// 	if err != nil {
// 		log.Printf("%s", err)
// 	}
// }