package main

import (
	"context"
	"log"
	"os"

	"github.com/soadmized/sentinel/cmd"
	"github.com/soadmized/sentinel/internal/config"
)

func main() {
	ctx := context.Background()
	conf := config.Read()
	args := os.Args[1:]

	err := cmd.Run(ctx, conf, args)
	if err != nil {
		log.Fatal(err)
	}
}
