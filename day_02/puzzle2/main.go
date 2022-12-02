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

func getMyMove(theirMove, myScore int) int {
	switch myScore {
	case win:
		switch theirMove {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	case draw:
		switch theirMove {
		case rock:
			return rock
		case paper:
			return paper
		case scissors:
			return scissors
		}
	case loss:
		switch theirMove {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	}

	panic("¯\\_(ツ)_/¯")
}

func main() {
	theirCodes := map[string]int{"A": rock, "B": paper, "C": scissors}
	myScoreCodes := map[string]int{"X": loss, "Y": draw, "Z": win}

	result := 0

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())

		if len(line) < 3 {
			continue
		}

		show := strings.Split(line, " ")
		theirMove := theirCodes[show[0]]
		myScore := myScoreCodes[show[1]]

		myMove := getMyMove(theirMove, myScore)
		result += myMove + myScore
	}

	fmt.Println(result)
}
