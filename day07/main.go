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

type dir struct {
	files  map[string]int
	dirs   map[string]*dir
	parent *dir
}

func newDir(parent *dir) *dir {
	return &dir{
		files:  make(map[string]int),
		dirs:   make(map[string]*dir),
		parent: parent,
	}
}

func sumTree(node *dir, sum, maxSize int, cache map[*dir]int) (int, int) {
	size := 0

	if s, ok := cache[node]; ok {
		size = s
	} else {
		for _, d := range node.dirs {
			childSize, childSum := sumTree(d, sum, maxSize, cache)
			size += childSize
			sum = childSum
		}

		for _, f := range node.files {
			size += f
		}

		cache[node] = size
	}

	if size <= maxSize {
		sum += size
	}

	return size, sum
}

func main() {
	var (
		nDirs = 0
		s     = bufio.NewScanner(os.Stdin)
		root  = newDir(nil)
		wd    = root
	)

	for s.Scan() {
		line := s.Text()

		switch {
		case line == "$ cd ..":
			wd = wd.parent
		case line == "$ cd /":
			wd = root
		case strings.HasPrefix(line, "$ cd "):
			items := strings.Split(line, " ")
			name := items[2]
			wd = wd.dirs[name]
		case line == "$ ls":
		case strings.HasPrefix(line, "dir "):
			items := strings.Split(line, " ")
			dirName := items[1]
			wd.dirs[dirName] = newDir(wd)
			nDirs++
		default:
			var (
				name string
				size int
			)

			if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err != nil {
				log.Fatalf("Could not read directory item: %v", err)
			}

			wd.files[name] = size
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read the input: %v", err)
	}

	cache := make(map[*dir]int, nDirs)

	totalSize, part1 := sumTree(root, 0, 100000, cache)

	log.Println("Part 1:", part1)

	const uninitialized = -1

	part2 := uninitialized

	for _, v := range cache {
		if 70000000-(totalSize-v) > 30000000 {
			if part2 == uninitialized || v < part2 {
				part2 = v
			}
		}
	}

	log.Println("Part 2:", part2)
}
