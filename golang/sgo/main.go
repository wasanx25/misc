package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"log"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "_gopher.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	//defsOrUses := map[*ast.Ident]types.Object{}
	//info := &types.Info{
	//	Defs: defsOrUses,
	//	Uses: defsOrUses,
	//}
	//
	//config := &types.Config{
	//	Importer: importer.Default(),
	//}
	//
	//_, err = config.Check("main", fset, []*ast.File{f}, info)
	//if err != nil {
	//	log.Fatal("Error: ", err)
	//}

	ast.Inspect(f, func(n ast.Node) bool {
		comment, ok := n.(*ast.Comment)
		if !ok {
			return true
		}
		fmt.Println(comment)

		//obj := defsOrUses[comment]
		//if obj == nil {
		//	return true
		//}

		fmt.Println(fset.Position(comment.Pos()))

		return true
	})

	pass := analysis.Pass{
		Analyzer: atomic.Analyzer,
		Fset: fset,
		Pkg: &types.Package{},
	}

	fmt.Println(pass.String())
}
