package command

import (
	"fmt"
	"strings"

	"os"

	"regexp"

	"github.com/codegangsta/cli"
	"github.com/yagi2/version-bump-up/util"
)

func CmdSet(c *cli.Context) {
	if (len(c.Args())) != 2 {
		fmt.Fprintf(os.Stderr, "Not enough arguments.\n$ version-bump-up set [versionName] [versionCode]\n")
		os.Exit(1)
	}

	gradle := util.OpenBuildGradle()

	newVersionName := "\"" + (c.Args()[0]) + "\""
	newVersionCode := c.Args()[1]

	newGradle := make([]string, 0, 100)
	for _, element := range gradle {
		if strings.Contains(element, "versionName ") {
			fmt.Printf("Set %s to %s\n", strings.TrimSpace(element), strings.TrimSpace(newVersionName))

			newGradle = append(newGradle, element[0:strings.Index(element, "\"")]+newVersionName+"\n")
		} else if strings.Contains(element, "versionCode") {
			fmt.Printf("Set %s to %s\n", strings.TrimSpace(element), strings.TrimSpace(newVersionCode))

			newGradle = append(newGradle, regexp.MustCompile(`\d`).ReplaceAllString(element, newVersionCode)+"\n")
		} else {
			newGradle = append(newGradle, element+"\n")
		}
	}

	util.WriteBuildGradle(newGradle)
}
