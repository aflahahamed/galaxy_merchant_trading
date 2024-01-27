package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type GalaxyKeys struct {
	transactionToNumber map[string]int
	metalValue          map[string]float64
	materialToRoman     map[string]string
}

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}

func DefineGalaxyKeys() *GalaxyKeys {
	return &GalaxyKeys{
		transactionToNumber: make(map[string]int),
		materialToRoman:     make(map[string]string),
		metalValue:          make(map[string]float64),
	}
}

func (gk *GalaxyKeys) Parsedict(line string) {
	if !strings.Contains(line, "Credits") {
		parts := strings.Split(line, " is ")
		key := parts[0]
		value := parts[1]
		gk.materialToRoman[key] = value
		transactionNumber, err := ConvertRomanToInt(value)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		gk.transactionToNumber[key] = transactionNumber

	} else {
		parts := strings.Split(line, " is ")
		key := parts[0]
		value := parts[1]
		keyss := strings.Split(key, " ")
		romanStringList := keyss[:len(keyss)-1]
		for i := 0; i < len(romanStringList); i++ {
			romanStringList[i] = gk.materialToRoman[romanStringList[i]]
		}
		metalName := keyss[len(keyss)-1]

		romanString := strings.Join(romanStringList, "")

		re := regexp.MustCompile(`\d+`)
		values, err := strconv.Atoi(re.FindAllString(value, -1)[0])
		if err != nil {
			fmt.Errorf(err.Error())
		}

		denominator, err := ConvertRomanToInt(romanString)
		if err != nil {
			fmt.Errorf(err.Error())

		}

		gk.metalValue[metalName] = float64(values) / float64(denominator)

	}
}

// ConvertRomanToInt converts a Roman numeral to an integer
// (implementation not provided, you need to implement this)
func ConvertRomanToInt(romanString string) (int, error) {
	if hasMoreThanThreeConsecutiveRepeats(romanString) {
		return 0, fmt.Errorf("Requested number is in invalid format error")
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
	// integer, err := rom.StringToInt(romanString)
	// if err != nil {
	// 	return 0, err
	// }
	// fmt.Println("lalalal", integer)

	// return integer, nil
}

// ConvertStringToInt converts a string to an integer
// (implementation not provided, you need to implement this)
func ConvertStringToInt(str string) int {
	// Your implementation here
	return 0
}

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
