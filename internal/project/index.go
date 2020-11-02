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

func Discover(dir string) *St {
	result := &St{}

	result.Uri = getUri(dir)
	result.EntitiesDirPath = getEntitiesDirPath(dir)
	result.DbDirPath = getDbDirPath(dir)

	return result
}

func getUri(dir string) string {
	fPath := filepath.Join(dir, "go.mod")

	if !util.IsFileExists(fPath) {
		log.Fatalf("%s file does not exist, is it go project?", fPath)
		return ""
	}

	fData, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Fatalf("Can not read %s file", fPath)
		return ""
	}

	fDataStr := strings.TrimSpace(string(fData))

	sm := uriRegexp.FindStringSubmatch(fDataStr)
	if len(sm) < 2 {
		log.Fatalf("Fail to parse %s", fPath)
		return ""
	}

	return sm[1]
}

func getEntitiesDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "domain", "entities"); util.IsDirExists(path) {
		return &PathSt{
			Abs: "internal/domain/entities",
			Rel: path,
		}
	}

	return nil
}

func getDbDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "adapters", "db", "pg"); util.IsDirExists(path) {
		return &PathSt{
			Abs: "internal/adapters/db/pg",
			Rel: path,
		}
	}

	return nil
}
