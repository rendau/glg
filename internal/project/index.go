package project

import (
	"github.com/rendau/glg/internal/util"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	uriRegexp = regexp.MustCompile(`(?m)^\s*module (.*)$`)
)

func Discover() *St {
	result := &St{}

	result.Uri = getUri()
	result.EntitiesDirPath = getEntitiesDirPath()
	result.DbDirPath = getDbDirPath()

	return result
}

func getUri() string {
	const fName = "go.mod"

	if !util.IsFileExists(fName) {
		log.Fatalf("%s file does not exist, is it go project?", fName)
		return ""
	}

	fData, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalf("Can not read %s file", fName)
		return ""
	}

	fDataStr := strings.TrimSpace(string(fData))

	sm := uriRegexp.FindStringSubmatch(fDataStr)
	if len(sm) < 2 {
		log.Fatalf("Fail to parse %s", fName)
		return ""
	}

	return sm[1]
}

func getEntitiesDirPath() string {
	if path := filepath.Join("internal", "domain", "entities"); util.IsDirExists(path) {
		return path
	}

	return ""
}

func getDbDirPath() string {
	if path := filepath.Join("internal", "adapters", "db", "pg"); util.IsDirExists(path) {
		return path
	}

	return ""
}
