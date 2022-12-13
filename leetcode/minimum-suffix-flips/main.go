package main

func minFlips(target string) int {
	flips := 0
	for i := 0; i < len(target); i++ {
		if target[i]-'0' != byte(flips%2) {
			flips += 1
		}
	}
	return flips
}
