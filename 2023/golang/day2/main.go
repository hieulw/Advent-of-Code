package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	possible := true
	threshold := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		gameId, err := strconv.Atoi(strings.Fields(split[0])[1])
		if err != nil {
			panic(err)
		}
		for _, set := range strings.Split(split[1], "; ") {
			for _, cube := range strings.Split(set, ", ") {
				split := strings.Split(cube, " ")
				color := split[1]
				number, err := strconv.Atoi(split[0])
				if err != nil {
					panic(err)
				}
				if number > threshold[color] {
					possible = false
				}
			}
		}
		if possible {
			sum += gameId
		}
		possible = true
	}
	fmt.Println(sum)
}
