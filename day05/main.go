package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type stack struct {
	top *crate
}

func (s *stack) getPeekCrate() (*crate, error) {
	top := s.top

	if top == nil {
		return nil, errors.New("stack is empty")
	}

	return top, nil
}

func (s *stack) peek() (rune, error) {
	c, err := s.getPeekCrate()
	if err != nil {
		return 0, err
	}

	return c.name, nil
}

func (s *stack) pop() (rune, error) {
	c, err := s.getPeekCrate()
	if err != nil {
		return 0, err
	}

	s.top = c.next
	c.next = nil

	return c.name, nil
}

func (s *stack) popN(n int) ([]rune, error) {
	names := make([]rune, n)

	for i := 0; i < n; i++ {
		r, err := s.pop()
		if err != nil {
			return nil, fmt.Errorf("could not pop item %d: %v", i+1, err)
		}

		names[i] = r
	}

	return names, nil
}

func (s *stack) push(r rune) {
	s.top = &crate{
		name: r,
		next: s.top,
	}
}

func (s *stack) pushInOrder(names []rune) {
	for i := len(names) - 1; i >= 0; i-- {
		s.push(names[i])
	}
}

type crate struct {
	name rune
	next *crate
}

func main() {
	var (
		readHeader    = false
		reverseStacks [][]rune
		s             = bufio.NewScanner(os.Stdin)
		stacksPart1   []*stack
		stacksPart2   []*stack
	)

	for s.Scan() {
		line := s.Text()

		if !readHeader {
			if line == "" {
				readHeader = true

				// Setup stacks
				for _, rs := range reverseStacks {
					s1 := &stack{}
					s1.pushInOrder(rs)

					stacksPart1 = append(stacksPart1, s1)

					s2 := &stack{}
					s2.pushInOrder(rs)

					stacksPart2 = append(stacksPart2, s2)
				}

				continue
			}

			for i := 1; i < len(line); i += 4 {
				r := line[i]

				if r < 'A' || r > 'Z' {
					continue
				}

				stackNumber := (i - 1) / 4

				// Grow reverseStacks if necessary
				for j := len(reverseStacks) - 1; j < stackNumber; j++ {
					reverseStacks = append(reverseStacks, make([]rune, 0))
				}

				reverseStacks[stackNumber] = append(reverseStacks[stackNumber], rune(r))
			}

			continue
		}

		var n, from, to int

		if _, err := fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to); err != nil {
			log.Fatalf("Could not read instruction %q: %v", line, err)
		}

		// Part 1
		stackFrom := stacksPart1[from-1]
		stackTo := stacksPart1[to-1]

		for i := 0; i < n; i++ {
			r, err := stackFrom.pop()
			if err != nil {
				log.Fatalf("Could not pop from %d: %v", from, err)
			}

			stackTo.push(r)
		}

		// Part 2
		names, err := stacksPart2[from-1].popN(n)
		if err != nil {
			log.Fatalf("could not pop %d items in order: %v", n, err)
		}

		stacksPart2[to-1].pushInOrder(names)
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error while reading the input: %v", err)
	}

	var (
		sb1 strings.Builder
		sb2 strings.Builder
	)

	for i := 0; i < len(reverseStacks); i++ {
		r1, err := stacksPart1[i].peek()
		if err != nil {
			log.Fatalf("Could not peek from %d: %v", i+1, err)
		}

		sb1.WriteRune(r1)

		r2, err := stacksPart2[i].peek()
		if err != nil {
			log.Fatalf("Could not peek from %d: %v", i+1, err)
		}

		sb2.WriteRune(r2)
	}

	log.Println("Part 1:", sb1.String())
	log.Println("Part 2:", sb2.String())
}
