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
	Name    NameSt
	Fields  []*FieldSt
	IdField *FieldSt
}

type FieldSt struct {
	Name          NameSt
	Type          string
	IsTypePointer bool
	IsTypeSlice   bool
	IsTypeInt     bool
	IsEmbedded    bool
	Tag           string
	JsonName      string
	ZeroValue     string
	PVZeroValue   string
	IsId          bool
	IsNullable    bool
}

func (o *NameSt) Normalize(changeOrigin bool) {
	if o.Origin != "" {
		o.Snake = util.Case2Snake(o.Origin)
		if changeOrigin {
			o.Origin = util.Case2Camel(o.Snake, false)
		}
		o.Camel = o.Origin
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

	for _, st := range []*StructSt{o.MainSt, o.GetParsSt, o.ListSt, o.ListParsSt, o.CuSt} {
		if st == nil {
			continue
		}
		result += "  " + st.Name.Origin + ":\n"
		for _, f := range st.Fields {
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
