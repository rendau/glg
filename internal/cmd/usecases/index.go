package usecases

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

//go:embed tmpl.tmpl
var tmp string

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.UsecasesDirPath == nil {
		fmt.Println("Usecases destination dir not found")
		return
	}

	t, err := template.New("usecases.tmp").Parse(tmp)
	if err != nil {
		log.Panicln(err)
	}

	fPath := filepath.Join(pr.UsecasesDirPath.Abs, eName.Snake+".go")

	outF, err := os.Create(fPath)
	if err != nil {
		log.Panicln(err)
	}
	defer outF.Close()

	err = t.Execute(outF, struct {
		Pr       *project.St
		EName    *entity.NameSt
		Ent      *entity.St
		Ctx4List map[string]any
	}{
		Pr:    pr,
		EName: eName,
		Ent:   ent,
	})
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}
