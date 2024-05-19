package main

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "gomver",
		Description: "Controll semantic vertion.",
		Commands: []*cli.Command{
			{
				Name:  "show",
				Usage: "Check if the entered value is semver and display details.",
				Action: func(cCtx *cli.Context) error {
					version, err := semver.NewVersion(cCtx.Args().Get(0))
					if err != nil {
						return fmt.Errorf("Wrong semantic version: %w", err)
					}
					fmt.Printf("inputed version: %+v\n", version.Original())
					fmt.Printf("Major value in version: %+v\n", version.Major())
					fmt.Printf("Minor value in version: %+v\n", version.Minor())
					fmt.Printf("Patch value in version: %+v\n", version.Patch())
					fmt.Printf("Prerelease value in version: %+v\n", version.Prerelease())
					fmt.Printf("Metadata value in version: %+v\n", version.Metadata())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
