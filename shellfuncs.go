package main

import (
	"os/exec"
	"fmt"
)



func runShellCommand(command string, args []string) error {
	defer waitGroup.Done()
	cmd, err:= exec.Command(command, args...).CombinedOutput()
	fmt.Printf("%s", cmd)
	return err
}