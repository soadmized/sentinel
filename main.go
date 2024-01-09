package main

import (
	"log"

	"sentinel/cmd"
	"sentinel/internal/config"
)

func main() {
	conf := config.Read()

	err := cmd.Run(conf)
	if err != nil {
		log.Fatal(err)
	}
}
