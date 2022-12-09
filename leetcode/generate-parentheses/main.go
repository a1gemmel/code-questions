package main

func generateParenthesisInner(s string, opened int, unopened int) []string {
	if unopened == 0 && opened == 0 {
		return []string{s}
	}
	possibilities := []string{}
	if opened > 0 {
		possibilities = append(possibilities, generateParenthesisInner(s+")", opened-1, unopened)...)
	}
	if unopened > 0 {
		possibilities = append(possibilities, generateParenthesisInner(s+"(", opened+1, unopened-1)...)
	}
	return possibilities
}

func generateParenthesis(n int) []string {
	return generateParenthesisInner("", 0, n)
}
