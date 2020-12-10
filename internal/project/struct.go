package project

type St struct {
	Uri             string
	EntitiesDirPath *PathSt
	DbDirPath       *PathSt
	CoreDirPath     *PathSt
	UsecasesDirPath *PathSt
}

type PathSt struct {
	Abs string
	Rel string
}
