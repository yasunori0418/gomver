package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yasunori0418/gomver/cmd"
	"github.com/yasunori0418/gomver/cmd/increment"
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
					err := cmd.Show(cCtx.Args().First())
					if err != nil {
						return fmt.Errorf("Wrong semantic version: %w", err)
					}
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
							v, err := increment.Major(cCtx.Args().First())
							if err != nil {
								return fmt.Errorf("Wrong semantic version: %w", err)
							}
							fmt.Println(v)
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
							v, err := increment.Minor(cCtx.Args().First())
							if err != nil {
								return fmt.Errorf("Wrong semantic version: %w", err)
							}
							fmt.Println(v)
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
							v, err := increment.Patch(cCtx.Args().First())
							if err != nil {
								return fmt.Errorf("Wrong semantic version: %w", err)
							}
							fmt.Println(v)
							return nil
						},
					},
				},
			},
			{
				Name:    "prerelease",
				Aliases: []string{"pre"},
				Usage: `
					Defines the prerelease value.
					Value must not include the required 'hyphen(-)' prefix.`,
				Action: func(cCtx *cli.Context) error {
					v, err := cmd.PreRelease(cCtx.Args().Get(0), cCtx.Args().Get(1))
					if err != nil {
						return fmt.Errorf("Wrong semantic version: %w", err)
					}
					fmt.Println(v)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
