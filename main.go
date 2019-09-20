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
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "cjcmd version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}