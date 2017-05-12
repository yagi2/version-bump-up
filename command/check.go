package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/yagi2/version-bump-up/util"
)

func CmdCheck(c *cli.Context) {
	if (len(c.Args())) != 0 {
		fmt.Fprintf(os.Stderr, "Not enough arguments.\n$ version-bump-up check\n")
		os.Exit(1)
	}

	gradle := util.OpenBuildGradle()

	for _, element := range gradle {
		if strings.Contains(element, "versionName ") || strings.Contains(element, "versionCode") {
			fmt.Println(strings.TrimSpace(element))
		}
	}
}
