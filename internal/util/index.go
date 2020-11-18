package util

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	camelCaseRegexp = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

func IsFileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return !info.IsDir()
}

func IsDirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return info.IsDir()
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func Case2Camel(v string) string {
	result := ""

	for _, w := range strings.Split(v, "_") {
		if w == "" {
			continue
		}
		result += strings.Title(w)
	}

	return result
}

func Case2Snake(v string) string {
	return strings.ToLower(camelCaseRegexp.ReplaceAllString(v, "${1}_${2}"))
}

func FmtFile(fPath string) {
	err := exec.Command("goimports", "-w", fPath).Run()
	if err != nil {
		fmt.Println("Fail to 'goimports'", fPath, err)

		err = exec.Command("gofmt", "-w", fPath).Run()
		if err != nil {
			fmt.Println("Fail to 'gofmt'", fPath, err)
		}
	}
}
