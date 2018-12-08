package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "_gopher.go", nil, 0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defsOrUses := map[*ast.Ident]types.Object{}
	info := &types.Info{
		Defs: defsOrUses,
		Uses: defsOrUses,
	}

	config := &types.Config{
		Importer: importer.Default(),
	}

	_, err = config.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		if ident.Name != "Gopher" {
			return true
		}

		obj := defsOrUses[ident]
		if obj == nil {
			return true
		}

		typ := obj.Type()
		if _, ok := typ.(*types.Named); !ok {
			return true
		}

		fmt.Println(fset.Position(ident.Pos()))

		return true
	})
}
