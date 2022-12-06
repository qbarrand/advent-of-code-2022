package main

import (
	"bufio"
	"container/ring"
	"errors"
	"io"
	"log"
	"os"
)

type markerQueue struct {
	m    map[rune]int
	r    *ring.Ring
	size int
}

func newMarkerBuffer(size int) *markerQueue {
	return &markerQueue{
		m:    make(map[rune]int, size),
		r:    ring.New(size),
		size: size,
	}
}

func (mq *markerQueue) dequeue() {
	val := mq.r.Value.(rune)
	mq.m[val]--

	if mq.m[val] == 0 {
		delete(mq.m, val)
	}
}

func (mq *markerQueue) enqueue(r rune) {
	mq.m[r]++
	mq.r.Value = r
	mq.r = mq.r.Next()
}

func (mq *markerQueue) isMarker() bool {
	return len(mq.m) == mq.size
}

func main() {
	const (
		initValue       = -1
		part1MarkerSize = 4
		part2MarkerSize = 14
	)

	var (
		mq1        = newMarkerBuffer(part1MarkerSize)
		mq2        = newMarkerBuffer(part2MarkerSize)
		part1Found = false
		part2Found = false
		s          = bufio.NewScanner(os.Stdin)
	)

	s.Split(bufio.ScanRunes)

	for i := 0; s.Scan() && (!part1Found || !part2Found); i++ {
		char := rune(s.Text()[0])

		if i >= part1MarkerSize {
			mq1.dequeue()
		}

		if i >= part2MarkerSize {
			mq2.dequeue()
		}

		mq1.enqueue(char)
		mq2.enqueue(char)

		if !part1Found && mq1.isMarker() {
			log.Println("Part 1:", i+1)
			part1Found = true
		}

		if !part2Found && mq2.isMarker() {
			log.Println("Part 2:", i+1)
			part2Found = true
		}
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Could not read input: %v", err)
	}
}
