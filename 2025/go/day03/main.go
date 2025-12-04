package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "input.txt"

func main() {
	content := readLines(fileName)

	var greatestNum []int
	var sum int

	for idx := range content {
		greatestNum = append(greatestNum, getGreatestNumber(idx, content)) 
	}

	for _, val := range greatestNum {
		sum += val
	}

	fmt.Println(sum)
}

func getGreatestNumber(idx int, content []string) int {
	numbers := strings.Split(content[idx], "")

	firstBest, err := strconv.ParseInt(numbers[0], 32, 32)

	if err != nil {
		panic("We couldnt parse the number")
	}

	bestJolt := -1

	for i := 1; i < len(numbers); i++ {
		second, err := strconv.ParseInt(numbers[i], 32, 32)

		if err != nil {
			panic("We couldnt parse the number")
		}

		currVal := firstBest * 10 + second

		if currVal > int64(bestJolt) {
			bestJolt = int(currVal)
		}

		if second > firstBest {
			firstBest = second
		}
	}


	return bestJolt
}

func readLines(fileName string) []string {
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

	return lines
}
