package main

import (
	"JuHaNi654/aoc23/pkg"
	"fmt"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	input := os.Args[2]
	day := os.Args[1]
	file := fmt.Sprintf("%s/%s", pwd, input)

	switch day {
	case "day1":
		pkg.Day1(file)
	case "day2":
		pkg.Day2(file)
	case "day3":
		pkg.Day3(file)
	default:
		fmt.Errorf("Day not found")
	}
}
