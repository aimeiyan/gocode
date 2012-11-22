package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {

	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile(fset, "tree/tree.go", nil, parser.DeclarationErrors)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fset)
	fmt.Printf("%#v", f)

	for _, v := range f.Decls {
		switch i := v.(type) {
		case *ast.FuncDecl:
			fmt.Println("func", i.Name, i.Pos(), i.End())
			fmt.Println("------------", *i.Type.Params, i.Body)
		case *ast.GenDecl:
			fmt.Println("gen", i.Tok, i.Pos(), i.End())
		}
		// fmt.Printf("----------------%#v\n", v)
	}

	fmt.Println(f.Name)

	for _, v := range f.Imports {
		fmt.Printf("%#v", v.Path)
	}

	fmt.Println("\n\n=====================================")

	fset.Iterate(func(f *token.File) bool{
		fmt.Println(f.Name(), f.LineCount(), f.Line(345))
		fmt.Println("-----------------")
		return true
	})
}
