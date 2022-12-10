package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cycles := []int{1}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		x := cycles[len(cycles)-1]
		cycles = append(cycles, x)
		if a := strings.TrimSpace(line[4:]); len(a) > 0 {
			i, err := strconv.Atoi(a)
			if err != nil {
				panic(err)
			}
			cycles = append(cycles, x+i)
		}
	}

	fmt.Println("Part 1 \t", puzzle1(cycles))
	fmt.Println("Part 2:")
	puzzle2(cycles)
}

func puzzle1(cycles []int) int {
	sum := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		sum += i * cycles[i-1]
	}

	return sum
}

func puzzle2(cycles []int) {
	for i, x := range cycles {
		if abs(i%40-x) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
