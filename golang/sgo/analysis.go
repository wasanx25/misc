package main

import (
	"fmt"
	"golang.org/x/tools/go/analysis"
)

func main() {
	analyzer := &analysis.Analyzer{
		Name: "used",
		Doc: "aa",
		Run: run,
	}

	pass := &analysis.Pass{
		Analyzer: analyzer,
	}

	fmt.Println(pass.String())
}

func run(pass *analysis.Pass) (interface{}, error) {
	pass.
}
