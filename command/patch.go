package command

import (
	"fmt"
	"regexp"
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
			nowPatch := nowVersionName[strings.LastIndex(nowVersionName, ".")+1 : strings.LastIndex(nowVersionName, ".")+2]

			var nextPatch int
			nextPatch, _ = strconv.Atoi(nowPatch)
			nextPatch++

			nextVersionName := nowVersionName[0:strings.LastIndex(nowVersionName, ".")+1] + strconv.Itoa(nextPatch)

			fmt.Printf("Bump up Patch %s to %s\n", nowVersionName, nextVersionName)

			newGradle = append(newGradle, element[0:strings.Index(element, "\"")]+"\""+nextVersionName+"\"\n")
		} else if strings.Contains(element, "versionCode ") {
			nowVersionCode := util.GetVersionCode(element)

			var nextVersionCode int
			nextVersionCode, _ = strconv.Atoi(nowVersionCode)
			nextVersionCode++

			fmt.Printf("Bump up Patch %s to %s\n", nowVersionCode, strconv.Itoa(nextVersionCode))

			newGradle = append(newGradle, regexp.MustCompile(`\d`).ReplaceAllString(element, strconv.Itoa(nextVersionCode))+"\n")
		} else {
			newGradle = append(newGradle, element+"\n")
		}
	}

	util.WriteBuildGradle(newGradle)
}
