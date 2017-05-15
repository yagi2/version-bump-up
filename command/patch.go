package command

import (
	"fmt"
	"strconv"
	"strings"

	"os"

	"github.com/codegangsta/cli"
	"github.com/yagi2/version-bump-up/util"
)

func CmdPatch(c *cli.Context) {
	if (len(c.Args())) != 0 {
		fmt.Fprintf(os.Stderr, "Not enough arguments.\n$ version-bump-up patch")
		os.Exit(1)
	}

	gradle := util.OpenBuildGradle()

	newGradle := make([]string, 0, 100)
	for _, element := range gradle {
		if strings.Contains(element, "versionName ") {
			nowVersionName := util.GetVersionName(element)
			nowPatch := nowVersionName[strings.LastIndex(nowVersionName, ".")+1 : len(nowVersionName)]

			var nextPatch int
			nextPatch, _ = strconv.Atoi(nowPatch)
			nextPatch++

			nextVersionName := nowVersionName[0:strings.LastIndex(nowVersionName, ".")+1] + strconv.Itoa(nextPatch)

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
