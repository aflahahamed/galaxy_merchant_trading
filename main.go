package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {

		if !strings.Contains(line, "?") {
			parser.parsedict(line)
		} else {
			// parser.solve(line)
		}
	}

}
