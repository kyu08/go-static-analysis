package shortvar

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const (
	doc               = "shortvar is ..."
	shortestVarLength = 3
)

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "shortvar",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

// run 3文字以下のパッケージ変数を探す
func run(pass *analysis.Pass) (any, error) {
	for _, n := range pass.Pkg.Scope().Names() {
		if len(n) <= shortestVarLength {
			obj := pass.Pkg.Scope().Lookup(n)
			_, ok := obj.(*types.Var)
			if ok {
				pass.Reportf(obj.Pos(), "%s should be more than 3 characters", n)
			}

		}
	}

	return nil, nil
}
