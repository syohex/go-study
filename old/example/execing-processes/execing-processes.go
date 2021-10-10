package main

import (
	"os"
	"syscall"
	"os/exec"
	"log"
)

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		log.Fatal(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatal(execErr)
	}
}
