package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FileName = "input.txt"

func ReadLines(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, errors.New("we cannot open file")
	}

	contents := string(data)

	contents = strings.ReplaceAll(contents, "\n", "")
	contents = strings.ReplaceAll(contents, "\r", "")

	parts := strings.Split(contents, ",")

	var result []string

	for _, p := range parts {
		p = strings.TrimSpace(p)

		if p != "" {
			result = append(result, p)
		}
	}

	return result, nil
}

func part1() {
	lines, err := ReadLines(FileName)

	if err != nil {
		panic(err)
	}

	var invalidIds []int
	for i, _ := range lines {
		lId, err := strconv.Atoi(strings.Split(lines[i], "-")[0])
		if err != nil {
			panic("We cannot parse Left Index")
		}
		rId, err := strconv.Atoi(strings.Split(lines[i], "-")[1])

		if err != nil {
			panic("We cannot parse Right Index")
		}

		prevNum := lId

		for j := lId; j < rId; j++ {
			if strconv.Itoa(prevNum)[0:len(strings.Split(strconv.Itoa(prevNum), "")) / 2] ==
			strconv.Itoa(prevNum)[len(strings.Split(strconv.Itoa(prevNum), "")) / 2:] {
				invalidIds = append(invalidIds, prevNum)
			}

			prevNum = prevNum + 1
		}

		if strconv.Itoa(rId)[0:len(strings.Split(strconv.Itoa(rId), "")) / 2] ==
		strconv.Itoa(rId)[len(strings.Split(strconv.Itoa(rId), "")) / 2:] {
			invalidIds = append(invalidIds, rId)
		}
	}

	var sum int
	for _, value := range invalidIds {
		sum += value
	}

	fmt.Println(sum)
}


