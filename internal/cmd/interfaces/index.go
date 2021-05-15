package interfaces

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/rendau/glg/assets"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	// var err error

	if pr.InterfacesDirPath == nil {
		fmt.Println("Interfaces dir not found")
		return
	}

	fPath := filepath.Join(pr.InterfacesDirPath.Abs, "db.go")

	if !util.IsFileExists(fPath) {
		fmt.Println("Db-interface file not found")
		return
	}

	removeCurrentMethods(fPath, eName)

	side1, side2 := getInjectPosSides(fPath)

	tData, err := assets.Asset("templates/interfaces.tmpl")
	if err != nil {
		log.Panicln(err)
	}

	t, err := template.New("interfaces.tmp").Parse(string(tData))
	if err != nil {
		log.Panicln(err)
	}

	tResultBuffer := &bytes.Buffer{}

	err = t.Execute(tResultBuffer, struct {
		Prefix   string
		Suffix   string
		Pr       *project.St
		EName    *entity.NameSt
		Ent      *entity.St
		Ctx4List map[string]interface{}
	}{
		Prefix:   side1,
		Suffix:   side2,
		Pr:       pr,
		EName:    eName,
		Ent:      ent,
		Ctx4List: getCtx4List(pr, eName, ent),
	})
	if err != nil {
		log.Panicln(err)
	}

	err = ioutil.WriteFile(fPath, []byte(side1+"\n"+tResultBuffer.String()+side2), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}

func removeCurrentMethods(fPath string, eName *entity.NameSt) {
	var re = regexp.MustCompile(`(?si)(?://\s*` + eName.Snake + `\n\s*)?` + eName.Camel + `(?:Get|List|IdExists|Create|Update|Delete)\([^\n]+\n`)

	fDataRaw, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Panicln(err)
	}

	newData := re.ReplaceAllString(string(fDataRaw), "")

	err = ioutil.WriteFile(fPath, []byte(newData), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}

func getInjectPosSides(fPath string) (string, string) {
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
				if strings.Contains(strings.ToLower(tSpec.Name.Name), "db") {
					switch decl := tSpec.Type.(type) {
					case *ast.InterfaceType:
						return fData[:decl.Methods.Closing-1], fData[decl.Methods.Closing-1:]
					}
				}
			}
		}
	}

	return "", ""
}

func getCtx4List(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	if ent.ListParsSt != nil {
		for _, field := range ent.ListParsSt.Fields {
			if strings.Contains(strings.ToLower(field.Type), "pagination") {
				result["hasPagination"] = true
				break
			}
		}
	}

	return result
}
