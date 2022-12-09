package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math"
	"os"
)

type coord struct {
	x, y int
}

func dist(a, b coord) int {
	sqrt := math.Sqrt(
		math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2),
	)

	return int(sqrt)

	//diffX := math.Abs(float64(a.x)) - math.Abs(float64(a.x))
	//diffY := math.Abs(float64(a.y)) - math.Abs(float64(a.y))
	//
	//return int(diffX + diffY)
}

func getDist1(head, tail coord) coord {
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

	for _, c := range coords {
		log.Printf("dist(%v, %v) == %d", head, c, dist(head, c))
		if dist(c, head) == 1 {
			return c
		}
	}

	return coord{0, 0}
}

func main() {
	var (
		head    = coord{0, 0}
		tail    = head
		s       = bufio.NewScanner(os.Stdin)
		visited = make(map[coord]struct{})
	)

	for s.Scan() {
		line := s.Text()

		direction := line[0]
		value := int(line[2]) - '0'

		log.Printf("head moving to %c %d", direction, value)

		for i := 1; i < value; i++ {
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

			log.Printf("Head now at %v, tail at %v", head, tail)

			//tailNewX := tail.x
			//tailNewY := tail.y
			//
			//diffX := math.Abs(float64(head.x)) - math.Abs(float64(tail.x))
			//diffY := math.Abs(float64(head.y)) - math.Abs(float64(tail.y))
			//
			//log.Println("x diff is", diffX)
			//log.Println("y diff is", diffY)
			//
			//if diffX > 1 {
			//	if head.x > tail.x {
			//		tailNewX++
			//	} else {
			//		tailNewX--
			//	}
			//}
			//
			//if diffY > 1 {
			//	if head.y > tail.y {
			//		tailNewY++
			//	} else {
			//		tailNewY--
			//	}
			//}
			//
			//tail.x = tailNewX
			//tail.y = tailNewY

			if dist(head, tail) != 0 {
				tail = getDist1(head, tail)
				log.Println("tail moving to", tail)
			}

			visited[tail] = struct{}{}
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read input: %v", err)
	}

	log.Println("Part 1:", len(visited))
}
