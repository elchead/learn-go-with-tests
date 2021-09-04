package main

import (
	"strings"
)

type RomanLiterals struct {
	Value  int
	Symbol string
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	allLiterals := []RomanLiterals{{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}
	for _, numeral := range allLiterals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}
