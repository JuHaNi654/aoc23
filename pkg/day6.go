package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseRow(s string) []int {
	items := []int{}
	val := strings.Fields(s)

	for i := 1; i < len(val); i++ {
		num, _ := strconv.Atoi(val[i])
		items = append(items, num)
	}

	return items
}

// Just modify the input file, too lazy to compine numbers trough code
func day6part1(t, d int) int {
	result := 0
	for i := 0; i <= t; i++ {
		if (t-i)*i > d {
			result++
		}
	}

	return result
}

func Day6(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	defer file.Close()
	s := bufio.NewScanner(file)

	var (
		time     []int
		distance []int
	)

	result := 1
	for s.Scan() {
		if strings.Contains(s.Text(), "Time:") {
			time = parseRow(s.Text())
		}

		if strings.Contains(s.Text(), "Distance:") {
			distance = parseRow(s.Text())
		}
	}

	for i := 0; i < len(time); i++ {
		result *= day6part1(time[i], distance[i])
	}

	fmt.Println("Result: ", result)
}
