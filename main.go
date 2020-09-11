package main

import (
	"log"

	"github.com/int128/ssoexec/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if err := cmd.Run(); err != nil {
		log.Printf("error: %s", err)
	}
}
