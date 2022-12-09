package main

var openForClose map[rune]rune = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	var stack []rune

	for _, r := range s {
		switch r {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			stack = append(stack, r)
		case ')':
			fallthrough
		case '}':
			fallthrough
		case ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != openForClose[r] {
				return false
			}
			stack = stack[0 : len(stack)-2]
		}
	}
	return len(stack) == 0
}
