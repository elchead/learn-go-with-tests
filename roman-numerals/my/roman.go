package main

import (
	"strings"
)

var allLiterals = romanNumerals{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}

type romanNumeral struct {
	Value  int
	Symbol string
}
type romanNumerals []romanNumeral

func (pairs romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, pair := range pairs {
		if symbol == pair.Symbol {
			return pair.Value
		}
	}
	return 0
}

func (pairs romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, pair := range pairs {
		if symbol == pair.Symbol {
			return true
		}
	}
	return false
}

func isSubtractive(currentSymbol byte) bool {
	return currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'

}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allLiterals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func ConvertToArabic(roman string) int {
	var result int
	for _, symbols := range windowedRoman(roman).Symbols() {
		result += allLiterals.ValueOf(symbols...)
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
