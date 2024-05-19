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
			{
				Name:    "increment",
				Aliases: []string{"inc"},
				Usage:   "Increment the version value.",
				Subcommands: []*cli.Command{
					{
						Name: "major",
						Usage: `
							Produces the next major version.
							Sets patch to 0.
							Sets minor to 0.
							Increments major number.
							Unsets metadata.
							Unsets prerelease status.`,
						Action: func(cCtx *cli.Context) error {
							version, err := semver.NewVersion(cCtx.Args().Get(0))
							if err != nil {
								return fmt.Errorf("Wrong semantic version: %w", err)
							}
							fmt.Println(version.IncMajor().String())
							return nil
						},
					},
					{
						Name: "minor",
						Usage: `
							Produces the next minor version.
							Sets patch to 0.
							Increments minor number.
							Unsets metadata.
							Unsets prerelease status.`,
						Action: func(cCtx *cli.Context) error {
							version, err := semver.NewVersion(cCtx.Args().Get(0))
							if err != nil {
								return fmt.Errorf("Wrong semantic version: %w", err)
							}
							fmt.Println(version.IncMinor().String())
							return nil
						},
					},
					{
						Name: "patch",
						Usage: `
						Produces the next patch version.
						If the current version does not have prerelease/metadata information,
						it unsets metadata and prerelease values, increments patch number.
						If the current version has any of prerelease or metadata information,
						it unsets both values and keeps current patch value`,
						Action: func(cCtx *cli.Context) error {
							version, err := semver.NewVersion(cCtx.Args().Get(0))
							if err != nil {
								return fmt.Errorf("Wrong semantic version: %w", err)
							}
							fmt.Println(version.IncPatch().String())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
