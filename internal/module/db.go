package module

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func DbMake(ns *NsSt, oDirPath string) {
	var err error

	t := template.New("db.tmp")

	t = t.Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
		"inc":     func(x int) int { return x + 1 },
	})

	t, err = t.ParseFiles("./templates/db.tmp")
	if err != nil {
		log.Panicln(err)
	}

	outF, err := os.Create(filepath.Join(oDirPath, "db.go"))
	if err != nil {
		log.Panicln(err)
	}

	err = t.Execute(outF, struct {
		Ns *NsSt
	}{
		Ns: ns,
	})
	if err != nil {
		log.Panicln(err)
	}
}
