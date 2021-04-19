package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet
	// can add options eg. fn = flags.String("file", "xxx.go", "name of generated go file")
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(mygen)
}
