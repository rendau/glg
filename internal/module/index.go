package module

func Run(fName, mName, oDirName string) {
	ns := Parse(fName, mName)

	DbMake(ns, oDirName)

	// fmt.Println(ns.Name)
	// if ns.MainSt != nil {
	// 	fmt.Println("MainSt:")
	// 	for _, f := range ns.MainSt.Fields {
	// 		fmt.Println("  ", *f)
	// 	}
	// }
	// if ns.CuSt != nil {
	// 	fmt.Println("CuSt:")
	// 	for _, f := range ns.CuSt.Fields {
	// 		fmt.Println("  ", *f)
	// 	}
	// }
}
