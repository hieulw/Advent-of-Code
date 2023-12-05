package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ": ")
		s = strings.Split(s[1], " | ")
		winNumbers := deleteEmptyElement(strings.Split(s[0], " "))
		numbers := deleteEmptyElement(strings.Split(s[1], " "))
		for _, number := range numbers {
			if slices.Contains(winNumbers, number) {
				count++
			}
		}
		if count != 0 {
			points := int(math.Pow(2, float64(count-1)))
			sum += points
		}
		count = 0
	}
	fmt.Println(sum)
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
