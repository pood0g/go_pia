package main

import (
	"log"
)

func logFatal(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
