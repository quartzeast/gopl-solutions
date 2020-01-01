package main

import (
	"bufio"
	"fmt"
	"os"
)

// countLines 统计文件中每行出现的次数，并将其对应的文件名添加到 occurrences。
func countLines(f *os.File, counts map[string]int, occurrences map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		addName(input.Text(), f.Name(), occurrences)
	}
}

// addName 先判断该行对应的文件名是否出现在 occurrences 中，若未出现，则添加。
func addName(line string, filename string, occurrences map[string][]string) {
	for _, name := range occurrences[line] {
		if name == filename {
			return
		}
	}
	occurrences[line] = append(occurrences[line], filename)
}

func main() {
	counts := make(map[string]int)
	occurrences := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occurrences)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occurrences)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, occurrences[line])
		}
	}
}
