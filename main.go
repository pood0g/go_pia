package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(1)
	ipAddr := makeGETRequest("http://api.ipify.org")
	fmt.Println(ipAddr)
	args := []string{"-c 10", "8.8.8.8"}
	go runShellCommand("ping", args)
	keyPair := genKeyPair()
	fmt.Println(keyPair.prvKey)
	fmt.Println(keyPair.pubKey)
	waitGroup.Wait()
}