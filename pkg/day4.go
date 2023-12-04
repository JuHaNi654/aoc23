package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type scratchcards struct {
	winning []string
	current []string
}

func newCards(i, j []string) *scratchcards {
	return &scratchcards{
		winning: i,
		current: j,
	}
}

func contains[T comparable](s []T, i T) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func loadCards(row string, cards *[]*scratchcards) {
	values := strings.Split(row, "|")
	win := strings.Fields(values[0])
	curr := strings.Fields(values[1])

	group := newCards(win, curr)
	*cards = append(*cards, group)
}

func day4part1(cards *[]*scratchcards) int {
	result := 0
	current := 0

	for _, group := range *cards {
		for _, number := range group.current {
			if contains(group.winning, number) {
				if current == 0 {
					current = 1
					continue
				}
				current = current << 1
			}
		}

		result += current
		current = 0
	}

	return result
}

func day4part2(cards *[]*scratchcards) int {
	stats := make([]int, len(*cards))
	for i := 0; i < len(*cards); i++ {
		stats[i] = 1
	}

	for i, group := range *cards {
		nextItem := i + 1
		for _, number := range group.current {
			if contains(group.winning, number) {
				stats[nextItem] = stats[nextItem] + stats[i]
				nextItem++
			}
		}
	}

	result := 0
	for _, val := range stats {
		result += val
	}

	return result
}

func Day4(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	var cards []*scratchcards
	result := 0
	for s.Scan() {
		card := strings.Split(s.Text(), ": ")
		loadCards(card[1], &cards)
	}

	result = day4part2(&cards)
	fmt.Println("Result:", result)
}
