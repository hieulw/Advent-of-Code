package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scratchcards := make(map[int]int)
	totalCard := 0
	sum := 0
	count := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ": ")
		c := strings.Split(s[0], " ")
		cardNumber, _ := strconv.Atoi(c[len(c)-1])
		s = strings.Split(s[1], " | ")
		winNumbers := deleteEmptyElement(strings.Split(s[0], " "))
		numbers := deleteEmptyElement(strings.Split(s[1], " "))
		for _, number := range numbers {
			if slices.Contains(winNumbers, number) {
				count++
			}
		}
		scratchcards[cardNumber]++
		if count != 0 {
			points := int(math.Pow(2, float64(count-1)))
			sum += points
			for i := cardNumber + 1; i < cardNumber+count+1; i++ {
				scratchcards[i] += scratchcards[cardNumber]
			}
		}
		count = 0
	}
	for _, c := range scratchcards {
		totalCard += c
	}
	fmt.Println(sum)
	fmt.Println(totalCard)
}

func deleteEmptyElement(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
