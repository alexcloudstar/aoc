package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "part1.txt"

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

		if err != nil {
			continue
		}

		if dir == directions["RIGHT"] {
			pos += num
		}

		if dir == directions["LEFT"] {
			pos = pos - num
		}

		if pos < 0 {
			pos += 100
		}

		if pos > 99 {
			pos = pos % 100
		}

		if pos == 0 {
			clicks += 1
		}

	}

	fmt.Println(clicks)
}
