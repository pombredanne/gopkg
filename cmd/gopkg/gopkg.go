package main

import (
	"github.com/go-pkg-org/gopkg/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		Level(zerolog.DebugLevel)

	app := cli.App{
		Name:    "gopkg",
		Version: "0.1.0",
		Usage:   "Package manager for Golang written applications",
		Authors: []*cli.Author{
			{"Aloïs Micard", "alois@micard.lu"},
			{"Fredrik Forsmo", "hello@frozzare.com"},
			{"Johannes Tegnér", "johannes@jitesoft.com"},
		},
		Commands: []*cli.Command{
			{
				Name:      "make",
				Usage:     "create a new package from import-path",
				ArgsUsage: "import-path",
				Action:    cmd.ExecMake,
			},
			{
				Name:      "build",
				Usage:     "build a package from control directory/package",
				ArgsUsage: "control-path",
				Action:    cmd.ExecBuild,
			},
			{
				Name:      "install",
				Usage:     "install a package from path",
				ArgsUsage: "pkg-path",
				Action:    cmd.ExecInstall,
			},
			{
				Name:      "remove",
				Usage:     "remove installed package",
				ArgsUsage: "pkg-name",
				Action:    cmd.ExecRemove,
			},
			{
				Name:  "list",
				Usage: "list packages",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "installed",
						Usage: "list only installed packages",
					},
				},
				Action: cmd.ExecList,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Err(err).Msg("error while running application")
		os.Exit(1)
	}
}
