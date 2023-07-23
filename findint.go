package findint

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "findint is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "findint",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// run int型の式を見つける
func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(nil, func(n ast.Node) {
		expr, _ := n.(ast.Expr)
		typ := pass.TypesInfo.TypeOf(expr)

		if types.Identical(typ, types.Typ[types.Int]) {
			pass.Reportf(n.Pos(), "int found")
		}
	})

	return nil, nil
}
