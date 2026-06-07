package main

import (
	"codegen/parser"
	"flag"
	"fmt"
)

func main() {
	newProjectPtr := flag.String("new-project", "example", "scaffold a new codegen backend")
	modPtr := flag.String("gen", "post title:string content:string likes:uint dislikes:uint*",
		`
	generate a new API module by defining the module name it's fields and datatypes

	example:
	post title:string content:string likes:uint dislikes:uint* user:user
	`)

	flag.Parse()

	fmt.Println(*newProjectPtr)
	parser.GenCode(*modPtr)
}
