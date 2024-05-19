package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"github.com/Masterminds/semver/v3"
)

func main() {
	app := &cli.App{
		Name:        "gomver",
		Description: "Controll semantic vertion.",
		Commands: []*cli.Command{
			{
				Name:    "parse",
				Aliases: []string{"p"},
				Usage:   "入力された値がsemverか確認し、詳細を表示します。",
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("Arg1: %v\n", cCtx.Args().Get(0))
					v, err := semver.NewVersion(cCtx.Args().Get(0))
					if err != nil {
						return fmt.Errorf("Wrong semantic version: %w", err)
					}
					fmt.Printf("Parsed argument semantic version: %+v\n", v)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
