package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
var y = 0

func aa(){ 
	var int yy 
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = "bl: " + x.Value
		case *ast.Ident:
			s = "id: " + x.Name
		}
		if s != "" {
			log.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})
}

/*
* imagine you have this piece of go code, and you want to find and print out all the constants and variable names, along with their positions in the code. This program does that automatically.
* read the code
* break it down into its fundamental parts (like var names and constants)
* find and print - it looks through these parts, and whenever it finds a number or a name, it prints out what it found and where it found it
 */
