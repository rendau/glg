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

	SyncNullableFields(result)

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
			field, isIdField, isNullableField := ParseField(f)
			if field != nil {
				stInst.Fields = append(stInst.Fields, field)

				if stInst == o.MainSt && isIdField {
					o.IdField = field
					stInst.IdField = field
				}

				if stInst == o.MainSt && (isNullableField || field.IsTypePointer) {
					field.IsNullable = true
				}
			}
		}
	}
}

func ParseField(f *ast.Field) (*FieldSt, bool, bool) {
	result := &FieldSt{}

	if len(f.Names) == 1 {
		result.Name.Origin = f.Names[0].Name
		result.Name.Normalize(false)
	}

	result.Type = ParseType(f.Type)

	if result.Type == "" {
		fmt.Printf("Strange type for field '%s': %v", result.Name.Origin, reflect.TypeOf(f.Type))
	}

	result.IsTypePointer = strings.HasPrefix(result.Type, "*")
	result.IsTypeSlice = strings.HasPrefix(result.Type, "[]") || strings.HasPrefix(result.Type, "*[]")
	result.IsTypeInt = strings.HasPrefix(result.Type, "int")

	var isIdField bool
	var isNullableField bool

	if f.Tag != nil {
		result.Tag = f.Tag.Value
		result.JsonName = TagParseJsonName(result.Tag)

		if TagHasGlgId(result.Tag) {
			isIdField = true
		}
		if TagHasGlgNullable(result.Tag) {
			isNullableField = true
		}
	}

	result.DefineZeroValue()

	result.DefinePVZeroValue()

	return result, isIdField, isNullableField
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

func TagHasGlgNullable(tag string) bool {
	if fRes := glgTagRegexp.FindStringSubmatch(tag); len(fRes) > 1 {
		for _, w := range strings.Split(fRes[1], ",") {
			if w == "nullable" {
				return true
			}
		}
	}

	return false
}

func FindOutIdField(o *St) {
	if o.IdField == nil {
		for _, field := range o.MainSt.Fields {
			if field.Name.Snake == "id" {
				field.IsId = true
				o.MainSt.IdField = field
				o.IdField = field
			}
		}
	}

	// set IsId flag on other struct fields
	if o.IdField != nil {
		for _, st := range []*StructSt{o.GetParsSt, o.ListSt, o.ListParsSt} {
			if st == nil {
				continue
			}

			for _, field := range st.Fields {
				if field.Name.Origin == o.IdField.Name.Origin &&
					(field.Type == o.IdField.Type || field.Type == ("*"+o.IdField.Type)) {
					field.IsId = true
					st.IdField = field
				}
			}
		}
	}
}

func SyncNullableFields(o *St) {
	for _, field := range o.MainSt.Fields {
		if !field.IsNullable {
			continue
		}

		for _, st := range []*StructSt{o.GetParsSt, o.ListSt, o.ListParsSt, o.CuSt} {
			if st == nil {
				continue
			}

			for _, f := range st.Fields {
				if f.Name.Origin == field.Name.Origin {
					f.IsNullable = true
				}
			}
		}
	}
}
