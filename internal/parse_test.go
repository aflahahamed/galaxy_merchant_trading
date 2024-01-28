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
