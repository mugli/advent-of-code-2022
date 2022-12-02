package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	win  = 6
	draw = 3
	loss = 0
)

const (
	// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
	rock     = 1
	paper    = 2
	scissors = 3
)

func getMyScore(myMove, theirMove int) int {
	switch myMove {
	case rock:
		switch theirMove {
		case rock:
			return draw
		case paper:
			return loss
		case scissors:
			return win
		}
	case paper:
		switch theirMove {
		case rock:
			return win
		case paper:
			return draw
		case scissors:
			return loss
		}
	case scissors:
		switch theirMove {
		case rock:
			return loss
		case paper:
			return win
		case scissors:
			return draw
		}
	}

	panic("¯\\_(ツ)_/¯")
}

func main() {
	theirCodes := map[string]int{"A": rock, "B": paper, "C": scissors}
	myCodes := map[string]int{"X": rock, "Y": paper, "Z": scissors}

	result := 0

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		if len(line) < 3 {
			continue
		}

		show := strings.Split(line, " ")
		theirMove := theirCodes[show[0]]
		myMove := myCodes[show[1]]

		myScore := getMyScore(myMove, theirMove)
		result += myMove + myScore
	}

	fmt.Println(result)
}
