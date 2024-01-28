package internal

import (
	"testing"
)

type TestInput[T comparable] struct {
	Name           string
	Input          string
	ExpectedResult T
	ExpectedError  string
}

func initializeTest() GalaxyKeys {
	parser := GalaxyKeys{
		transactionToNumber: map[string]int{
			"glob": 1,
			"pish": 10,
			"prok": 5,
			"tegj": 50,
		},
		metalValue: map[string]float64{
			"Gold":   14450,
			"Iron":   195.5,
			"Silver": 17,
		},
		materialToRoman: map[string]string{
			"glob": "I",
			"pish": "X",
			"prok": "V",
			"tegj": "L",
		},
	}

	return parser
}

const (
	validRomanAlias = "Parse Valid Roman Alias"
	invalidFormat   = "requested number is in invalid format"
	noIdea          = "i have no idea what you are talking about"
)

func TestParseString(t *testing.T) {
	ParseTest := []TestInput[string]{
		{Name: validRomanAlias, Input: "glob is I", ExpectedResult: "success", ExpectedError: ""},
		{Name: "Parse InValid Roman Alias", Input: "glob is I am", ExpectedResult: "success", ExpectedError: invalidFormat},
		{Name: validRomanAlias, Input: "prok is V", ExpectedResult: "success", ExpectedError: ""},
		{Name: validRomanAlias, Input: "pish is X", ExpectedResult: "success", ExpectedError: ""},
		{Name: validRomanAlias, Input: "tegj is L", ExpectedResult: "success", ExpectedError: ""},
	}

	parser := initializeTest()

	for _, test := range ParseTest {
		result, err := parser.Parsedict(test.Input)

		t.Logf("Test Name: %s\n", test.Name)

		if err != nil {
			if test.ExpectedError == "" {
				t.Errorf("Not expected any error, but got %v", err)
			}
		} else if result != test.ExpectedResult {
			t.Errorf("Expected %v, got %v", test.ExpectedResult, result)
		}
	}
}

func TestPrintString(t *testing.T) {

	printTests := []TestInput[string]{
		{Name: "Calculating Credits before assigning", Input: "how many Credits is glob prok Silver ?", ExpectedResult: "glob prok Silver is 68 credits", ExpectedError: invalidFormat},
		{Name: "Query valid metal credits", Input: "how many Credits is glob glob Gold ?", ExpectedResult: "glob glob Gold is 28900 credits", ExpectedError: ""},
		{Name: "Query valid roman alias", Input: "how much is pish tegj glob glob ?", ExpectedResult: "pish tegj glob glob is 42", ExpectedError: ""},
		{Name: "Query comparison of two metal's credits", Input: "Does glob glob Gold has less Credits than pish tegj glob glob Iron?", ExpectedResult: "glob glob Gold has more Credits than pish tegj glob glob Iron", ExpectedError: ""},
		{Name: "Invalid Input", Input: "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", ExpectedResult: "", ExpectedError: "i have no idea what you are talking about"},
	}

	parser := initializeTest()

	for _, test := range printTests {
		result, err := parser.AnswerPrinter(test.Input)
		t.Logf("Test Name: %s\n", test.Name)

		if err != nil {
			if test.ExpectedError == "" {
				t.Errorf("Not expected any error, but got %v", err)
			}
		} else if result != test.ExpectedResult {
			t.Errorf("Expected %v, got %v", test.ExpectedResult, result)
		}
	}
}
