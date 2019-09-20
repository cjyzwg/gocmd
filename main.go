package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cjcmd"
	app.Usage = "cjcmd工具集"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "o",
					Value:       "",
					Usage:       "project owner for create project",
					Destination: &p.Owner,
				},
				cli.StringFlag{
					Name:        "d",
					Value:       "",
					Usage:       "project directory for create project",
					Destination: &p.Path,
				},
				cli.BoolFlag{
					Name:        "proto",
					Usage:       "whether to use protobuf for create project",
					Destination: &p.WithGRPC,
				},
				cli.StringFlag{
					Name:        "m",
					Usage:       "project module name for create project, for `go mod init`",
					Destination: &p.ModuleName,
				},
			},
			Action: runNew,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "cjcmd build",
			Action:  buildAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "cjcmd run",
			Action:  runAction,
		},
		{
			Name:            "tool",
			Aliases:         []string{"t"},
			Usage:           "cjcmd tool",
			Action:          toolAction,
			SkipFlagParsing: true,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "cjcmd version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
		{
			Name:   "self-upgrade",
			Usage:  "cjcmd self-upgrade",
			Action: upgradeAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}