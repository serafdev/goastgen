package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

func Parse(filename string) (dump string, err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)

	var bf bytes.Buffer
	ast.Fprint(&bf, fset, f, func(string, reflect.Value) bool {
		return true
	})

	return string(bf.Bytes()), nil
}
