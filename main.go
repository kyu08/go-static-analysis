package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
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

		for _, l := range f.Body.List {
			s, ok := l.(*ast.ReturnStmt)
			if !ok {
				continue
			}

			r, ok := s.Results[0].(*ast.BasicLit)
			if !ok {
				continue
			}

			emptyStr, err := strconv.Unquote(r.Value)
			if err != nil {
				continue
			}

			if r.Kind == token.STRING && len(emptyStr) == 0 {
				fmt.Printf("%s returns empty string!\n", f.Name.Name)
			}
		}
	}
}
