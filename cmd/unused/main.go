package main

import (
	"github.com/kyu08/unused"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unused.Analyzer) }
