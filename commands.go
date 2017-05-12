package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/yagi2/version-bump-up/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "major",
		Usage:  "[WIP]Bump up major version e.g. 1.2.2 to 2.0.0",
		Action: command.CmdMajor,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "minor",
		Usage:  "[WIP]Bump up minor version e.g. 1.2.2 to 1.3.0",
		Action: command.CmdMinor,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "patch",
		Usage:  "[WIP]Bump up patch version e.g. 1.2.2 to 1.2.3",
		Action: command.CmdPatch,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "set",
		Usage:  "Set to version from command line argument.",
		Action: command.CmdSet,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "check",
		Usage:  "[WIP]Check now version.",
		Action: command.CmdCheck,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
