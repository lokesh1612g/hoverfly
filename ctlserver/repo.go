package main

import (
	"os"
	"os/exec"
)

var current State

func RepoGetCurrentState() State {
	return current
}

func RepoCreateState(s State) State {
	cmd := exec.Command("sh", "-c", s.Command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	current = s
	return current
}
