package module

type NsSt struct {
	Name   string
	MainSt *StructSt
	CuSt   *StructSt
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
