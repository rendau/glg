package project

type St struct {
	Uri             string
	EntitiesDirPath *PathSt
	RepoDirPath     *PathSt
	RepoPgDirPath   *PathSt
	CoreDirPath     *PathSt
	UsecasesDirPath *PathSt
	RestDirPath     *PathSt
}

type PathSt struct {
	Abs string
	Rel string
}
