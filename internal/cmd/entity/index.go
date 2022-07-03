package entity

import (
	"bytes"
	_ "embed"
	"io/ioutil"
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

func Run(dir, name string) {
	pr := project.Discover(dir)

	eName := &entity.NameSt{Origin: name}
	eName.Normalize(true)

	if pr.EntitiesDirPath == nil {
		log.Fatalln("entity dir not found")
		return
	}

	fPath := filepath.Join(pr.EntitiesDirPath.Abs, eName.Snake+".go")

	t, err := template.New("entity.tmp").Parse(tmp)
	if err != nil {
		log.Panicln(err)
	}

	tResultBuffer := &bytes.Buffer{}

	err = t.Execute(tResultBuffer, struct {
		EName *entity.NameSt
	}{
		EName: eName,
	})
	if err != nil {
		log.Panicln(err)
	}

	err = ioutil.WriteFile(fPath, tResultBuffer.Bytes(), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}
