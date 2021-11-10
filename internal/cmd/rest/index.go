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
	"strings"
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
		"getQueryParParser": getQueryParParser,
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
		Ctx4List map[string]interface{}
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
	const fName = "router.go"

	fPath := filepath.Join(pr.RestDirPath.Abs, fName)

	// remove current routes

	fDataRaw, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Panicln(err)
	}

	fData := string(fDataRaw)

	re := regexp.MustCompile(`(?si)(?://\s*` + eName.Snake + `\n\s*)?r.Handle(?:Func)?\("/` + eName.Snake + `["/][^\n]+\n`)
	fData = re.ReplaceAllString(fData, "")

	err = ioutil.WriteFile(fPath, []byte(fData), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)

	// register

	side1, side2, ok := util.DivideFuncReturnPosSides(fPath, "router")
	if !ok || side1 == "" || side2 == "" {
		fmt.Println("Fail to register routes in rest. Not found 'router' function in `" + fName + "` file")
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
		Ctx4List       map[string]interface{}
		IdPathParRegex string
		WithMetrics    bool
	}{
		Pr:             pr,
		EName:          eName,
		Ent:            ent,
		Ctx4List:       getCtx4List(pr, eName, ent),
		IdPathParRegex: idPathParRegex,
		WithMetrics:    strings.Contains(side1, "mh := func("),
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

func getCtx4List(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	if ent.ListParsSt != nil {
		for _, field := range ent.ListParsSt.Fields {
			if strings.Contains(strings.ToLower(field.Type), "pagination") {
				result["hasPagination"] = true
				break
			}
		}

		result["parsFields"] = ent.ListParsSt.Fields
	}

	return result
}

func getQueryParParser(field *entity.FieldSt) string {
	switch field.Type {
	case "*bool":
		return "uQpParseBool"
	case "*int":
		return "uQpParseInt"
	case "*int64":
		return "uQpParseInt64"
	case "*float64":
		return "uQpParseFloat64"
	case "*time.Time":
		return "uQpParseTime"
	case "*[]int64":
		return "uQpParseInt64Slice"
	case "*string":
		return "uQpParseString"
	case "*[]string":
		return "uQpParseStringSlice"

	case "bool":
		return "uQpParseBoolV"
	case "int":
		return "uQpParseIntV"
	case "int64":
		return "uQpParseInt64V"
	case "float64":
		return "uQpParseFloat64V"
	case "[]int64":
		return "uQpParseInt64SliceV"
	case "string":
		return "uQpParseStringV"
	case "[]string":
		return "uQpParseStringSliceV"
	}

	return ""
}
