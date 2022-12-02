package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type outcome uint

const (
	ShouldLose outcome = 'X'
	ShouldDraw outcome = 'Y'
	ShouldWin  outcome = 'Z'
)

type hand struct {
	score        int
	losesAgainst *hand
	winsAgainst  *hand
}

var (
	rock     = &hand{score: 1}
	paper    = &hand{score: 2}
	scissors = &hand{score: 3}
)

func init() {
	rock.winsAgainst = scissors
	rock.losesAgainst = paper

	paper.winsAgainst = rock
	paper.losesAgainst = scissors

	scissors.winsAgainst = paper
	scissors.losesAgainst = rock
}

func charToHand(c uint8) *hand {
	switch c {
	case 'A', 'X':
		return rock
	case 'B', 'Y':
		return paper
	case 'C', 'Z':
		return scissors
	}

	return nil
}

func part1Score(theirs, ours *hand) int {
	if theirs == ours {
		return 3
	}

	if ours.winsAgainst == theirs {
		return 6
	}

	return 0
}

func part2Score(theirHand *hand, outcome outcome) int {
	switch outcome {
	case ShouldLose:
		return 0 + theirHand.winsAgainst.score
	case ShouldDraw:
		return 3 + theirHand.score
	case ShouldWin:
		return 6 + theirHand.losesAgainst.score
	}

	return 0
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	scorePart1 := 0
	scorePart2 := 0

	for s.Scan() {
		line := s.Text()

		ourChar := line[2]

		theirHand := charToHand(line[0])
		ourHandPart1 := charToHand(ourChar)

		scorePart1 += ourHandPart1.score
		scorePart1 += part1Score(theirHand, ourHandPart1)

		scorePart2 += part2Score(theirHand, outcome(ourChar))
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read the input: %v", err)
	}

	log.Println("Part 1:", scorePart1)
	log.Println("Part 2:", scorePart2)
}
