package module

import (
	"log"

	"github.com/rendau/glg/internal/cmd/core"
	"github.com/rendau/glg/internal/cmd/repo"
	"github.com/rendau/glg/internal/cmd/repo_pg"
	"github.com/rendau/glg/internal/cmd/rest"
	"github.com/rendau/glg/internal/cmd/usecases"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
)

func Run(dir, name string) {
	pr := project.Discover(dir)

	eName := &entity.NameSt{Origin: name}
	eName.Normalize(true)

	if pr.EntitiesDirPath == nil {
		log.Fatalln("entity file not found")
		return
	}

	ent := entity.Parse(pr.EntitiesDirPath.Abs, eName)

	// fmt.Println(ent)

	repo_pg.Make(pr, eName, ent)
	repo.Make(pr, eName, ent)
	core.Make(pr, eName, ent)
	usecases.Make(pr, eName, ent)
	rest.Make(pr, eName, ent)
}
