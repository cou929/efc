package main

import (
	"github.com/cou929/efc"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(efc.Analyzer) }