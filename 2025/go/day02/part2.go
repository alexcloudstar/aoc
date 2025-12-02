package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := ReadLines(FileName)

	if err != nil {
		panic(err)
	}

	var invalidIds []int

	for i, _ := range lines {
		lNum, err := strconv.Atoi(strings.Split(lines[i], "-")[0])
		if err != nil {
			panic("We cannot parse Left Index")
		}
		rNum, err := strconv.Atoi(strings.Split(lines[i], "-")[1])

		if err != nil {
			panic("We cannot parse Right Index")
		}

		for n := lNum; n <= rNum; n++ {
			if isInvalid(n) {
				invalidIds = append(invalidIds, n)
			}
		}
	}

	var sum int
	for _, value := range invalidIds {
		sum += value
	}

	fmt.Println(sum)
}

func isInvalid(num int) bool {
	s := strconv.Itoa(num)
	l := len(s)

	for b := 1; b <= l / 2; b++ {
		if l % b != 0 {
			continue
		}

		block := s[0:b]
		count := l / b
		repeated := strings.Repeat(block, count)

		if repeated == s {
			return true
		}
	}	

	return false
}
