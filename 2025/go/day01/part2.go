package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "part2.txt"

func main() {
	directions := map[string]string{
		"LEFT":  "L",
		"RIGHT": "R",
	}

	pos := 50
	clicks := 0

	reader, err := os.Open(fileName)

	if err != nil {
		panic("An error occured reading file")
	}

	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	defer reader.Close()

	for _, line := range lines {
		dir := strings.Split(line, "")[0]
		num, err := strconv.Atoi(line[1:])
		start := pos

		hits := 0

		if err != nil {
			continue
		}

		if dir == directions["RIGHT"] {
			hits = (start + num) / 100
			pos = (((start + num) % 100) + 100) % 100
		}

		if dir == directions["LEFT"] {
			if start == 0 {
				hits = num / 100
			}

			if start != 0 && num >= start {
				hits = ((num - start) / 100) + 1
			}

			pos = (((start - num) % 100) + 100) % 100
		}

		clicks += hits

	}

	fmt.Println(clicks)
}
