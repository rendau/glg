package db

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

//go:embed tmpl.tmpl
var tmp string

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.DbDirPath == nil {
		fmt.Println("Db destination dir not found")
		return
	}

	t, err := template.New("db.tmp").Funcs(template.FuncMap{
		"parsFieldAssocName":   parsFieldAssocName,
		"fieldSubQueryForIn":   fieldSubQueryForIn,
		"fieldPgTypeNullValue": fieldPgTypeNullValue,
	}).Parse(tmp)
	if err != nil {
		log.Panicln(err)
	}

	fPath := filepath.Join(pr.DbDirPath.Abs, eName.Snake+".go")

	outF, err := os.Create(fPath)
	if err != nil {
		log.Panicln(err)
	}
	defer outF.Close()

	err = t.Execute(outF, struct {
		Pr         *project.St
		EName      *entity.NameSt
		TName      string
		Ent        *entity.St
		Ctx4Get    map[string]interface{}
		Ctx4List   map[string]interface{}
		Ctx4CuArgs map[string]interface{}
	}{
		Pr:         pr,
		EName:      eName,
		TName:      tableName(eName),
		Ent:        ent,
		Ctx4Get:    getCtx4Get(pr, eName, ent),
		Ctx4List:   getCtx4List(pr, eName, ent),
		Ctx4CuArgs: getCtx4CuArgs(pr, eName, ent),
	})
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}

func getCtx4Get(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	result["scanableFields"] = scanableFields(ent.MainSt.Fields)

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

		for _, field := range ent.ListParsSt.Fields {
			if field.Name.Snake == "only_count" && field.Type == "bool" {
				result["onlyCountFieldName"] = field.Name.Origin
				break
			}
		}

		result["parsFields"] = ent.ListParsSt.Fields
	}

	result["fields"] = ent.ListSt.Fields
	result["scanableFields"] = scanableFields(ent.ListSt.Fields)

	return result
}

func getCtx4CuArgs(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	result["fields"] = ent.CuSt.Fields

	return result
}

func scanableFields(fields []*entity.FieldSt) []*entity.FieldSt {
	result := make([]*entity.FieldSt, 0)

	for _, f := range fields {
		switch f.Type {
		case "bool", "string",
			"*bool", "*string":
			result = append(result, f)
		case "int", "int8", "int16", "int32", "int64",
			"*int", "*int8", "*int16", "*int32", "*int64":
			result = append(result, f)
		case "uint", "uint8", "uint16", "uint32", "uint64",
			"*uint", "*uint8", "*uint16", "*uint32", "*uint64":
			result = append(result, f)
		case "float32", "float64",
			"*float32", "*float64":
			result = append(result, f)
		case "time.Time",
			"*time.Time":
			result = append(result, f)

		case "[]bool", "[]string",
			"*[]bool", "*[]string":
			result = append(result, f)
		case "[]int", "[]int8", "[]int16", "[]int32", "[]int64",
			"*[]int", "*[]int8", "*[]int16", "*[]int32", "*[]int64":
			result = append(result, f)
		case "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64",
			"*[]uint", "*[]uint8", "*[]uint16", "*[]uint32", "*[]uint64":
			result = append(result, f)
		case "[]float32", "[]float64",
			"*[]float32", "*[]float64":
			result = append(result, f)
		}
	}

	return result
}

func parsFieldAssocName(ent *entity.St, field *entity.FieldSt) string {
	if strings.ToLower(field.Name.Snake) == "ids" {
		return "id"
	}

	if ent.MainSt != nil {
		for _, f := range ent.MainSt.Fields {
			if f.Name.Camel == field.Name.Camel {
				return f.JsonName
			}
		}
	}

	if ent.ListSt != nil {
		for _, f := range ent.ListSt.Fields {
			if f.Name.Camel == field.Name.Camel {
				return f.JsonName
			}
		}
	}

	return ""
}

func fieldSubQueryForIn(field *entity.FieldSt, name string) string {
	switch field.Type {
	case "[]string", "*[]string":
		return `(select * from unnest(${` + name + `} :: string[]))`
	case "[]int", "[]int8", "[]int16", "[]int32", "[]int64",
		"[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64",
		"*[]int", "*[]int8", "*[]int16", "*[]int32", "*[]int64",
		"*[]uint", "*[]uint8", "*[]uint16", "*[]uint32", "*[]uint64":
		return `(select * from unnest(${` + name + `} :: bigint[]))`
	}
	return ""
}

func tableName(eName *entity.NameSt) string {
	switch eName.Snake {
	case "role":
		return `"role"`
	case "group":
		return `"group"`
	case "user":
		return `"user"`
	}

	return eName.Snake
}

func fieldPgTypeNullValue(field *entity.FieldSt) string {
	switch field.Type {
	case "*int", "*int8", "*int16",
		"*uint", "*uint8", "*uint16":
		return `pgtype.Int2{Status: pgtype.Null}`
	case "*int32", "*uint32":
		return `pgtype.Int4{Status: pgtype.Null}`
	case "*int64", "*uint64":
		return `pgtype.Int8{Status: pgtype.Null}`
	case "*float32":
		return `pgtype.Float4{Status: pgtype.Null}`
	case "*float64":
		return `pgtype.Float8{Status: pgtype.Null}`
	}
	return ""
}
