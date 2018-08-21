//+build ignore

package main

import (
	"log"

	"github.com/Gonzih/screensaver-off-go/icons"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(icons.Assets, vfsgen.Options{
		Filename:     "icons/icons_vfsdata.go",
		PackageName:  "icons",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
