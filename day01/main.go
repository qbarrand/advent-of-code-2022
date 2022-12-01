package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type listItem struct {
	next *listItem
	prev *listItem
	val  int
}

type list struct {
	bottom *listItem
	count  int
	size   int
	sum    int
	top    *listItem
}

func (l *list) addElement(val int) {
	var (
		curr = l.top
		prev *listItem
	)

	for i := 0; i < l.size; i++ {
		if curr == nil || val > curr.val {
			log.Printf("Adding %d at index %d", val, i)

			l.count++
			l.sum += val

			newElem := &listItem{
				val:  val,
				prev: prev,
			}

			if i == 0 {
				l.top = newElem
			} else {
				prev.next = newElem
			}

			if curr == nil {
				// we are placing this node at the end of the list
				l.bottom = newElem
			} else {
				newElem.next = curr
			}

			break
		}

		prev = curr
		curr = curr.next
	}

	log.Printf("Before removal")
	printList(l)

	if l.count > l.size {
		log.Printf("%v", l.bottom)
		l.count--
		l.sum -= l.bottom.val
		l.bottom = l.bottom.prev
		l.bottom.next = nil
	}

	log.Printf("After removal")
	printList(l)
}

func printList(l *list) {
	curr := l.top

	fmt.Printf("top --> ")

	for curr != nil {
		fmt.Printf(" %d --> ", curr.val)
		curr = curr.next
	}

	fmt.Printf(" end\n")
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	curr := 0
	list := &list{size: 3}

	for s.Scan() {
		line := s.Text()

		if line == "" {
			log.Printf("Adding %d", curr)
			list.addElement(curr)
			curr = 0

			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Could not parse %q as integer: %v", line, err)
		}

		curr += cal
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error while reading the input: %v", err)
	}

	log.Println("Part 1:", list.top.val)
	log.Println("Part 2:", list.sum)
}
