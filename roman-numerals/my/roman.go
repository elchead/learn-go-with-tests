package main

import (
	"strings"
)

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	switch {
	case arabic == 4:
		result.WriteString("IV")
	case arabic < 4:
		for i := 0; i < arabic; i++ {
			result.WriteString("I")
		}
	}
	return result.String()
}
