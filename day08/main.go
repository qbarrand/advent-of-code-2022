package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

type coord struct {
	x, y int
}

func part1(matrix [][]rune) int {
	var (
		max      rune
		selected = make(map[coord]struct{})
	)

	for y := 1; y < len(matrix)-1; y++ {
		line := matrix[y]

		max = line[0]

		for x := 1; x < len(line)-1; x++ {
			if line[x] > max {
				max = line[x]
				selected[coord{x: x, y: y}] = struct{}{}
			}
		}

		max = line[len(line)-1]

		for x := len(line) - 2; x > 0; x-- {
			if line[x] > max {
				max = line[x]
				selected[coord{x: x, y: y}] = struct{}{}
			}
		}
	}

	for x := 1; x < len(matrix[0])-1; x++ {
		max = matrix[0][x]

		for y := 1; y < len(matrix)-1; y++ {
			if matrix[y][x] > max {
				max = matrix[y][x]
				selected[coord{x: x, y: y}] = struct{}{}
			}
		}

		max = matrix[len(matrix)-1][x]

		for y := len(matrix) - 2; y >= 0; y-- {
			if matrix[y][x] > max {
				max = matrix[y][x]
				selected[coord{x: x, y: y}] = struct{}{}
			}
		}
	}

	return (len(matrix)-1)*2 + (len(matrix[0])-1)*2 + len(selected)
}

func part2(matrix [][]rune) int {
	max := 0

	// Skip edges as their viewing distance is 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[y])-1; x++ {
			currentHeight := matrix[y][x]

			top := 0
			for i := y - 1; i >= 0; i-- {
				top++

				if matrix[i][x] >= currentHeight {
					break
				}
			}

			bottom := 0
			for i := y + 1; i < len(matrix); i++ {
				bottom++

				if matrix[i][x] >= currentHeight {
					break
				}
			}

			left := 0
			for i := x - 1; i >= 0; i-- {
				left++

				if matrix[y][i] >= currentHeight {
					break
				}
			}

			right := 0
			for i := x + 1; i < len(matrix[y]); i++ {
				right++

				if matrix[y][i] >= currentHeight {
					break
				}
			}

			if score := top * bottom * left * right; max < score {
				max = score
			}
		}
	}

	return max
}

func main() {
	var (
		matrix = make([][]rune, 0)
		s      = bufio.NewScanner(os.Stdin)
	)

	for s.Scan() {
		line := s.Text()
		charLine := make([]rune, 0, len(line))

		for _, c := range line {
			charLine = append(charLine, c-'0')
		}

		matrix = append(matrix, charLine)
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read input: %v", err)
	}

	log.Println("Part 1:", part1(matrix))
	log.Println("Part 2:", part2(matrix))
}
