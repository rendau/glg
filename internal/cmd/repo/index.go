package repo

import (
	"bytes"
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
	// var err error

	if pr.RepoDirPath == nil {
		fmt.Println("Interfaces dir not found")
		return
	}

	fPath := filepath.Join(pr.RepoDirPath.Abs, "interfaces.go")

	if !util.IsFileExists(fPath) {
		fmt.Println("Repo-interface file not found")
		return
	}

	removeCurrentMethods(fPath, eName)

	side1, side2, ok := util.DivideInterfaceEndPosSides(fPath, "Repo")
	if !ok || side1 == "" || side2 == "" {
		fmt.Println("Fail to register module in db-interfaces. Not found 'Repo' interface type in `" + fPath + "` file")
	}

	t, err := template.New("interfaces.tmp").Parse(tmp)
	if err != nil {
		log.Panicln(err)
	}

	tResultBuffer := &bytes.Buffer{}

	err = t.Execute(tResultBuffer, struct {
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

	err = ioutil.WriteFile(fPath, []byte(side1+"\n\n"+tResultBuffer.String()+side2), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}

func removeCurrentMethods(fPath string, eName *entity.NameSt) {
	fDataRaw, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Panicln(err)
	}

	fData := string(fDataRaw)

	re := regexp.MustCompile(`(?si)(?://\s*` + eName.Snake + `\n\s*)?` + eName.Camel + `(?:Get|List|IdExists|Create|Update|Delete)\([^\n]+\n`)
	fData = re.ReplaceAllString(fData, "")

	re = regexp.MustCompile(`(?si)(?://\s*` + eName.Snake + `\n\s*)?` + eName.Camel + `\w+Exists\([^\n]+\n`)
	fData = re.ReplaceAllString(fData, "")

	err = ioutil.WriteFile(fPath, []byte(fData), os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	util.FmtFile(fPath)
}
