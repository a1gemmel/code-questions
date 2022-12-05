package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	matchText   = "EVEN"
	noMatchText = "NOT EVEN"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	lineNum := 0

	for {
		lineNum++
		scanner.Scan()
		line := scanner.Text()
		if line == "END" {
			break
		}
		line = strings.TrimPrefix(strings.TrimSuffix(line, "*"), "*")
		whitespaces := strings.Split(line, "*")

		if len(whitespaces) == 0 {
			fmt.Println(lineNum, matchText)
			continue
		}
		gap := len(whitespaces[0])
		even := true
		for _, w := range whitespaces {
			if len(w) != gap {
				fmt.Println(lineNum, noMatchText)
				even = false
				break
			}
		}
		if even {
			fmt.Println(lineNum, matchText)
		}

	}
}
