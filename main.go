package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aflahahamed/galaxy_merchant_trading/internal"
)

func main() {
	keys := internal.DefineGalaxyKeys()

	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if len(line) != 0 {
			if !strings.Contains(line, "?") {
				_, err := keys.Parsedict(line)
				if err != nil {
					fmt.Println("Error during parsing")
				}
			} else {
				result, err := keys.AnswerPrinter(line)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println(result)
			}
		}

	}

}
