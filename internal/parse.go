package internal

import (
	"regexp"
	"strconv"
	"strings"

	utils "github.com/aflahahamed/galaxy_merchant_trading/utils"
)

type GalaxyKeys struct {
	transactionToNumber map[string]int
	metalValue          map[string]float64
	materialToRoman     map[string]string
}

func DefineGalaxyKeys() *GalaxyKeys {
	return &GalaxyKeys{
		transactionToNumber: make(map[string]int),
		materialToRoman:     make(map[string]string),
		metalValue:          make(map[string]float64),
	}
}

// Parses the lines and fills up the GalaxyKeys struct
func (gk *GalaxyKeys) Parsedict(line string) (result string, err error) {
	if !strings.Contains(line, "Credits") {
		parts := strings.Split(line, " is ")
		key := parts[0]
		value := parts[1]
		gk.materialToRoman[key] = value
		transactionNumber, err := utils.ConvertRomanToInt(value)
		if err != nil {
			return "", err
		}

		gk.transactionToNumber[key] = transactionNumber

		return "success", nil

	} else {
		parts := strings.Split(line, " is ")
		key := parts[0]
		value := parts[1]
		keys := strings.Split(key, " ")
		romanStringList := keys[:len(keys)-1]
		for i := 0; i < len(romanStringList); i++ {
			romanStringList[i] = gk.materialToRoman[romanStringList[i]]
		}
		metalName := keys[len(keys)-1]

		romanString := strings.Join(romanStringList, "")

		re := regexp.MustCompile(`\d+`)
		values, err := strconv.Atoi(re.FindAllString(value, -1)[0])
		if err != nil {
			return "", err
		}

		denominator, err := utils.ConvertRomanToInt(romanString)
		if err != nil {
			return "", err
		}

		gk.metalValue[metalName] = float64(values) / float64(denominator)
		return "success", nil

	}

}
