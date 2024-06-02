package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("imports:")
	for _, i := range f.Imports {
		log.Println(" ", i.Path.Value)
	}

	log.Println("functions:")
	for _, f := range f.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		log.Printf("  %s\n", fn.Name.Name)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			log.Println(fmt.Sprintf("return statement found in line %d", fset.Position(ret.Pos()).Line))
			return true
		}
		return true
	})
}

/*
* this go code is designed to parse and analyze a go source file named main.go.
* it identifies and logs the imported packages, defined functions, and locations of return statements within the file.
 */
