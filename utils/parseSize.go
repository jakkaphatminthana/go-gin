package utils

import (
	"strconv"
	"strings"
)

func ParseSize(size string) int64 {
	size = strings.ToUpper(size)
	multiplier := int64(1)

	switch {
	case strings.HasSuffix(size, "K"):
		multiplier = 1024
		size = strings.TrimSuffix(size, "K")
	case strings.HasSuffix(size, "M"):
		multiplier = 1024 * 1024
		size = strings.TrimSuffix(size, "M")
	case strings.HasSuffix(size, "G"):
		multiplier = 1024 * 1024 * 1024
		size = strings.TrimSuffix(size, "G")
	}

	val, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		panic("Invalid body limit format: " + err.Error())
	}

	return val * multiplier
}
