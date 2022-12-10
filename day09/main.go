package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func dist(a, b coord) (int, int, int) {
	distX := abs(a.x - b.x)
	distY := abs(a.y - b.y)

	return distX, distY, distX + distY
}

func touches(a, b coord) bool {
	distX, distY, sum := dist(a, b)

	return sum <= 1 || (distX == 1 && distY == 1)
}

func getClosest(head, tail coord) coord {
	coords := []coord{
		{x: tail.x - 1, y: tail.y - 1},
		{x: tail.x - 1, y: tail.y},
		{x: tail.x - 1, y: tail.y + 1},
		{x: tail.x, y: tail.y - 1},
		{x: tail.x, y: tail.y + 1},
		{x: tail.x + 1, y: tail.y - 1},
		{x: tail.x + 1, y: tail.y},
		{x: tail.x + 1, y: tail.y + 1},
	}

	var (
		closest coord
		minDist = math.MaxInt
	)

	for _, c := range coords {
		if _, _, d := dist(c, head); d < minDist {
			minDist = d
			closest = c
		}
	}

	return closest
}

func main() {
	const ropeLength = 9

	var (
		head         = coord{}
		rope         = make([]*coord, ropeLength)
		s            = bufio.NewScanner(os.Stdin)
		tail         = head
		visitedPart1 = make(map[coord]struct{})
		visitedPart2 = make(map[coord]struct{})
	)

	for i := 0; i < ropeLength; i++ {
		rope[i] = &coord{}
	}

	for i := 0; s.Scan(); i++ {
		var (
			direction rune
			value     int
		)

		if _, err := fmt.Sscanf(s.Text(), "%c %d", &direction, &value); err != nil {
			log.Fatalf("Could not read line %d: %v", i+1, err)
		}

		for j := 0; j < value; j++ {
			switch direction {
			case 'L':
				head.x--
			case 'R':
				head.x++
			case 'U':
				head.y++
			case 'D':
				head.y--
			}

			if !touches(head, tail) {
				tail = getClosest(head, tail)
			}

			visitedPart1[tail] = struct{}{}

			currHead := head

			for k := 0; k < ropeLength; k++ {
				if !touches(currHead, *rope[k]) {
					newPos := getClosest(currHead, *rope[k])
					rope[k] = &newPos
				}

				currHead = *rope[k]

				if k == ropeLength-1 {
					visitedPart2[*rope[k]] = struct{}{}
				}
			}
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read input: %v", err)
	}

	log.Println("Part 1:", len(visitedPart1))
	log.Println("Part 2:", len(visitedPart2))
}
