package entity

import (
	"github.com/rendau/glg/internal/util"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	jsonTagRegexp = regexp.MustCompile(`(?iU)json:"(.*)"`)
)

func Parse(dirPath, name string) *St {
	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, filepath.Join(dirPath, name+".go"), nil, 0)
	if err != nil {
		log.Panicln(err)
	}

	return ParseF(f, name)
}

func ParseF(f *ast.File, name string) *St {
	result := &St{
		Name: NameSt{
			Camel: util.CaseSnake2Camel(name),
			Snake: name,
		},
	}

	for _, decl := range f.Decls {
		switch gDecl := decl.(type) {
		case *ast.GenDecl:
			if gDecl.Tok == token.TYPE && len(gDecl.Specs) == 1 {
				tSpec := gDecl.Specs[0].(*ast.TypeSpec)
				ParseSt(result, tSpec.Name.Name, tSpec.Type)
			}
		}
	}

	return result
}

func ParseSt(o *St, stName string, expr ast.Expr) {
	switch decl := expr.(type) {
	case *ast.StructType:
		nameLower := strings.ToLower(stName)
		eNameLower := strings.ToLower(o.Name.Camel)

		var stInst *StructSt

		switch {
		case nameLower == eNameLower+"st":
			o.MainSt = &StructSt{}
			stInst = o.MainSt
		case nameLower == eNameLower+"cust":
			o.CuSt = &StructSt{}
			stInst = o.CuSt
		default:
			return
		}

		for _, f := range decl.Fields.List {
			field := ParseField(f)
			if field != nil {
				stInst.Fields = append(stInst.Fields, field)
			}
		}
	}
}

func ParseField(f *ast.Field) *FieldSt {
	if len(f.Names) != 1 {
		return nil
	}

	result := &FieldSt{
		Name: f.Names[0].Name,
	}

	result.Type = ParseType(f.Type)

	if f.Tag != nil {
		result.Tag = f.Tag.Value
		result.JsonName = ParseTagJson(result.Tag)
	}

	return result
}

func ParseType(expr ast.Expr) string {
	switch decl := expr.(type) {
	case *ast.Ident:
		return decl.Name
	case *ast.StarExpr:
		if tp := ParseType(decl.X); tp != "" {
			return "*" + tp
		}
	case *ast.ArrayType:
		if tp := ParseType(decl.Elt); tp != "" {
			return "[]" + tp
		}
	}

	return ""
}

func ParseTagJson(tag string) string {
	if fRes := jsonTagRegexp.FindStringSubmatch(tag); len(fRes) > 1 {
		for _, w := range strings.Split(fRes[1], ",") {
			if w == "" || w == "-" || w == "omitempty" {
				continue
			}
			return w
		}
	}
	return ""
}
