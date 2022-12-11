package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
)

var lcm = 1

type operation = func(int) int

type monkey struct {
	div         int
	inspections int
	items       []int
	op          operation
	testFalse   int
	testTrue    int
}

func (m *monkey) clone() *monkey {
	c := *m

	c.items = make([]int, len(m.items))
	copy(c.items, m.items)

	return &c
}

func getMonkeyLevelBusiness(monkeys []*monkey, rounds int, worryLevelDiv int) int {
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for _, b := range m.items {
				m.inspections++

				var (
					newWorryLevel = (m.op(b) % lcm) / worryLevelDiv
					nextMonkey    int
				)

				if newWorryLevel%m.div == 0 {
					nextMonkey = m.testTrue
				} else {
					nextMonkey = m.testFalse
				}

				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newWorryLevel)
			}

			m.items = make([]int, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}

func getOperation(c rune, right string) (operation, error) {
	var (
		rightBigInt int
		rightOld    = false
	)

	if right == "old" {
		rightOld = true
	} else {
		var err error

		rightBigInt, err = strconv.Atoi(right)
		if err != nil {
			return nil, fmt.Errorf("could not parse %q as integer: %v", right, err)
		}
	}

	switch c {
	case '+':
		return func(i int) int { return i + rightBigInt }, nil
	case '-':
		return func(i int) int { return i - rightBigInt }, nil
	case '*':
		if rightOld {
			return func(i int) int { return i * i }, nil
		}

		return func(i int) int { return i * rightBigInt }, nil
	}

	return nil, fmt.Errorf("unhandled operation %q", c)
}

func main() {
	var (
		monkeysPart1 = make([]*monkey, 0)
		r            = bufio.NewReader(os.Stdin)
	)

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	for {
		// Header

		if _, err := r.ReadString('\n'); err != nil {
			log.Fatalf("Could not read header: %v", err)
		}

		m := &monkey{}

		monkeysPart1 = append(monkeysPart1, m)

		// Starting items

		line, err := r.ReadString('\n')
		if err != nil {
			log.Panicf("Could not read line: %v", err)
		}

		listStart := strings.Index(line, ": ") + 1
		listString := line[listStart+1 : len(line)-1]

		var items []string

		if strings.Contains(listString, ",") {
			items = strings.Split(listString, ", ")
		} else {
			items = []string{listString}
		}

		m.items = make([]int, 0, len(items))

		for _, elem := range items {
			i, err := strconv.Atoi(elem)
			if err != nil {
				log.Panicf("Could not parse element %q: %v", elem, err)
			}

			m.items = append(m.items, i)
		}

		// Operation

		var (
			c     rune
			right string
		)

		if _, err = fmt.Fscanf(r, "  Operation: new = old %c %s\n", &c, &right); err != nil {
			log.Panicf("Could not parse operation line %q: %v", line, err)
		}

		op, err := getOperation(c, right)
		if err != nil {
			log.Panicf("Could not create operation: %v", err)
		}

		m.op = op

		// Test
		if _, err = fmt.Fscanf(r, "  Test: divisible by %d\n", &m.div); err != nil {
			log.Panicf("Could not parse test line %q: %v", line, err)
		}

		lcm *= m.div

		// If true

		if _, err = fmt.Fscanf(r, "    If true: throw to monkey %d\n", &m.testTrue); err != nil {
			log.Panicf("Could not parse if true line: %v", err)
		}

		// If false

		if _, err = fmt.Fscanf(r, "    If false: throw to monkey %d\n", &m.testFalse); err != nil {
			log.Panicf("Could not parse if false line: %v", err)
		}

		if _, err = r.ReadString('\n'); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Panicf("Error while reading an empty line: %v", err)
		}
	}

	monkeysPart2 := make([]*monkey, 0, len(monkeysPart1))

	for _, m := range monkeysPart1 {
		monkeysPart2 = append(monkeysPart2, m.clone())
	}

	log.Println("Part 1:", getMonkeyLevelBusiness(monkeysPart1, 20, 3))
	log.Println("Part 2:", getMonkeyLevelBusiness(monkeysPart2, 10000, 1))
}
