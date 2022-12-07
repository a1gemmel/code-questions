package main

import "strings"

var symbolToInt map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	i := 0
	for {
		if len(s) == 0 {
			break
		}
		if strings.HasPrefix(s, "IV") {
			i += 4
			s = s[2:]
			continue
		}
		if strings.HasPrefix(s, "IX") {
			i += 9
			s = s[2:]
			continue
		}
		if strings.HasPrefix(s, "XL") {
			i += 40
			s = s[2:]
			continue
		}
		if strings.HasPrefix(s, "XC") {
			i += 90
			s = s[2:]
			continue
		}
		if strings.HasPrefix(s, "CD") {
			i += 400
			s = s[2:]
			continue
		}
		if strings.HasPrefix(s, "CM") {
			i += 900
			s = s[2:]
			continue
		}
		i += symbolToInt[s[0]]
		s = s[1:]
	}
	return i
}
