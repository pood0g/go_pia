package main

import (
	// "fmt"
	"bytes"
	"io"
	"log"
	"net/http"
	// "os/exec"
)

func handleFatal(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func makeGETRequest(url string) []byte {

	resp, err := http.Get(url)
	handleFatal(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	handleFatal(err)

	return body
}

func makePOSTRequest(url, contentType string, body []byte) []byte {
	reqBody := bytes.NewReader(body)
	resp, err := http.Post(url, contentType, reqBody)
	handleFatal(err)
	defer resp.Body.Close()
	respBody , err := io.ReadAll(resp.Body)
	handleFatal(err)

	return respBody
}

// func runShellCommand(command string, args []string) {
// 	defer waitGroup.Done()
// 	cmd, err:= exec.Command(command, args...).Output()
// 	handleFatal(err)

// 	fmt.Printf("%s", cmd)
// }
