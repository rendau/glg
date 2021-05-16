package rest

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rendau/glg/assets"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.RestDirPath == nil {
		fmt.Println("REST destination dir not found")
		return
	}

	tData, err := assets.Asset("templates/rest_h.tmpl")
	if err != nil {
		log.Panicln(err)
	}

	t, err := template.New("rest.tmp").Funcs(template.FuncMap{
		"getListParParser": getListParParser,
	}).Parse(string(tData))
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
		Ctx4Get  map[string]interface{}
		Ctx4List map[string]interface{}
	}{
		Pr:       pr,
		EName:    eName,
		Ent:      ent,
		Ctx4Get:  getCtx4Get(pr, eName, ent),
		Ctx4List: getCtx4List(pr, eName, ent),
	})
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}

func getCtx4Get(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	if ent.IdField != nil {
		if ent.GetParsSt != nil {
			for _, field := range ent.GetParsSt.Fields {
				if field.Name.Origin == ent.IdField.Name.Origin {
					result["idFieldInGetParsSt"] = field
					break
				}
			}
		}
	}

	return result
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

func getListParParser(field *entity.FieldSt) string {
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
	}

	return ""
}
