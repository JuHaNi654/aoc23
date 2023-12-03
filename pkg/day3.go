package pkg

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coords struct {
	y, x int
}

var directions = []coords{
	coords{y: -1, x: 0},  // Top	  - Center
	coords{y: -1, x: 1},  // Top		- Right
	coords{y: 0, x: 1},   // Middle	- Right
	coords{y: 1, x: 1},   // Bottom - Right
	coords{y: 1, x: 0},   // Bottom - Center
	coords{y: 1, x: -1},  // Bottom - Left
	coords{y: 0, x: -1},  // Middle - Left
	coords{y: -1, x: -1}, // Top    - Left
}

func checkMarker(s *[][]string, c coords) bool {
	if c.y < 0 || c.x < 0 {
		return false
	}

	if c.y > len(*s)-1 || c.x > len((*s)[0])-1 {
		return false
	}

	re := regexp.MustCompile(`[0-9]|\.`)
	item := (*s)[c.y][c.x]
	return !re.MatchString(item)
}

func d3part1(s *[][]string) []string {
	parts := []string{}
	part := ""
	found := false

	for i, row := range *s {
		for j, col := range row {
			r := rune(col[0])
			if r < 48 || r > 57 {
				if found {
					parts = append(parts, part)
				}

				part = ""
				found = false
				continue
			}

			part += col
			for _, direction := range directions {
				y := i + direction.y
				x := j + direction.x

				if checkMarker(s, coords{y: y, x: x}) {
					found = true
				}
			}
		}
	}

	return parts
}

func Day3(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	defer file.Close()

	var schematic [][]string
	s := bufio.NewScanner(file)
	result := 0
	for s.Scan() {
		r := strings.Split(s.Text(), "")
		schematic = append(schematic, r)
	}

	parts := d3part1(&schematic)
	for _, part := range parts {
		val, _ := strconv.Atoi(part)
		result += val
	}

	fmt.Println("Result:", result)
}
