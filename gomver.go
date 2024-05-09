package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "gomver",
		Description: "Controll semantic vertion.",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
