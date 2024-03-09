package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elfCaloriesList []int
	var topThreeCaloriesTotal int = 0

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var maxCalories, elfCalories int = 0, 0
	for scanner.Scan() {
		calories, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			if elfCalories > maxCalories {
				maxCalories = elfCalories
			}
			elfCaloriesList = append(elfCaloriesList, elfCalories)
			elfCalories = 0
		}
		elfCalories += int(calories)
	}
	sort.SliceStable(elfCaloriesList, func(i, j int) bool {
		return elfCaloriesList[i] > elfCaloriesList[j]
	})
	fmt.Println("max calories", maxCalories)

	for _, calories := range elfCaloriesList[:3] {
		topThreeCaloriesTotal += calories
	}
	fmt.Println(elfCaloriesList)
	fmt.Println("top three calories total", topThreeCaloriesTotal)
}
