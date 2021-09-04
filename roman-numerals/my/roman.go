package main

import "errors"

func ConvertToRoman(arabic int) (string, error) {
	switch {
	case arabic < 5:
		var result string
		for i := 0; i < arabic; i++ {
			result += "I"
		}
		return result, nil
	}
	return "", errors.New("No roman literal found")
}
