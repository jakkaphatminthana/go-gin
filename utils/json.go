package utils

import (
	"strings"
	"unicode"

	"github.com/go-playground/validator"
)

func ToCamelCase(input string) string {
	isToUpper := false
	output := ""

	for i, r := range input {
		if r == '_' {
			isToUpper = true
			continue
		}
		if isToUpper {
			output += string(unicode.ToUpper(r))
			isToUpper = false
		} else {
			if i == 0 {
				output += string(unicode.ToLower(r))
			} else {
				output += string(r)
			}
		}
	}
	return output
}

func GetJsonFieldNameByErrorField(fe validator.FieldError) string {
	jsonName := fe.Field() // fallback
	if fe.StructField() != "" {
		jsonName = strings.ToLower(string(fe.StructField()[0])) + fe.StructField()[1:]
	}

	return jsonName
}
