package util

import (
	"os"
	"regexp"
	"strings"
)

var (
	camelCaseRegexp = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

func IsFileExists(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && !info.IsDir()
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CaseSnake2Camel(v string) string {
	result := ""

	for _, w := range strings.Split(v, "_") {
		if w == "" {
			continue
		}
		result += strings.Title(w)
	}

	return result
}

func CaseCamel2Snake(v string) string {
	return strings.ToLower(camelCaseRegexp.ReplaceAllString(v, "${1}_${2}"))
}