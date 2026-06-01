package main

import (
	"fmt"
	"flag"
)

func main() {
	newProjectPtr := flag.String("new-project", "example", "scaffold a new codegen backend")
	genPtr := flag.String("gen", "post title:string content:string likes:uint dislikes:uint*",
	`
	generate a new API module by defining the module name it's fields and datatypes

	example:
	post title:string content:string likes:uint dislikes:uint* user:user
	`)

	flag.Parse()

	fmt.Println("new project: ", *newProjectPtr)
	fmt.Println("generate: ", *genPtr)
}
