package command

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/yagi2/version-bump-up/util"
)

func CmdMinor(c *cli.Context) {
	if (len(c.Args())) != 0 {
		fmt.Fprintf(os.Stderr, "Not enough arguments.\n$ version-bump-up patch")
		os.Exit(1)
	}

	gradle := util.OpenBuildGradle()

	newGradle := make([]string, 0, 100)
	for _, element := range gradle {
		if strings.Contains(element, "versionName ") {
			nowVersionName := util.GetVersionName(element)
			nowMinor := nowVersionName[strings.Index(nowVersionName, ".")+1 : strings.LastIndex(nowVersionName, ".")]

			var nextMinor int
			nextMinor, _ = strconv.Atoi(nowMinor)
			nextMinor++

			nextVersionName := nowVersionName[0:strings.Index(nowVersionName, ".")+1] + strconv.Itoa(nextMinor) + ".0"

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
