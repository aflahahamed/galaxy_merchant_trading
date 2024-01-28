package utils

import (
	"fmt"
)

// ConvertRomanToInt converts a Roman numeral to an integer
func ConvertRomanToInt(romanString string) (int, error) {
	if hasMoreThanThreeConsecutiveRepeats(romanString) {
		return 0, fmt.Errorf("requested number is in invalid format")
	}
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for i := 0; i < len(romanString)-1; i++ {

		if roman[string(romanString[i])] >= roman[string(romanString[i+1])] {
			result += roman[string(romanString[i])]
		}
		if roman[string(romanString[i])] < roman[string(romanString[i+1])] {
			result -= roman[string(romanString[i])]
		}
	}

	return result + roman[string(romanString[len(romanString)-1])], nil
}

// Validating the roman numbers repeating values
func hasMoreThanThreeConsecutiveRepeats(text string) bool {
	count := 1
	for i := 1; i < len(text); i++ {
		if text[i] == text[i-1] {
			count++
			if count > 3 {
				return true
			}
		} else {
			count = 1
		}
	}
	return false
}
