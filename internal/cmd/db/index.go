package db

import (
	"fmt"
	"github.com/rendau/glg/assets"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.DbDirPath == nil {
		fmt.Println("Db destination dir not found")
		return
	}

	tData, err := assets.Asset("templates/db.tmpl")
	if err != nil {
		log.Panicln(err)
	}

	t, err := template.New("db.tmp").Funcs(template.FuncMap{
		"notLastI":           func(i, len int) bool { return (i + 1) < len },
		"fieldPgType":        fieldPgType,
		"fieldsPgTypeFilter": fieldsPgTypeFilter,
	}).Parse(string(tData))
	if err != nil {
		log.Panicln(err)
	}

	outF, err := os.Create(filepath.Join(pr.DbDirPath.Rel, eName.Snake+".go"))
	if err != nil {
		log.Panicln(err)
	}
	defer outF.Close()

	err = t.Execute(outF, struct {
		Pr      *project.St
		EName   *entity.NameSt
		Ent     *entity.St
		Ctx4Get map[string]interface{}
	}{
		Pr:      pr,
		EName:   eName,
		Ent:     ent,
		Ctx4Get: getCtx4Get(pr, eName, ent),
	})
	if err != nil {
		log.Panicln(err)
	}
}

func getCtx4Get(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	result["scanableFields"] = scanableFields(ent.MainSt.Fields)

	return result
}

func scanableFields(fields []*entity.FieldSt) []*entity.FieldSt {
	result := make([]*entity.FieldSt, 0)

	for _, f := range fields {
		switch f.Type {
		case "bool", "string":
			result = append(result, f)
		case "int", "int8", "int16", "int32", "int64":
			result = append(result, f)
		case "uint", "uint8", "uint16", "uint32", "uint64":
			result = append(result, f)

		case "[]bool", "[]string":
			result = append(result, f)
		case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
			result = append(result, f)
		case "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
			result = append(result, f)
		}
	}

	return result
}

func fieldPgType(field *entity.FieldSt) string {
	switch field.Type {
	case "[]bool":
		return "pgtype.BoolArray"
	case "[]string":
		return "pgtype.TextArray"
	case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
		return "pgtype.Int8Array"
	case "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
		return "pgtype.Int8Array"
	}
	return ""
}

func fieldsPgTypeFilter(fields []*entity.FieldSt) []*entity.FieldSt {
	result := make([]*entity.FieldSt, 0)

	for _, field := range fields {
		if fieldPgType(field) != "" {
			result = append(result, field)
		}
	}

	return result
}
