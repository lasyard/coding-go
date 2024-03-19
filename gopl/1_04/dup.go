package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err == nil {
				countLines(f, counts)
				f.Close()
			} else {
				fmt.Fprintf(os.Stderr, "e_1_4: %v\n", err)
			}
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s\n", len(files), line)
			fmt.Printf("\tin %v\n", files)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		k := input.Text()
		if counts[k] == nil {
			counts[k] = make([]string, 1)
			counts[k][0] = f.Name()
		} else {
			counts[k] = append(counts[k], f.Name())
		}
	}
}
