package util

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

func DivideStructEndPosSides(fPath, structName string) (string, string, bool) {
	fDataRaw, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Panicln(err)
	}

	fData := string(fDataRaw)

	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, filepath.Join(fPath), nil, 0)
	if err != nil {
		log.Panicln(err)
	}

	for _, decl := range f.Decls {
		switch gDecl := decl.(type) {
		case *ast.GenDecl:
			if gDecl.Tok == token.TYPE && len(gDecl.Specs) == 1 {
				tSpec := gDecl.Specs[0].(*ast.TypeSpec)
				if strings.ToLower(tSpec.Name.Name) == strings.ToLower(structName) {
					switch decl := tSpec.Type.(type) {
					case *ast.StructType:
						return fData[:decl.Fields.Closing-1], fData[decl.Fields.Closing-1:], true
					}
				}
			}
		}
	}

	return "", "", false
}

func DivideFuncReturnPosSides(fPath, funcName string) (string, string, bool) {
	fDataRaw, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Panicln(err)
	}

	fData := string(fDataRaw)

	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, filepath.Join(fPath), nil, 0)
	if err != nil {
		log.Panicln(err)
	}

	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			if strings.ToLower(fn.Name.Name) == strings.ToLower(funcName) {
				returnPos := strings.LastIndex(fData[:fn.Body.End()-1], "return")
				if returnPos > 0 {
					return fData[:returnPos-1], fData[returnPos-1:], true
				}
			}
		}
	}

	return "", "", false
}
