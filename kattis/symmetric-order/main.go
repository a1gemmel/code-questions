package main

import "fmt"

func main() {
	set := 1
	for {
		var c int
		fmt.Scanf("%d", &c)
		if c == 0 {
			break
		}
		fmt.Printf("SET %d\n", set)
		readnames(c)
		set++
	}
}

func readnames(n int) {
	for i := 0; i < n; i++ {
		var name string
		fmt.Scanf("%s", &name)
		if i%2 == 0 {
			fmt.Println(name)
		} else {
			defer fmt.Println(name)
		}
	}
}
