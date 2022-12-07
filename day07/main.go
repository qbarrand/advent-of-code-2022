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

type node interface {
	GetSize() int
}

type baseNode struct {
	name string
}

type fileNode struct {
	baseNode

	size int
}

func newFileNode(name string, size int) *fileNode {
	return &fileNode{
		baseNode: baseNode{name: name},
		size:     size,
	}
}

func (fn *fileNode) GetSize() int {
	return fn.size
}

type dirNode struct {
	baseNode

	children map[string]node
	parent   *dirNode
}

func newDirNode(name string, parent *dirNode) *dirNode {
	return &dirNode{
		baseNode: baseNode{name: name},
		children: make(map[string]node),
		parent:   parent,
	}
}

func (dn *dirNode) GetSize() int {
	size := 0

	for _, v := range dn.children {
		size += v.GetSize()
	}

	return size
}

func sumTree(node *dirNode, minSize int, cache map[*dirNode]int) int {
	size := 0

	if s, ok := cache[node]; ok {
		size = s
	} else {
		size = node.GetSize()
		cache[node] = size
	}

	return node.GetSize()
}

func main() {
	const maxDirSize = 100000

	var (
		s    = bufio.NewScanner(os.Stdin)
		root = newDirNode("/", nil)
		wd   = root
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
			wd = wd.children[name].(*dirNode)
		case line == "$ ls":
		case strings.HasPrefix(line, "dir "):
			items := strings.Split(line, " ")
			dirName := items[1]
			wd.children[dirName] = newDirNode(dirName, wd)
		default:
			var (
				name string
				size int
			)

			if _, err := fmt.Sscanf(line, "%d %s", &size, &name); err != nil {
				log.Fatalf("Could not read directory item: %v", err)
			}

			wd.children[name] = newFileNode(name, size)
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read the input: %v", err)
	}
}
