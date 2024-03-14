package main

import (
	"clidev/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
