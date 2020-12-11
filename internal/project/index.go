package project

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rendau/glg/internal/util"
)

var (
	uriRegexp = regexp.MustCompile(`(?m)^\s*module (.*)$`)
)

func Discover(dir string) *St {
	result := &St{}

	result.Uri = getUri(dir)
	result.EntitiesDirPath = getEntitiesDirPath(dir)
	result.DbDirPath = getDbDirPath(dir)
	result.InterfacesDirPath = getInterfacesDirPath(dir)
	result.CoreDirPath = getCoreDirPath(dir)
	result.UsecasesDirPath = getUsecasesDirPath(dir)
	result.RestDirPath = getRestDirPath(dir)

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
			Abs: path,
			Rel: "internal/domain/entities",
		}
	}

	return nil
}

func getDbDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "adapters", "db", "pg"); util.IsDirExists(path) {
		return &PathSt{
			Abs: path,
			Rel: "internal/adapters/db/pg",
		}
	}

	return nil
}

func getInterfacesDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "interfaces"); util.IsDirExists(path) {
		return &PathSt{
			Abs: path,
			Rel: "internal/interfaces",
		}
	}

	return nil
}

func getCoreDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "domain", "core"); util.IsDirExists(path) {
		return &PathSt{
			Abs: path,
			Rel: "internal/domain/core",
		}
	}

	return nil
}

func getUsecasesDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "domain", "usecases"); util.IsDirExists(path) {
		return &PathSt{
			Abs: path,
			Rel: "internal/domain/usecases",
		}
	}

	return nil
}

func getRestDirPath(dir string) *PathSt {
	if path := filepath.Join(dir, "internal", "adapters", "httpapi", "rest"); util.IsDirExists(path) {
		return &PathSt{
			Abs: path,
			Rel: "internal/adapters/httpapi/rest",
		}
	}

	return nil
}
