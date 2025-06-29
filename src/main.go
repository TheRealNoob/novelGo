package main

import (
	"log"
	"os"

	"github.com/therealnoob/novelGo/cmd"
)

func main() {
	// Initialize and execute Cobra commands
	if err := cmd.Execute(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
