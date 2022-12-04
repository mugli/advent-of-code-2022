package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
	"strconv"
	"strings"
)

func makeSet(rangeStr string) mapset.Set[int] {
	set := mapset.NewSet[int]()

	split := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	for i := start; i <= end; i++ {
		set.Add(i)
	}

	return set
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)

	var result int

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		if len(line) < 1 {
			continue
		}

		pair := strings.Split(line, ",")

		elf1 := makeSet(pair[0])
		elf2 := makeSet(pair[1])

		if elf1.Intersect(elf2).Cardinality() > 0 {
			result++
		}
	}

	fmt.Println(result)
}
