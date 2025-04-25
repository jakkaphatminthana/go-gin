package utils

import (
	"fmt"
	"strconv"
)

func ParseStringToUint64(value string) (uint64, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse string to uint64: %w", err)
	}
	return parsedValue, nil
}
