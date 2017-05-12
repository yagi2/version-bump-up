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
		Usage:  "",
		Action: command.CmdMajor,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "minor",
		Usage:  "",
		Action: command.CmdMinor,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "patch",
		Usage:  "",
		Action: command.CmdPatch,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "set",
		Usage:  "",
		Action: command.CmdSet,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "check",
		Usage:  "",
		Action: command.CmdCheck,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
