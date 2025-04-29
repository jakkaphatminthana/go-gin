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
	// Get the top-level struct type from the namespace
	t := fe.StructField()

	// fallback to json key using camelCase
	return ToCamelCase(strings.ToLower(t))
}
