package entity

type St struct {
	Name   NameSt
	MainSt *StructSt
	CuSt   *StructSt
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
