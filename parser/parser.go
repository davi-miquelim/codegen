package parser

import (
	"strings"
	"fmt"
)

type ModuleField struct {
	name     string
	dataType string
}

type Module struct {
	name   string
	fields []ModuleField
}

func CreateModule(modStr string) Module {
	cmd := strings.Split(modStr, " ")
	modName := cmd[0]
	modFields := cmd[1:]
	fieldsCount := len(modFields)

	// TODO: support relations
	fields := make([]ModuleField, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		fieldDef := strings.Split(modFields[i], ":")
		fmt.Println(fieldDef)
		fields[i] = ModuleField{
			name:     fieldDef[0],
			dataType: fieldDef[1],
		}
	}

	return Module{
		name:   modName,
		fields: fields,
	}
}
