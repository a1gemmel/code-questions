package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	letterToIndex := map[byte]int{s[0]: 0}
	longest := 1
	start := 0
	end := 0

	for start <= len(s)-1 && end <= len(s)-2 {
		end++
		if _, found := letterToIndex[s[end]]; !found { // if the next letter is not already in the prefix

			if end-start+1 > longest {
				longest = end - start + 1
			}
		} else { // next letter was already in the prefix
			oldLetterPos := letterToIndex[s[end]]
			for i := start; i <= oldLetterPos; i++ {
				delete(letterToIndex, s[i])
			}
			start = oldLetterPos + 1
		}
		letterToIndex[s[end]] = end
	}
	return longest
}
