package main

import (
	"log"
	"mikrotik-script-generator/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("Error occurred when execute command", err)
	}
}
