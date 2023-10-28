package main

import (
	"log"
	"os"
)

func logFatal(err string) {
	log.Printf("%s %s", LOGERROR, err)
	os.Exit(1)
}

func logWarn(warning string) {
	log.Printf("%s %s", LOGWARN, warning)
}

func logInfo(info string) {
	log.Printf("%s %s", LOGINFO, info)
}
