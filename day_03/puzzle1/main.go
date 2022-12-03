package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
	"strings"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(c rune) int {
	return strings.IndexRune(priority, c) + 1
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var result int

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		if len(line) < 1 {
			continue
		}

		comp1 := line[:len(line)/2]
		comp2 := line[len(line)/2:]

		comp1Set := mapset.NewSet[rune]()
		comp2Set := mapset.NewSet[rune]()

		for _, char := range comp1 {
			comp1Set.Add(char)
		}
		for _, char := range comp2 {
			comp2Set.Add(char)
		}

		common := comp1Set.Intersect(comp2Set)
		it := common.Iterator()
		for elem := range it.C {
			result += getPriority(elem)
		}
	}

	fmt.Println(result)
}
