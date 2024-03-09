package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func MapInt(list []string) [2]int {
	result := [2]int{}
	for i, v := range list {
		v, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result[i] = v
	}
	return result
}

func RangeContain(first, second string) bool {
	firstRange := MapInt(strings.Split(first, "-"))
	secondRange := MapInt(strings.Split(second, "-"))
	return firstRange[0] >= secondRange[0] && firstRange[1] <= secondRange[1]
}

func RangeOverlap(first, second string) bool {
	firstRange := MapInt(strings.Split(first, "-"))
	secondRange := MapInt(strings.Split(second, "-"))
	return !(firstRange[1] < secondRange[0] || firstRange[0] > secondRange[1])
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	countContain, countOverlap := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		if RangeContain(pair[0], pair[1]) || RangeContain(pair[1], pair[0]) {
			countContain++
		}
		if RangeOverlap(pair[0], pair[1]) || RangeOverlap(pair[1], pair[0]) {
			countOverlap++
		}
	}
	println(countContain, countOverlap)
}
