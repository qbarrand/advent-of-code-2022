package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 'a' + 1
	}

	return int(r) - 'A' + 27
}

type charSet map[rune]struct{}

func (cs charSet) inter(o charSet) {
	for c := range cs {
		if _, ok := o[c]; !ok {
			delete(cs, c)
		}
	}
}

func main() {
	var (
		groupChars charSet
		part1      = 0
		part2      = 0
		s          = bufio.NewScanner(os.Stdin)
	)

	for i := 0; s.Scan(); i++ {
		line := s.Text()
		mid := len(line) / 2

		firstContainerItems := make(charSet, mid)
		allItems := make(charSet, len(line))

		for _, c := range line[:mid] {
			firstContainerItems[c] = struct{}{}
			allItems[c] = struct{}{}
		}

		foundDuplicate := false

		for _, c := range line[mid:] {
			allItems[c] = struct{}{}

			if _, ok := firstContainerItems[c]; ok && !foundDuplicate {
				part1 += getPriority(c)
				foundDuplicate = true
			}
		}

		if i%3 == 0 {
			groupChars = allItems
		} else {
			groupChars.inter(allItems)
		}

		if i%3 == 2 {
			for k := range groupChars {
				part2 += getPriority(k)
			}
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error while reading input: %v", err)
	}

	log.Println("Part 1:", part1)
	log.Println("Part 2:", part2)
}
