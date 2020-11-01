package module

import (
	"fmt"
	"github.com/rendau/glg/internal/entity"
	"github.com/rendau/glg/internal/project"
)

func Run(name string) {
	pr := project.Discover()

	ent := entity.Parse(pr.EntitiesDirPath, name)

	if ent.MainSt != nil {
		fmt.Println("MainSt:")
		for _, f := range ent.MainSt.Fields {
			fmt.Println("  ", *f)
		}
	}
	if ent.CuSt != nil {
		fmt.Println("CuSt:")
		for _, f := range ent.CuSt.Fields {
			fmt.Println("  ", *f)
		}
	}
}
