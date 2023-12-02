package pkg

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var values = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func convert(s string, skip, idx int, l *[]string) {
	if idx >= len(s) {
		return
	}

	item := rune(s[idx])

	if item >= 47 && item <= 57 {
		*l = append(*l, string(item))
		convert(s, (idx + 1), idx+1, l)
	} else {
		check := s[skip : idx+1]
		re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
		result := re.FindAllStringIndex(check, -1)

		if len(result) != 0 {
			r := result[0]
			skip = skip + (r[1] - 1)
			key := check[r[0]:r[1]]
			val := values[key]
			*l = append(*l, val)
		}

		convert(s, skip, idx+1, l)
	}
}

func Day1(f string) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	result := 0
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		var l []string
		convert(s.Text(), 0, 0, &l)
		val, _ := strconv.Atoi((l[0] + l[len(l)-1]))
		result += val
	}

	fmt.Println("Result:", result)
}
