package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func handleFatal(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func makeGETRequest(url string) string {

	resp, err := http.Get(url)
	handleFatal(err)

	body, err := io.ReadAll(resp.Body)
	handleFatal(err)

	return string(body)
}

func runShellCommand(command string, args []string) {
	cmd, err:= exec.Command(command, args...).Output()
	handleFatal(err)

	fmt.Print(cmd)
}
