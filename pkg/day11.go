package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Just modify this value between part 1 and 2
const extra_steps = 999_999

type galaxyCoords struct {
	y, x int
}

func newGalaxyCoords(y, x int) *galaxyCoords {
	return &galaxyCoords{
		y: y,
		x: x,
	}
}

func checkVertical(row, col int, galaxies *[][]string, check *[]int) {
	if row > len(*galaxies)-1 {
		*check = append(*check, col)
		return
	}

	if (*galaxies)[row][col] == "#" {
		return
	}

	checkVertical(row+1, col, galaxies, check)
}

func checkHorizontal(galaxies *[][]string, check *[]int) {
	for i, row := range *galaxies {
		if !contains(row, "#") {
			*check = append(*check, i)
		}
	}
}

func day11part1(galaxies [][]string, r, c *[]int) int {
	steps := 0
	var coords []*galaxyCoords

	row_idx := 0
	for i, row := range galaxies {
		col_idx := 0
		if contains(*r, i) {
			row_idx += extra_steps
		}
		for j, val := range row {
			if contains(*c, j) {
				col_idx += extra_steps
			}

			if val == "#" {
				coords = append(coords, newGalaxyCoords(i+row_idx, j+col_idx))
			}
		}

	}

	for i, galaxy := range coords {
		for j := i + 1; j < len(coords); j++ {
			steps += (coords[j].y - galaxy.y)
			if galaxy.x > coords[j].x {
				steps += (galaxy.x - coords[j].x)
			} else {
				steps += (coords[j].x - galaxy.x)
			}
		}
	}
	return steps
}

func Day11(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	var galaxies [][]string
	var expand_col []int
	var expand_row []int

	defer file.Close()
	s := bufio.NewScanner(file)

	for s.Scan() {
		row := strings.Split(s.Text(), "")
		galaxies = append(galaxies, row)
	}

	for i, _ := range galaxies[0] {
		checkVertical(0, i, &galaxies, &expand_col)
	}
	checkHorizontal(&galaxies, &expand_row)
	result := day11part1(galaxies, &expand_row, &expand_col)

	fmt.Println("Result: ", result)
}
