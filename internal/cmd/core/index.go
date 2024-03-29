package core

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

//go:embed tmpl.tmpl
var tmp string

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	var err error

	if pr.CoreDirPath == nil {
		fmt.Println("Core destination dir not found")
		return
	}

	t, err := template.New("core.tmp").Parse(tmp)
	if err != nil {
		log.Panicln(err)
	}

	fPath := filepath.Join(pr.CoreDirPath.Abs, eName.Snake+".go")

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

	registerModule(pr.CoreDirPath.Abs, eName)
}

func registerModule(coreDirPath string, eName *entity.NameSt) {
	const fName = "index.go"

	fPath := filepath.Join(coreDirPath, fName)

	// struct

	side1, side2, ok := util.DivideStructEndPosSides(fPath, "St")
	if !ok || side1 == "" || side2 == "" {
		fmt.Println("Fail to register module in core. Not found 'St' struct type in `" + fName + "` file")
	}

	if regexp.MustCompile(eName.Camel+` +\*`+eName.Camel).FindString(side1) == "" {
		err := ioutil.WriteFile(fPath, []byte(side1+"\n"+eName.Camel+" *"+eName.Camel+side2), os.ModePerm)
		if err != nil {
			log.Panicln(err)
		}

		util.FmtFile(fPath)
	}

	// fun

	side1, side2, ok = util.DivideFuncReturnPosSides(fPath, "New")
	if !ok || side1 == "" || side2 == "" {
		fmt.Println("Fail to register module in core. Not found 'New' function in `" + fName + "` file")
	}

	if regexp.MustCompile(eName.Camel+` += +New`+eName.Camel).FindString(side1) == "" {
		err := ioutil.WriteFile(fPath, []byte(side1+"\nc."+eName.Camel+" = New"+eName.Camel+"(c)"+side2), os.ModePerm)
		if err != nil {
			log.Panicln(err)
		}
	}

	util.FmtFile(fPath)
}
