package module

import (
	"log"

	"github.com/rendau/glg/internal/cmd/core"
	"github.com/rendau/glg/internal/cmd/db"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
)

func Run(dir, name string) {
	pr := project.Discover(dir)

	eName := &entity.NameSt{Origin: name}
	eName.Normalize()

	if pr.EntitiesDirPath == nil {
		log.Fatalln("entity file not found")
		return
	}

	ent := entity.Parse(pr.EntitiesDirPath.Abs, eName)

	// fmt.Println(ent)

	core.Make(pr, eName, ent)
	db.Make(pr, eName, ent)
}
