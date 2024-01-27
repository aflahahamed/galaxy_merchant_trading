package internal

import (
	"fmt"
	"strings"
)

func (gk *GalaxyKeys) Questionparser(line string) {
	if strings.Contains(line, "how much is") {
		cleanedText := strings.ReplaceAll(line, "how much is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, " ?", "")
		keyss := strings.Split(cleanedText, " ")

		for i := 0; i < len(keyss); i++ {
			keyss[i] = gk.materialToRoman[keyss[i]]
		}
		romanString := strings.Join(keyss, "")
		answer, err := ConvertRomanToInt(romanString)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(cleanedText, "is", answer)
	} else if strings.Contains(line, "how many Credits is ") {
		cleanedText := strings.ReplaceAll(line, "how many Credits is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, " ?", "")

		keyss := strings.Split(cleanedText, " ")
		for i := 0; i < len(keyss)-1; i++ {
			keyss[i] = gk.materialToRoman[keyss[i]]
		}
		metalValueString := keyss[len(keyss)-1]
		romanString := strings.Join(keyss[:len(keyss)-1], "")
		answer, err := ConvertRomanToInt(romanString)
		if err != nil {
			fmt.Println(err)
			return
		}

		resultValue := float64(answer) * gk.metalValue[metalValueString]
		fmt.Printf("%s is %.1f credits\n", cleanedText, resultValue)

	} else if strings.Contains(line, " has more Credits than ") {
		cleanedText := strings.Replace(line, " has more Credits than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Does ", "")
		cleanedText = strings.ReplaceAll(cleanedText, " ?", "")
		comparators := strings.Split(cleanedText, "-")

		a, err := gk.TextToValueWithMetal(comparators[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := gk.TextToValueWithMetal(comparators[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		if a < b {
			fmt.Printf("%s has less credit than %s\n", comparators[0], comparators[1])

		} else {
			fmt.Printf("%s has more credit than %s\n", comparators[0], comparators[1])

		}
		return
	} else if strings.Contains(line, " has less Credits than ") {
		cleanedText := strings.Replace(line, " has less Credits than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Does ", "")
		cleanedText = strings.ReplaceAll(cleanedText, "?", "")
		comparators := strings.Split(cleanedText, "-")

		a, err := gk.TextToValueWithMetal(comparators[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := gk.TextToValueWithMetal(comparators[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		if a < b {
			fmt.Printf("%s has less Credits than %s\n", comparators[0], comparators[1])

		} else {
			fmt.Printf("%s has more Credits than %s\n", comparators[0], comparators[1])

		}
		return
	} else if strings.Contains(line, " larger than ") {
		cleanedText := strings.Replace(line, " larger than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, "?", "")
		comparators := strings.Split(cleanedText, "-")

		leftSide := strings.Split(comparators[0], " ")
		a, err := gk.TextToValue(leftSide)
		if err != nil {
			fmt.Println(err)
			return
		}

		rightSide := strings.Split(comparators[1], " ")
		b, err := gk.TextToValue(rightSide)
		if err != nil {
			fmt.Println(err)
			return
		}

		if a < b {
			fmt.Printf("%s is smaller than %s\n", comparators[0], comparators[1])
		} else {
			fmt.Printf("%s is larger than %s\n", comparators[0], comparators[1])
		}
		return
	} else if strings.Contains(line, " smaller than ") {
		cleanedText := strings.Replace(line, " smaller than ", "-", 1)
		cleanedText = strings.ReplaceAll(cleanedText, "Is ", "")
		cleanedText = strings.ReplaceAll(cleanedText, "?", "")
		comparators := strings.Split(cleanedText, "-")

		leftSide := strings.Split(comparators[0], " ")
		a, err := gk.TextToValue(leftSide)
		if err != nil {
			fmt.Println(err)
			return
		}

		rightSide := strings.Split(comparators[1], " ")
		b, err := gk.TextToValue(rightSide)
		if err != nil {
			fmt.Println(err)
			return
		}

		if a < b {
			fmt.Printf("%s is smaller than %s\n", comparators[0], comparators[1])
		} else {
			fmt.Printf("%s is larger than %s\n", comparators[0], comparators[1])
		}
		return
	} else {
		fmt.Printf("I have no idea what you are talking about\n")
	}

}

func (gk *GalaxyKeys) TextToValueWithMetal(text string) (float64, error) {

	keyss := strings.Split(text, " ")
	for i := 0; i < len(keyss)-1; i++ {
		keyss[i] = gk.materialToRoman[keyss[i]]
	}
	metalValueString := keyss[len(keyss)-1]
	romanString := strings.Join(keyss[:len(keyss)-1], "")
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
