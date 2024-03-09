package main

import (
	"bufio"
	"os"
	"strings"
)

func CalculatePriorityOfSharableItems(first, second string) int {
	for _, char := range first {
		if !strings.Contains(second, string(char)) {
			continue
		}

		if priority := char - 96; priority > 0 {
			return int(priority) // a-z
		}
		return int(char - 38) // A-Z
	}
	return 0
}

func CalculatePriorityOfSharableItemsGroup(first, second, third string) int {
	for _, char := range first {
		if !(strings.Contains(second, string(char)) && strings.Contains(third, string(char))) {
			continue
		}

		if priority := char - 96; priority > 0 {
			return int(priority) // a-z
		}
		return int(char - 38) // A-Z
	}
	return 0
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	priorities1, priorities2 := 0, 0
	rucksackGroup := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack := scanner.Text()
		rucksackGroup = append(rucksackGroup, rucksack)

		index := len(rucksack) / 2
		firstCompartment := rucksack[:index]
		secondCompartment := rucksack[index:]
		priorities1 += CalculatePriorityOfSharableItems(firstCompartment, secondCompartment)

		if len(rucksackGroup) == 3 {
			priorities2 += CalculatePriorityOfSharableItemsGroup(rucksackGroup[0], rucksackGroup[1], rucksackGroup[2])
			rucksackGroup = []string{}
		}
	}
	println(priorities1, priorities2)
}
