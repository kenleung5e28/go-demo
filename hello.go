package main

import (
	"fmt"
)

func main() {
	var s [2]string
	s[0] = "fff"
	s[1] = "kkk"
	for i := 1; i < 100; i++ {
		fmt.Println(s[i%2])
	}
}
