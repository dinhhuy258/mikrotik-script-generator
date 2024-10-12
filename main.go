package main

import (
	"mikrotik-script-generator/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("Error occurred when execute command", err)
	}
}
