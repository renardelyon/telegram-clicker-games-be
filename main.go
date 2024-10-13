package main

import (
	"log"
	"os"
	"telegram-clicker-game-be/cmd"
	"telegram-clicker-game-be/config"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}

	if err := cmd.Cli(cfg).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
