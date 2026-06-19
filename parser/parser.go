package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strings"
	"codegen/base"
)

type Field struct {
	Name string
	Type string
}

type StructDecl struct {
	Name   string
	Fields []Field
}

func CreateModule(modStr string) StructDecl {
	cmd := strings.Split(modStr, " ")
	modName := cmd[0]
	modFields := cmd[1:]
	fieldsSize := len(modFields)

	fields := make([]Field, fieldsSize)
	for i, field := range modFields {
		fieldDef := strings.Split(field, ":")
		fields[i] = Field{
			Name: base.Capitalize(fieldDef[0]),
			Type: fieldDef[1],
		}
	}

	return StructDecl{
		Name: modName,
		Fields: fields,
	}
}

func NewStruct(newStruct StructDecl) *ast.GenDecl {
	structName := ast.NewIdent(newStruct.Name)

	fieldsSize := len(newStruct.Fields)
	fields := make([]*ast.Field, fieldsSize)
	for i, field := range newStruct.Fields {
		fieldName := ast.NewIdent(field.Name)
		fieldType := ast.NewIdent(field.Type)

		fields[i] = &ast.Field{
			Names: []*ast.Ident{fieldName},
			Type:  fieldType,
		}

	}

	fieldList := &ast.FieldList{
		List: fields,
	}

	baseStructType := &ast.StructType{
		Fields: fieldList,
	}

	baseStructSpec := &ast.TypeSpec{
		Name: structName,
		Type: baseStructType,
	}

	return &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{baseStructSpec},
	}
}

func GenCode(modStr string) string {
	mod := CreateModule(modStr)
	lowerName := mod.Name

	mod.Name = base.Capitalize(mod.Name)
	baseDecl := NewStruct(mod)

	upsertDecl := CreateUpserStruct(mod)

	fileAST := &ast.File{
		Name:  &ast.Ident{Name: lowerName},
		Decls: []ast.Decl{baseDecl},
	}

	var buf bytes.Buffer
	fset := token.NewFileSet()
	formatErr := format.Node(&buf, fset, fileAST)
	if formatErr != nil {
		fmt.Println("Failed to generate module", formatErr)
		os.Exit(1)
	}

	bufStr := buf.String()
	fmt.Println(bufStr)
	return bufStr
}
