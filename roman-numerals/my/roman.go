package main

import (
	"strings"
)

type RomanLiterals struct {
	Value  int
	Symbol string
}
type allRomanNumerals []RomanLiterals

func (pairs allRomanNumerals) ValueOf(roman string) int {
	for _, pair := range pairs {
		if roman == pair.Symbol {
			return pair.Value
		}
	}
	return 0
}

var allLiterals = allRomanNumerals{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}

func ConvertToArabic(roman string) int {
	var result int
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) {
			nextSymbol := roman[i+1]
			potentialNumber := string([]byte{roman[i], nextSymbol})
			value := allLiterals.ValueOf(potentialNumber)
			if value != 0 {
				i++
				result += value
			} else {
				result++
			}
		} else {
			result++
		}
	}
	return result
}
func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allLiterals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}
