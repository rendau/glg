package module

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"regexp"
	"strings"
)

var (
	jsonTagRegexp = regexp.MustCompile(`(?iU)json:"(.*)"`)
)

func Run(fName, mName string) {
	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, fName, nil, 0)
	if err != nil {
		log.Panicln("Fail to parse file", err)
	}

	ns := ParseF(f, mName)

	fmt.Println(ns.Name)
	if ns.MainSt != nil {
		fmt.Println("MainSt:")
		for _, f := range ns.MainSt.Fields {
			fmt.Println("  ", *f)
		}
	}
	if ns.CuSt != nil {
		fmt.Println("CuSt:")
		for _, f := range ns.CuSt.Fields {
			fmt.Println("  ", *f)
		}
	}
}

func ParseF(f *ast.File, mName string) *NsSt {
	result := &NsSt{
		Name: mName,
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

func ParseSt(ns *NsSt, name string, expr ast.Expr) {
	switch decl := expr.(type) {
	case *ast.StructType:
		nameLower := strings.ToLower(name)
		mNameLower := strings.ToLower(ns.Name)

		var stInst *StructSt

		switch {
		case nameLower == mNameLower+"st":
			ns.MainSt = &StructSt{}
			stInst = ns.MainSt
		case nameLower == mNameLower+"cust":
			ns.CuSt = &StructSt{}
			stInst = ns.CuSt
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
