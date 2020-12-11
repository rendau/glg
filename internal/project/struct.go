package project

type St struct {
	Uri               string
	EntitiesDirPath   *PathSt
	DbDirPath         *PathSt
	InterfacesDirPath *PathSt
	CoreDirPath       *PathSt
	UsecasesDirPath   *PathSt
	RestDirPath       *PathSt
}

type PathSt struct {
	Abs string
	Rel string
}
