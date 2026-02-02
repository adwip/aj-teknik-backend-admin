package sql_lib

import (
	"fmt"
	"strings"
)

func visualizeQuery(sql string, args []any) {
	parts := strings.Split(sql, "?")
	var builder strings.Builder

	for i, arg := range args {
		if i < len(parts) {
			builder.WriteString(parts[i])
			var v string
			switch a := arg.(type) {
			case string:
				v = "'" + a + "'"
			default:
				v = fmt.Sprint(a)
			}
			builder.WriteString(v)
		}
	}

	if len(parts) > len(args) {
		builder.WriteString(parts[len(parts)-1])
	}

	fmt.Println(yellow + builder.String() + reset)
}
