package entity

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

var (
	jsonTagRegexp = regexp.MustCompile(`(?iU)json:"([^"]*)"`)
	glgTagRegexp  = regexp.MustCompile(`(?iU)glg:"([^"]*)"`)
)

func Parse(dirPath string, eName *NameSt) *St {
	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, filepath.Join(dirPath, eName.Snake+".go"), nil, 0)
	if err != nil {
		log.Panicln(err)
	}

	result := ParseF(f, eName)

	FindOutIdField(result)

	return result
}

func ParseF(f *ast.File, eName *NameSt) *St {
	result := &St{
		Name: eName,
	}

	for _, decl := range f.Decls {
		switch gDecl := decl.(type) {
		case *ast.GenDecl:
			if gDecl.Tok == token.TYPE && len(gDecl.Specs) == 1 {
				tSpec := gDecl.Specs[0].(*ast.TypeSpec)
				ParseSt(result, tSpec.Name.Name, eName, tSpec.Type)
			}
		}
	}

	return result
}

func ParseSt(o *St, stName string, eName *NameSt, expr ast.Expr) {
	switch decl := expr.(type) {
	case *ast.StructType:
		nameLower := strings.ToLower(stName)
		eNameLower := strings.ToLower(eName.Camel)

		var stInst *StructSt

		switch {
		case nameLower == eNameLower+"st":
			o.MainSt = &StructSt{}
			stInst = o.MainSt
		case nameLower == eNameLower+"getparsst":
			o.GetParsSt = &StructSt{}
			stInst = o.GetParsSt
		case nameLower == eNameLower+"listst":
			o.ListSt = &StructSt{}
			stInst = o.ListSt
		case nameLower == eNameLower+"listparsst":
			o.ListParsSt = &StructSt{}
			stInst = o.ListParsSt
		case nameLower == eNameLower+"cust":
			o.CuSt = &StructSt{}
			stInst = o.CuSt
		default:
			return
		}

		for _, f := range decl.Fields.List {
			field, isIdField := ParseField(f)
			if field != nil {
				stInst.Fields = append(stInst.Fields, field)

				if stInst == o.MainSt && isIdField {
					o.IdField = field
				}
			}
		}
	}
}

func ParseField(f *ast.Field) (*FieldSt, bool) {
	result := &FieldSt{}

	if len(f.Names) == 1 {
		result.Name.Origin = f.Names[0].Name
		result.Name.Normalize()
	}

	result.Type = ParseType(f.Type)

	if result.Type == "" {
		fmt.Printf("Strange type for field '%s': %v", result.Name.Origin, reflect.TypeOf(f.Type))
	}

	result.IsTypePointer = strings.HasPrefix(result.Type, "*")
	result.IsTypeSlice = strings.HasPrefix(result.Type, "[]") || strings.HasPrefix(result.Type, "*[]")

	var isIdField bool

	if f.Tag != nil {
		result.Tag = f.Tag.Value
		result.JsonName = TagParseJsonName(result.Tag)

		if TagHasGlgId(result.Tag) {
			isIdField = true
		}
	}

	return result, isIdField
}

func ParseType(expr ast.Expr) string {
	switch decl := expr.(type) {
	case *ast.Ident:
		return decl.Name
	case *ast.SelectorExpr:
		var selName string
		if decl.Sel != nil {
			selName = decl.Sel.Name
		}
		if tp := ParseType(decl.X); tp != "" {
			return tp + "." + selName
		}
		return selName
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

func TagParseJsonName(tag string) string {
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

func TagHasGlgId(tag string) bool {
	if fRes := glgTagRegexp.FindStringSubmatch(tag); len(fRes) > 1 {
		for _, w := range strings.Split(fRes[1], ",") {
			if w == "id" {
				return true
			}
		}
	}

	return false
}

func FindOutIdField(o *St) {
	if o.IdField != nil {
		return
	}

	for _, field := range o.MainSt.Fields {
		if field.Name.Snake == "id" {
			o.IdField = field
			break
		}
	}
}
