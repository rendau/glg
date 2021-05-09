package main

//go:generate -command templates_debug go-bindata -debug -pkg assets -o assets/index.go templates/
//go:generate templates_debug
//go:generate -command templates go-bindata -pkg assets -o assets/index.go templates/
//go:generate templates

import "github.com/rendau/glg/cmd"

func main() {
	cmd.Execute()
}
