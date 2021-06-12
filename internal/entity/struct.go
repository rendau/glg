package entity

import (
	"fmt"

	"github.com/rendau/glg/internal/util"
)

type St struct {
	Name       *NameSt
	MainSt     *StructSt
	GetParsSt  *StructSt
	ListSt     *StructSt
	ListParsSt *StructSt
	CuSt       *StructSt
	IdField    *FieldSt
}

type NameSt struct {
	Origin string
	Camel  string
	LCamel string
	Snake  string
}

type StructSt struct {
	Fields  []*FieldSt
	IdField *FieldSt
}

type FieldSt struct {
	Name          NameSt
	Type          string
	IsTypePointer bool
	IsTypeSlice   bool
	IsTypeInt     bool
	Tag           string
	JsonName      string
	ZeroValue     string
	PVZeroValue   string
	IsId          bool
	IsNullable    bool
}

func (o *NameSt) Normalize() {
	if o.Origin != "" {
		o.Snake = util.Case2Snake(o.Origin)
		o.Origin = util.Case2Camel(o.Snake, false)
		o.Camel = util.Case2Camel(o.Snake, false)
		o.LCamel = util.Case2Camel(o.Snake, true)
	} else {
		o.Camel = ""
		o.Snake = ""
	}
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
	if o.GetParsSt != nil {
		result += "  " + eName + "GetParsSt:\n"
		for _, f := range o.GetParsSt.Fields {
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

func (o *FieldSt) DefineZeroValue() {
	switch o.Type {
	case "bool":
		o.ZeroValue = "false"
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"time.Time":
		o.ZeroValue = "0"
	case "string":
		o.ZeroValue = `""`
	default:
		o.ZeroValue = "nil"
	}
}

func (o *FieldSt) DefinePVZeroValue() {
	switch o.Type {
	case "*int", "*int8", "*int16", "*int32", "*int64",
		"*uint", "*uint8", "*uint16", "*uint32", "*uint64",
		"*float32", "*float64":
		o.PVZeroValue = "0"
	default:
		o.PVZeroValue = ""
	}
}
