package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	start, end int
}

func parsePair(s string) (pair, error) {
	pairElems := strings.Split(s, "-")

	if l := len(pairElems); l != 2 {
		return pair{}, fmt.Errorf("expected 2 elements, got %d", l)
	}

	left, err := strconv.Atoi(pairElems[0])
	if err != nil {
		return pair{}, fmt.Errorf("could not parse %s as int: %v", pairElems[0], err)
	}

	right, err := strconv.Atoi(pairElems[1])
	if err != nil {
		return pair{}, fmt.Errorf("could not parse %s as int: %v", pairElems[1], err)
	}

	return pair{start: left, end: right}, nil
}

func (p pair) containsPair(o pair) bool {
	return o.start >= p.start && o.end <= p.end
}

func (p pair) overlaps(o pair) bool {
	return (o.start >= p.start && o.start <= p.end) || (o.end >= p.start && o.end <= p.end)
}

func main() {
	var (
		part1 = 0
		part2 = 0
		s     = bufio.NewScanner(os.Stdin)
	)

	for s.Scan() {
		strRanges := strings.Split(s.Text(), ",")

		left, err := parsePair(strRanges[0])
		if err != nil {
			log.Fatalf("could not parse left pair: %v", err)
		}

		right, err := parsePair(strRanges[1])
		if err != nil {
			log.Fatalf("could not parse right pair: %v", err)
		}

		if left.containsPair(right) || right.containsPair(left) {
			part1++
		}

		if left.overlaps(right) || right.overlaps(left) {
			part2++
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("could not read input: %v", err)
	}

	log.Println("Part 1:", part1)
	log.Println("Part 2:", part2)
}
