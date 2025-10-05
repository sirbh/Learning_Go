package main

import (
	"fmt"
	"learn/testProgs"
)

func main() {
	str := "hello world"
	counts := testProgs.CharCounter(str)

	for ch, count := range counts {
		fmt.Printf("%q : %d\n", ch, count)
	}
}
