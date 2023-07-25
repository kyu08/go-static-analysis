package main

import (
	"github.com/kyu08/shortvar"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(shortvar.Analyzer) }
