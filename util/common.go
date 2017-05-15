package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func OpenBuildGradle() []string {
	file, err := os.Open(`./app/build.gradle`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open app/build.gradle")
		os.Exit(1)
	}

	defer file.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "Can't read app/build.gradle")
	}

	return lines
}

func WriteBuildGradle(newGradle []string) {
	file, err := os.Create(`./app/build.gradle`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open app/build.gradle")
		os.Exit(1)
	}

	defer file.Close()
	for _, element := range newGradle {
		file.Write(([]byte)(element))
	}

	fmt.Println("Complete write to app/build.gradle")
}

func GetVersionName(element string) string {
	return element[strings.Index(element, "\"")+1 : len(element)-1]
}

func GetVersionCode(element string) string {
	return regexp.MustCompile(`\d`).FindAllStringSubmatch(element, -1)[0][0]
}
