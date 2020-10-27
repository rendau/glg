package module

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func Run(fName, mName string) error {
	fSet := token.NewFileSet()

	f, err := parser.ParseFile(fSet, fName, nil, 0)
	if err != nil {
		log.Panicln("Fail to parse file", err)
	}

	ns, err := ParseF(f, mName)
	if err != nil {
		log.Panicln("Fail to parse file", err)
	}

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

	return nil
}

func ParseF(f *ast.File, mName string) (*NsSt, error) {
	result := &NsSt{
		Name: mName,
	}

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			if decl.Tok != token.TYPE || len(decl.Specs) != 1 {
				continue
			}
			tSpec := decl.Specs[0].(*ast.TypeSpec)
			switch decl := tSpec.Type.(type) {
			case *ast.StructType:
				err := ParseSt(result, tSpec.Name.Name, decl)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return result, nil
}

func ParseSt(ns *NsSt, name string, s *ast.StructType) error {
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
		return nil
	}

	for _, f := range s.Fields.List {
		field, err := ParseFld(f)
		if err != nil {
			return err
		}
		if field != nil {
			stInst.Fields = append(stInst.Fields, field)
		}
	}

	return nil
}

func ParseFld(f *ast.Field) (*FieldSt, error) {
	var err error

	if len(f.Names) != 1 {
		return nil, nil
	}

	result := &FieldSt{
		Name: f.Names[0].Name,
	}

	result.Type, err = ParseType(f.Type)
	if err != nil {
		return nil, err
	}

	if f.Tag != nil {
		result.Tag = f.Tag.Value
	}

	return result, nil
}

func ParseType(expr ast.Expr) (string, error) {
	switch decl := expr.(type) {
	case *ast.Ident:
		return decl.Name, nil
	case *ast.StarExpr:
		tp, err := ParseType(decl.X)
		if err != nil {
			return "", err
		}
		return "*" + tp, nil
	case *ast.ArrayType:
		tp, err := ParseType(decl.Elt)
		if err != nil {
			return "", err
		}
		return "[]" + tp, nil
	}

	return "", nil
}
