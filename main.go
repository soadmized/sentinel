package main

import (
	"log"

	"github.com/soadmized/sentinel/cmd"
	"github.com/soadmized/sentinel/internal/config"
)

func main() {
	conf := config.Read()

	err := cmd.Run(conf)
	if err != nil {
		log.Fatal(err)
	}
}
