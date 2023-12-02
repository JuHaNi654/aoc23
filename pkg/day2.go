package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(line string) int {
	var (
		id, val int
		label   string
	)

	split := strings.Split(line, ":")
	fmt.Sscanf(split[0], "Game %d", &id)

	for _, results := range strings.Split(split[1], ";") {
		items := strings.Split(results, ",")
		for _, item := range items {
			fmt.Sscanf(item, "%d %s", &val, &label)

			if label == "blue" && val > 14 ||
				label == "red" && val > 12 ||
				label == "green" && val > 13 {
				return 0
			}
		}
	}

	return id
}

func part2(line string) int {
	var (
		val              int
		red, blue, green int
		label            string
	)

	split := strings.Split(line, ":")

	for _, results := range strings.Split(split[1], ";") {
		items := strings.Split(results, ",")
		for _, item := range items {
			fmt.Sscanf(item, "%d %s", &val, &label)

			if label == "blue" && val > blue {
				blue = val
			}

			if label == "red" && val > red {
				red = val
			}

			if label == "green" && val > green {
				green = val
			}
		}
	}

	return red * green * blue
}

func Day2(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	result := 0
	for s.Scan() {
		result += part2(s.Text())
	}

	fmt.Println("Result:", result)
}
