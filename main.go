package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aflahahamed/galaxy_merchant_trading/internal"
)

func main() {

	arguments := os.Args[1:]

	if len(arguments) != 1 {
		log.Fatal("Expected one argument of the input text file")
		os.Exit(1)
	}

	file, err := os.Open(arguments[0])
	scanner := bufio.NewScanner(file)
	if err != nil {
		log.Fatal(err)
	}
	keys := internal.DefineGalaxyKeys()

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
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
