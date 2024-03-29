package repo_pg

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

	if pr.RepoPgDirPath == nil {
		fmt.Println("'repo/pg' destination dir not found")
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

	fPath := filepath.Join(pr.RepoPgDirPath.Abs, eName.Snake+".go")

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
		Ctx4Get    map[string]any
		Ctx4List   map[string]any
		Ctx4CuArgs map[string]any
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

func getCtx4Get(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]any {
	result := map[string]any{}

	result["scanableFields"] = scanableFields(ent.MainSt.Fields)

	return result
}

func getCtx4List(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]any {
	result := map[string]any{}

	if ent.ListParsSt != nil {
		result["parsFields"] = ent.ListParsSt.Fields
	}

	result["fields"] = ent.ListSt.Fields
	result["scanableFields"] = scanableFields(ent.ListSt.Fields)

	return result
}

func getCtx4CuArgs(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]any {
	result := map[string]any{}

	result["fields"] = ent.CuSt.Fields

	return result
}

func scanableFields(fields []*entity.FieldSt) []*entity.FieldSt {
	return fields
}

func parsFieldAssocName(ent *entity.St, field *entity.FieldSt) string {
	if strings.ToLower(field.Name.Snake) == "ids" {
		return "id"
	}

	if ent.MainSt != nil {
		for _, f := range ent.MainSt.Fields {
			if f.Name.Camel == field.Name.Camel || (f.Name.Camel+"s") == field.Name.Camel {
				return f.DbName
			}
		}
	}

	if ent.ListSt != nil {
		for _, f := range ent.ListSt.Fields {
			if f.Name.Camel == field.Name.Camel || (f.Name.Camel+"s") == field.Name.Camel {
				return f.DbName
			}
		}
	}

	return ""
}

func fieldSubQueryForIn(field *entity.FieldSt, name string) string {
	switch field.Type {
	case "[]string", "*[]string":
		return `(select * from unnest(${` + name + `} :: text[]))`
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
