package main

import "errors"

func ConvertToRoman(arabic int) (string, error) {
	switch {
	case arabic == 1:
		return "I", nil
	}
	return "", errors.New("No valid arabic number")
}
