package internal

import (
	"fmt"
	"go/ast"
	"go/token"
)

//need to find all struct type
//add json tag for all the fields
func FormatCode(fileset *token.FileSet, file *ast.File) {
	ast.Inspect(file, func(node ast.Node) bool {
		st, ok := node.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range st.Fields.List {
			if field.Tag == nil {
				field.Tag = &ast.BasicLit{
					ValuePos: 0,
					Kind:     0,
					Value:    "",
				}
			}
			if len(field.Tag.Value) == 0 {
				field.Tag.Value = fmt.Sprintf("`json:\"%s\"`", toSnakeCase(field.Names[0].String()))
			} else {
				field.Tag.Value = fmt.Sprintf("`%s;json:\"%s\"`", field.Tag.Value[1:len(field.Tag.Value) - 1],toSnakeCase(field.Names[0].String()))
			}
		}

		return false
	})
}
