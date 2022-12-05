package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() (r rune) {
	r, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return
}

func (s *stack) unshift(r rune) {
	*s = append([]rune{r}, *s...)
}

func main() {
	f, err := os.Open("../input_small.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var stacks []stack

	for sc.Scan() {
		line := sc.Text()

		if len(line) < 1 {
			continue
		}

		switch true {
		case strings.HasPrefix(strings.TrimSpace(line), "["):
			if len(stacks) <= 0 {
				// init
				stacks = make([]stack, (len(line)+1)/4)
			}

			for i, r := range line {
				if r != ' ' && r != '[' && r != ']' {
					stacks[i/4].unshift(r)
				}
			}
		case strings.HasPrefix(line, "move"):
			var move, from, to int
			_, _ = fmt.Sscanf(sc.Text(), "move %d from %d to %d", &move, &from, &to)

			for i := 0; i < move; i++ {
				stacks[to-1].push(stacks[from-1].pop())
			}
		default:
			// ignore 1   2   3... line
		}
	}

	for _, s := range stacks {
		fmt.Print(string(s.pop()))
	}
}
