package main

import (
	"fmt"
	"strings"
)

func ConvertToRoman(arabic int) string {

	// A Builder is used to efficiently build a string using Write methods.
	// It minimizes memory copying.
	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	fmt.Println("result:", result.String())
	return result.String()
}
