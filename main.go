package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	expr, err := parser.ParseFile(fset, "target.go", nil, 0)
	if err != nil {
		panic(err)
	}

	for _, d := range expr.Decls {
		f, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if f.Name.IsExported() {
			continue
		}

		fmt.Printf("%s is private function!\n", f.Name.Name)
	}
}
