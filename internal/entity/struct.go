package entity

import (
	"fmt"
)

type St struct {
	Name       *NameSt
	MainSt     *StructSt
	ListSt     *StructSt
	ListParsSt *StructSt
	CuSt       *StructSt
}

type NameSt struct {
	Camel string
	Snake string
}

type StructSt struct {
	Fields []*FieldSt
}

type FieldSt struct {
	Name     string
	Type     string
	Tag      string
	JsonName string
}

func (o *St) String() string {
	eName := ""

	if o.Name != nil {
		eName = o.Name.Camel
	}

	result := ""

	if eName != "" {
		result = eName + ":\n"
	} else {
		result = "Entity:\n"
	}

	if o.MainSt != nil {
		result += "  " + eName + "St:\n"
		for _, f := range o.MainSt.Fields {
			result += fmt.Sprintln("    ", *f)
		}
	}
	if o.ListSt != nil {
		result += "  " + eName + "ListSt:\n"
		for _, f := range o.ListSt.Fields {
			result += fmt.Sprintln("    ", *f)
		}
	}
	if o.ListParsSt != nil {
		result += "  " + eName + "ListParsSt:\n"
		for _, f := range o.ListParsSt.Fields {
			result += fmt.Sprintln("    ", *f)
		}
	}
	if o.CuSt != nil {
		result += "  " + eName + "CUSt:\n"
		for _, f := range o.CuSt.Fields {
			result += fmt.Sprintln("    ", *f)
		}
	}

	return result
}
