package db

import (
	"fmt"
	"github.com/rendau/glg/assets"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.DbDirPath == nil {
		fmt.Println("Db destination dir not found")
		return
	}

	tData, err := assets.Asset("templates/db.tmp")
	if err != nil {
		log.Panicln(err)
	}

	t, err := template.New("db.tmp").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
		"inc":     func(x int) int { return x + 1 },
	}).Parse(string(tData))
	if err != nil {
		log.Panicln(err)
	}

	outF, err := os.Create(filepath.Join(pr.DbDirPath.Rel, eName.Snake+".go"))
	if err != nil {
		log.Panicln(err)
	}
	defer outF.Close()

	err = t.Execute(outF, struct {
		Pr    *project.St
		EName *entity.NameSt
		Ent   *entity.St
	}{
		Pr:    pr,
		EName: eName,
		Ent:   ent,
	})
	if err != nil {
		log.Panicln(err)
	}
}
