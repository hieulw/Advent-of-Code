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
	sumPower := 0
	power := 1
	possible := true
	threshold := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	max := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
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
				if max[color] < number {
					max[color] = number
				}
			}
		}
		if possible {
			sum += gameId
		}
		possible = true
		for _, number := range max {
			power *= number
		}
		sumPower += power
		power = 1
		max = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
	}
	fmt.Println("part 1:", sum)
	fmt.Println("part 2:", sumPower)
}
