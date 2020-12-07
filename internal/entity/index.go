package entity

import (
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

func Parse(dirPath string, eName *NameSt) *St {
	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, filepath.Join(dirPath, eName.Snake+".go"), nil, 0)
	if err != nil {
		log.Panicln(err)
	}

	return ParseF(f, eName)
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
			field := ParseField(f)
			if field != nil {
				stInst.Fields = append(stInst.Fields, field)
			}
		}
	}
}

func ParseField(f *ast.Field) *FieldSt {
	result := &FieldSt{}

	if len(f.Names) == 1 {
		result.Name.Origin = f.Names[0].Name
		result.Name.Normalize()
	}

	result.Type = ParseType(f.Type)

	result.IsTypePointer = strings.HasPrefix(result.Type, "*")
	result.IsTypeSlice = strings.HasPrefix(result.Type, "[]") || strings.HasPrefix(result.Type, "*[]")

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
