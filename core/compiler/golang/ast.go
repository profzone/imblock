package golang

import (
	"go/token"
	"go/parser"
	"io/ioutil"
	"go/ast"
	"fmt"
)

func Parse(fileName string) error {
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	t := token.NewFileSet()
	astFile, err := parser.ParseFile(t, fileName, source, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(astFile, func(node ast.Node) bool {
		s, ok := node.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range s.Fields.List {
			fmt.Printf("Field: %s\n", field.Names[0].Name)
			fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}
		return false
	})

	return nil
}
