package main

import (
	"NanjoFan/Tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd: Execute err: %v\n", err)
	}
}
