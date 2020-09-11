package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/int128/ssoexec/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if err := cmd.Run(os.Args); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			os.Exit(err.ExitCode())
		}
		log.Fatalf("error: %s", err)
	}
}
