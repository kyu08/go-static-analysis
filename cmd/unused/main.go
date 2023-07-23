package main

import (
	"github.com/kyu08/findint"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findint.Analyzer) }
