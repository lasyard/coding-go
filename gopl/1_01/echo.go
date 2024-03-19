package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[0]
	for i := 1; i < len(os.Args); i++ {
		s += " " + os.Args[i]
	}
	fmt.Println(s)
}
