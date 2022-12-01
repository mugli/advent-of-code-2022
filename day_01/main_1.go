package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var top, cur int

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		switch line {
		case "":
			if cur > top {
				top = cur
			}

			cur = 0
		default:
			num, _ := strconv.Atoi(line)
			cur = cur + num
		}
	}

	if cur > top {
		top = cur
	}

	fmt.Println(top)
}
