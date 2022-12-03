package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"io"
	"os"
	"strings"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(c rune) int {
	return strings.IndexRune(priority, c) + 1
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wholeFile, _ := io.ReadAll(f)
	lines := strings.Split(strings.ReplaceAll(string(wholeFile), "\r\n", "\n"), "\n")

	var result int

	for i := 0; (i + 3) < len(lines); i += 3 {
		line1 := strings.TrimSpace(lines[i])
		line2 := strings.TrimSpace(lines[i+1])
		line3 := strings.TrimSpace(lines[i+2])

		elv1 := mapset.NewSet[rune]()
		elv2 := mapset.NewSet[rune]()
		elv3 := mapset.NewSet[rune]()

		for _, char := range line1 {
			elv1.Add(char)
		}
		for _, char := range line2 {
			elv2.Add(char)
		}
		for _, char := range line3 {
			elv3.Add(char)
		}

		common := elv1.Intersect(elv2).Intersect(elv3)
		it := common.Iterator()
		for elem := range it.C {
			result += getPriority(elem)
		}
	}

	fmt.Println(result)
}
