package lib

func ReverseStr(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return string(b)
}