package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var targetCycles = map[int]struct{}{
	20:  {},
	60:  {},
	100: {},
	140: {},
	180: {},
	220: {},
}

func main() {
	const crtWidth = 40

	var (
		part1      = 0
		part2Lines = make([]string, 0, 6)
		s          = bufio.NewScanner(os.Stdin)
		sb         strings.Builder
		X          = 1
	)

	sb.Grow(crtWidth)

	for i := 1; s.Scan(); i++ {
		carry := 0
		endTime := i
		words := strings.Split(s.Text(), " ")

		if len(words) == 2 {
			var err error

			carry, err = strconv.Atoi(words[1])
			if err != nil {
				log.Fatalf("Could not parse %q as integer: %v", words[1], err)
			}

			endTime++
		}

		for j := i; j <= endTime; j++ {
			if _, ok := targetCycles[j]; ok {
				part1 += j * X
			}

			pixelBeingDraw := (j - 1) % crtWidth

			if pixelBeingDraw >= X-1 && pixelBeingDraw <= X+1 {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}

			if sb.Len() == 40 {
				part2Lines = append(part2Lines, sb.String())
				sb.Reset()
			}
		}

		i = endTime
		X += carry
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error while reading input: %v", err)
	}

	log.Println("Part 1:", part1)
	log.Println("Part 2:")

	for _, line := range part2Lines {
		log.Println(line)
	}
}
