package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/yagi2/version-bump-up/util"
)

func CmdMajor(c *cli.Context) {
	if (len(c.Args())) != 0 {
		fmt.Fprintf(os.Stderr, "Not enough arguments.\n$ version-bump-up patch")
		os.Exit(1)
	}

	gradle := util.OpenBuildGradle()

	newGradle := make([]string, 0, 100)
	for _, element := range gradle {
		if strings.Contains(element, "versionName ") {
			nowVersionName := util.GetVersionName(element)
			nowMajor := nowVersionName[0:strings.Index(nowVersionName, ".")]

			var nextMajor int
			nextMajor, _ = strconv.Atoi(nowMajor)
			nextMajor++

			nextVersionName := strconv.Itoa(nextMajor) + ".0.0"

			fmt.Printf("Bump up Patch %s to %s\n", nowVersionName, nextVersionName)

			newGradle = append(newGradle, element[0:strings.Index(element, "\"")]+"\""+nextVersionName+"\"\n")
		} else if strings.Contains(element, "versionCode ") {
			var nextVersionCode int
			nextVersionCode, _ = strconv.Atoi(util.GetVersionCode(element))
			nextVersionCode++

			fmt.Printf("Bump up Patch %s to %s\n", util.GetVersionCode(element), strconv.Itoa(nextVersionCode))

			newGradle = append(newGradle, util.GetNextVersionCode(element))
		} else {
			newGradle = append(newGradle, element+"\n")
		}
	}

	util.WriteBuildGradle(newGradle)
}
