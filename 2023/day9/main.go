package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		numbers := [][]int{}
		numbers = append(numbers, []int{})
		split := strings.Fields(scanner.Text())
		for i := range split {
			num, err := strconv.Atoi(split[i])
			if err != nil {
				panic(err)
			}
			numbers[0] = append(numbers[0], num)
		}
		for i := 0; ; i++ {
			numbers = append(numbers, sequenceDiff(numbers[i]))
			minNumber, maxNumber := minMax(numbers[i+1])
			if minNumber == 0 && maxNumber == 0 {
				break
			}
		}
		sum += predictPrevious(numbers, 0)
	}
	fmt.Println(sum)
}

func predictNext(numbers [][]int, i, j int) int {
	last := numbers[i][j-1]
	if last == 0 {
		return 0
	}
	return last + predictNext(numbers, i+1, j-1)
}

func predictPrevious(numbers [][]int, i int) int {
	last := numbers[i][0]
	minNumber, maxNumber := minMax(numbers[i])
	if minNumber == 0 && maxNumber == 0 {
		return 0
	}
	return last - predictPrevious(numbers, i+1)
}

func sequenceDiff(numbers []int) []int {
	result := []int{}
	for i := 1; i < len(numbers); i++ {
		result = append(result, numbers[i]-numbers[i-1])
	}
	return result
}

func minMax(x []int) (int, int) {
	minInt, maxInt := x[0], x[0]
	for i := 1; i < len(x); i++ {
		minInt = min(minInt, x[i])
		maxInt = max(maxInt, x[i])
	}
	return minInt, maxInt
}
