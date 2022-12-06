package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, _ := io.ReadAll(f)

	fmt.Println("Puzzle 1 \t", findUniqueWindow(data, 4))
	fmt.Println("Puzzle 2 \t", findUniqueWindow(data, 14))
}

func findUniqueWindow(data []byte, windowSize int) int {
	start := 0
	seen := make(map[byte]int)

	for end := 0; end < len(data); end++ {
		// see more from end pointer
		seen[data[end]]++

		if (end - start) >= windowSize {
			// "unsee" from the start pointer
			seen[data[start]]--
			if seen[data[start]] <= 0 {
				delete(seen, data[start])
			}

			start++
		}

		if len(seen) == windowSize {
			return end + 1
		}
	}

	return -1
}
