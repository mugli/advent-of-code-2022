package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var cur int
	top := []int{0, 0, 0}

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		switch line {
		case "":
			if cur > top[0] {
				top[0] = cur
				sort.Ints(top)
			}

			cur = 0
		default:
			num, _ := strconv.Atoi(line)
			cur += num
		}
	}

	if cur > top[0] {
		top[0] = cur
		sort.Ints(top)
	}

	total := 0
	for i, n := range top {
		fmt.Println(i, "\t", n)
		total += n
	}

	fmt.Println("Total", "\t", total)
}
