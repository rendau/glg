package module

import (
	"fmt"
	"github.com/rendau/glg/internal/util"
	"path/filepath"
)

func Run(name string) {
	fPath := filepath.Join("internal", "domain", "entities", name+".go")
	if !util.IsFileExists(fPath) {
		_ = fmt.Errorf("file %s does not exists", fPath)
		return
	}

	// ns := Parse(fName, mName)
	//
	// DbMake(ns, oDirName)

	// fmt.Println(ns.Name)
	// if ns.MainSt != nil {
	// 	fmt.Println("MainSt:")
	// 	for _, f := range ns.MainSt.Fields {
	// 		fmt.Println("  ", *f)
	// 	}
	// }
	// if ns.CuSt != nil {
	// 	fmt.Println("CuSt:")
	// 	for _, f := range ns.CuSt.Fields {
	// 		fmt.Println("  ", *f)
	// 	}
	// }
}
