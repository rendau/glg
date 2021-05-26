package interfaces

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/rendau/glg/assets"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
	"github.com/rendau/glg/internal/util"
)

func Make(pr *project.St, eName *entity.NameSt, ent *entity.St) {
	// var err error

	if pr.InterfacesDirPath == nil {
		fmt.Println("Interfaces dir not found")
		return
	}

	fPath := filepath.Join(pr.InterfacesDirPath.Abs, "db.go")

	if !util.IsFileExists(fPath) {
		fmt.Println("Db-interface file not found")
		return
	}

	removeCurrentMethods(fPath, eName)

	side1, side2, ok := util.DivideInterfaceEndPosSides(fPath, "Db")
	if !ok || side1 == "" || side2 == "" {
		fmt.Println("Fail to register module in db-interfaces. Not found 'Db' interface type in `" + fPath + "` file")
	}

	tData, err := assets.Asset("templates/interfaces.tmpl")
	if err != nil {
		log.Panicln(err)
	}

	t, err := template.New("interfaces.tmp").Parse(string(tData))
	if err != nil {
		log.Panicln(err)
	}

	tResultBuffer := &bytes.Buffer{}

	err = t.Execute(tResultBuffer, struct {
		Pr       *project.St
		EName    *entity.NameSt
		Ent      *entity.St
		Ctx4List map[string]interface{}
	}{
		Pr:       pr,
		EName:    eName,
		Ent:      ent,
		Ctx4List: getCtx4List(pr, eName, ent),
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

func getCtx4List(pr *project.St, eName *entity.NameSt, ent *entity.St) map[string]interface{} {
	result := map[string]interface{}{}

	if ent.ListParsSt != nil {
		for _, field := range ent.ListParsSt.Fields {
			if strings.Contains(strings.ToLower(field.Type), "pagination") {
				result["hasPagination"] = true
				break
			}
		}
	}

	return result
}
