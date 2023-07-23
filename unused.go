package unused

import (
	"go/types"

	"github.com/gostaticanalysis/ident"
	"github.com/samber/lo"
	"golang.org/x/tools/go/analysis"
)

const doc = "unused is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "unused",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		ident.Analyzer,
	},
}

// run 未使用の関数を検出する
func run(pass *analysis.Pass) (any, error) {
	m := pass.ResultOf[ident.Analyzer].(ident.Map)

	for o := range m {
		if isTarget(o) && len(m[o]) == 1 {
			pass.Reportf(o.Pos(), "%q is unused", o.Name())
		}
	}

	return nil, nil
}

func isTarget(o types.Object) bool {
	_, ok := o.(*types.Func)
	excludeFuncNames := []string{"main", "init"}

	return ok && !o.Exported() && !lo.Contains(excludeFuncNames, o.Name())
}
