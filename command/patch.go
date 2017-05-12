package command

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/yagi2/version-bump-up/util"
)

func CmdPatch(c *cli.Context) {
	gradle := util.OpenBuildGradle()

	// debug
	for _, element := range gradle {
		if strings.Contains(element, "versionName ") || strings.Contains(element, "versionCode ") {
			fmt.Println(element)
		}
	}
}
