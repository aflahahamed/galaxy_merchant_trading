package internal

import (
	"fmt"
	"strings"
)

type GalaxyKeys struct {
	transaction map[string]int
	metals      map[string]int
}

func (gk *GalaxyKeys) Parsedict(line string) {
	fmt.Println("parsedict")
	if !strings.Contains(line, "Credits") {
		parts := strings.Split(line, " is ")
		key := parts[0]
		value := parts[1]
		fmt.Println("if", key, value)

		// Assuming the value is a Roman numeral, convert it to an integer
		// and add it to the transaction map
		gk.transaction[key] = ConvertRomanToInt(value)
	} else {
		parts := strings.Split(line, " is ")
		key := parts[0]
		value := parts[1]
		// gk.metals[key] = ConvertRomanToInt(value)
		fmt.Println("else", key, value)
	}

}

// ConvertRomanToInt converts a Roman numeral to an integer
// (implementation not provided, you need to implement this)
func ConvertRomanToInt(romanString string) int {
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

	return result + roman[string(romanString[len(romanString)-1])]
}

// ConvertStringToInt converts a string to an integer
// (implementation not provided, you need to implement this)
func ConvertStringToInt(str string) int {
	// Your implementation here
	return 0
}
