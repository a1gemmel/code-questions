package main

func longestCommonPrefix(strs []string) string {
	s1 := strs[0]
	var i int
	for i = 0; i < len(s1); i++ {
		for _, s := range strs[1:] {
			if i >= len(s) || s[i] != s1[i] {
				return s1[0:i]
			}
		}
	}
	return s1
}
