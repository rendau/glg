package rest

import (
	"bytes"
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

//go:embed tmplHandlers.tmpl
var tmpHandlers string

//go:embed tmplRouter.tmpl
var tmpRouter string

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.RestDirPath == nil {
		fmt.Println("REST destination dir not found")
		return
	}

	t, err := template.New("rest.tmp").Funcs(template.FuncMap{
		"getTypeForSwag": getTypeForSwag,
	}).Parse(tmpHandlers)
	if err != nil {
		log.Panicln(err)
	}

	fPath := filepath.Join(pr.RestDirPath.Abs, "h_"+eName.Snake+".go")

	outF, err := os.Create(fPath)
	if err != nil {
		log.Panicln(err)
	}
	defer outF.Close()

	err = t.Execute(outF, struct {
		Pr       *project.St
		EName    *entity.NameSt
		Ent      *entity.St
		Ctx4List map[string]any
	}{
		Pr:       pr,
		EName:    eName,
		Ent:      ent,
		Ctx4List: getCtx4List(pr, eName, ent),
	})
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)

	registerRoutes(pr, eName, ent)
}

func registerRoutes(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	const fName = "index.go"

	fPath := filepath.Join(pr.RestDirPath.Abs, fName)

	// remove current routes

	fDataRaw, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Panicln(err)
	}

	fData := string(fDataRaw)

	re := regexp.MustCompile(`(?si)(?://\s+` + eName.Snake + `\n\s+)?r\.[^\(]+\("/` + eName.Snake + `["/][^\n]+\n`)
	fData = re.ReplaceAllString(fData, "")

	err = ioutil.WriteFile(fPath, []byte(fData), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)

	// register

	side1, side2, ok := util.DivideFuncReturnPosSides(fPath, "GetHandler")
	if !ok || side1 == "" || side2 == "" {
		fmt.Println("Fail to register routes in rest. Not found 'GetHandler' function in `" + fName + "` file")
	}

	t, err := template.New("rest_router.tmp").Parse(tmpRouter)
	if err != nil {
		log.Panicln(err)
	}

	tResultBuffer := &bytes.Buffer{}

	var idPathParRegex string

	if ent.IdField != nil {
		if ent.IdField.IsTypeInt {
			idPathParRegex = `[0-9]+`
		} else {
			idPathParRegex = `[^/]+`
		}
	}

	err = t.Execute(tResultBuffer, struct {
		Pr             *project.St
		EName          *entity.NameSt
		Ent            *entity.St
		Ctx4List       map[string]any
		IdPathParRegex string
	}{
		Pr:             pr,
		EName:          eName,
		Ent:            ent,
		Ctx4List:       getCtx4List(pr, eName, ent),
		IdPathParRegex: idPathParRegex,
	})
	if err != nil {
		log.Panicln(err)
	}

	err = ioutil.WriteFile(fPath, []byte(side1+"\n\n"+tResultBuffer.String()+side2), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}

func getCtx4List(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]any {
	result := map[string]any{}

	if ent.ListParsSt != nil {
		result["parsFields"] = ent.ListParsSt.Fields
	}

	return result
}

func getTypeForSwag(field *entity.FieldSt) string {
	switch field.Type {
	case "bool":
		return `boolean`
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return `integer`
	case "float32", "float64":
		return `number`
	default:
		return "string"
	}
}
