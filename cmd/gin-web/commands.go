package main

import (
	"github.com/urfave/cli"
)

type (
	Commands struct {
		PathFile string
		CliApp   *cli.App
	}
)

func NewCommands() *Commands {
	return &Commands{}
}

func (cmd *Commands) Commands() *Commands {
	cmd.CliApp = cli.NewApp()

	cmd.CliApp.Name = "gin web server"
	cmd.CliApp.Usage = "gin web server"
	cmd.CliApp.Author = "yx shi"
	cmd.CliApp.Email = "357705148@qq.com"
	cmd.CliApp.Version = "0.0.1"

	cmd.CliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config,c",
			Value:       "config/local/config.yaml",
			Destination: &cmd.PathFile,
			Usage:       "server config path",
		},
	}

	return cmd
}
