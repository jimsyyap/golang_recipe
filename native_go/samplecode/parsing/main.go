package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
    package main
    func main() {
        println("Hello, World!")
    }`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)
}

/*
* this code takes a simple go program and shos you the building blocks it's made of.
* it's like a toy x-ray machine for looking at other toys to see how they work inside.
* by understanding how other go programs are written, you can learn to write your own go programs better.
* this is also a starting point for building tools that can automatically fix mistakes in code, change how code looks, or even create new programs based on patterns in existing code
*
 */
