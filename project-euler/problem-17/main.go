package main

import "fmt"

var numbers = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	80: "eighty",
}

func sayNumber(i int) string {
	s := ""
	thousands := i / 1000
	if thousands != 0 {
		s += fmt.Sprintf("%sthousand", numbers[thousands])
	}
	hundreds := i % 1000 / 100
	if hundreds != 0 {
		if thousands != 0 {
			s += "and"
		}
		s += fmt.Sprintf("%shundred", numbers[hundreds])
	}
	teens := i % 100
	if teens != 0 && (hundreds != 0 || thousands != 0) {
		s += "and"
	}
	if teens != 0 {
		if word, found := numbers[teens]; found {
			s += word
		} else {
			tens := teens / 10
			ones := teens % 10
			if tens != 0 {
				if word, found := numbers[tens*10]; found {
					s += word
				} else {
					s += numbers[tens] + "ty"
				}
			}
			if ones != 0 {
				if tens != 0 {
					s += ""
				}
				s += numbers[ones]
			}

		}

	}
	return s
}

func main() {

	letters := 0
	for i := 1; i <= 1000; i++ {
		letters += len(sayNumber(i))
	}

	fmt.Println(letters)

}
