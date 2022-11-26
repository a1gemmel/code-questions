package main

import (
	"fmt"

	"github.com/a1gemmel/project-euler/golang/lib"
)

func main() {

	fmt.Println(lib.SumInts(lib.BigIntFactorial(100).Repr()))
}
