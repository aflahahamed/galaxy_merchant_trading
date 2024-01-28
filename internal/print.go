package internal

import (
	"fmt"
	"math"
	"strings"
)

var (
	result string
)

func (gk *GalaxyKeys) AnswerPrinter(line string) (string, error) {
	if strings.Contains(line, "how much is") {
		cleanedText := strings.ReplaceAll(line, "how much is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, " ?", "")
		cleanedTextSplit := strings.Split(cleanedText, " ")

		for i := 0; i < len(cleanedTextSplit); i++ {
			cleanedTextSplit[i] = gk.materialToRoman[cleanedTextSplit[i]]
		}
		romanString := strings.Join(cleanedTextSplit, "")
		answer, err := ConvertRomanToInt(romanString)

		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		result = fmt.Sprintf("%s is %d", cleanedText, answer)
		return result, nil

	} else if strings.Contains(line, "how many Credits is ") {
		cleanedText := strings.ReplaceAll(line, "how many Credits is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, " ?", "")
		cleanedTextSplit := strings.Split(cleanedText, " ")

		for i := 0; i < len(cleanedTextSplit)-1; i++ {
			cleanedTextSplit[i] = gk.materialToRoman[cleanedTextSplit[i]]
		}

		metalValueString := cleanedTextSplit[len(cleanedTextSplit)-1]
		romanString := strings.Join(cleanedTextSplit[:len(cleanedTextSplit)-1], "")
		answer, err := ConvertRomanToInt(romanString)

		if err != nil {
			return "", err
		}

		resultValue := float64(answer) * gk.metalValue[metalValueString]

		if resultValue == math.Floor(resultValue) {
			result = fmt.Sprintf("%s is %d credits", cleanedText, int(resultValue))

		} else {
			result = fmt.Sprintf("%s is %.1f credits", cleanedText, resultValue)

		}

		return result, nil

	} else if strings.Contains(line, " has more Credits than ") {
		cleanedText := strings.Replace(line, " has more Credits than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Does ", "")
		cleanedText = strings.ReplaceAll(cleanedText, " ?", "")
		comparators := strings.Split(cleanedText, "-")
		a, err := gk.TextToValueWithMetal(comparators[0])

		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}
		b, err := gk.TextToValueWithMetal(comparators[1])

		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		if a < b {
			result = fmt.Sprintf("%s has less credit than %s", comparators[0], comparators[1])

		} else {
			result = fmt.Sprintf("%s has more credit than %s", comparators[0], comparators[1])

		}

		return result, nil

	} else if strings.Contains(line, " has less Credits than ") {
		cleanedText := strings.Replace(line, " has less Credits than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Does ", "")
		cleanedText = strings.ReplaceAll(cleanedText, "?", "")
		comparators := strings.Split(cleanedText, "-")

		a, err := gk.TextToValueWithMetal(comparators[0])
		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		b, err := gk.TextToValueWithMetal(comparators[1])
		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		if a < b {
			result = fmt.Sprintf("%s has less Credits than %s", comparators[0], comparators[1])

		} else {
			result = fmt.Sprintf("%s has more Credits than %s", comparators[0], comparators[1])

		}

		return result, nil

	} else if strings.Contains(line, " larger than ") {
		cleanedText := strings.Replace(line, " larger than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, "?", "")
		comparators := strings.Split(cleanedText, "-")
		leftSide := strings.Split(comparators[0], " ")

		a, err := gk.TextToValue(leftSide)
		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		rightSide := strings.Split(comparators[1], " ")
		b, err := gk.TextToValue(rightSide)
		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		if a < b {
			result = fmt.Sprintf("%s is smaller than %s", comparators[0], comparators[1])

		} else {
			result = fmt.Sprintf("%s is larger than %s", comparators[0], comparators[1])
		}

		return result, nil

	} else if strings.Contains(line, " smaller than ") {
		cleanedText := strings.Replace(line, " smaller than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, "?", "")
		comparators := strings.Split(cleanedText, "-")
		leftSide := strings.Split(comparators[0], " ")

		a, err := gk.TextToValue(leftSide)
		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		rightSide := strings.Split(comparators[1], " ")
		b, err := gk.TextToValue(rightSide)
		if err != nil {
			return "", fmt.Errorf("requested number is in invalid format")
		}

		if a < b {
			result = fmt.Sprintf("%s is smaller than %s", comparators[0], comparators[1])

		} else {
			result = fmt.Sprintf("%s is larger than %s", comparators[0], comparators[1])
		}
		return result, nil

	} else {

		return "", fmt.Errorf("i have no idea what you are talking about")

	}

}

func (gk *GalaxyKeys) TextToValueWithMetal(text string) (float64, error) {

	cleanedTextSplit := strings.Split(text, " ")

	for i := 0; i < len(cleanedTextSplit)-1; i++ {
		cleanedTextSplit[i] = gk.materialToRoman[cleanedTextSplit[i]]
	}

	metalValueString := cleanedTextSplit[len(cleanedTextSplit)-1]
	romanString := strings.Join(cleanedTextSplit[:len(cleanedTextSplit)-1], "")

	answer, err := ConvertRomanToInt(romanString)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	resultValue := float64(answer) * gk.metalValue[metalValueString]

	return resultValue, nil
}

func (gk *GalaxyKeys) TextToValue(text []string) (int, error) {

	for i := 0; i < len(text); i++ {
		text[i] = gk.materialToRoman[text[i]]
	}

	romanString := strings.Join(text, "")
	answer, err := ConvertRomanToInt(romanString)

	if err != nil {
		fmt.Println(err.Error())

		return 0, err
	}

	return answer, nil
}
