package base

import (
	"strings"
)

func Capitalize(str string) string {
	upper := strings.ToUpper(string(str[0]))
	rest := str[1:]	
	builder := strings.Builder{}
	builder.WriteString(upper)
	builder.WriteString(rest)

	return builder.String()
}
