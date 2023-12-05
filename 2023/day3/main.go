package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var schematic []string
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}

	sum := 0
	for i, line := range schematic {
		numStr := ""
		partNumber := false
		for j := range line {
			if isDigit(line[j]) {
				if isAdjacentToSymbol(schematic, i, j) {
					partNumber = true
				}
				numStr += string(line[j])
			}
			// number separator by '.', symbol and end of line
			if line[j] == '.' || isSymbol(line[j]) || j == len(line)-1 {
				num, err := strconv.Atoi(numStr)
				if err == nil && partNumber {
					sum += num
				}
				numStr = ""
				partNumber = false
			}
		}
	}
	fmt.Println(sum)

	sumGearRatio := 0
	for i := range schematic {
		for j := range schematic[i] {
			if schematic[i][j] == '*' {
				part := findGearRatio(schematic, i, j)
				sumGearRatio += part[0] * part[1]
			}
		}
	}
	fmt.Println(sumGearRatio)
}

func isDigit(number byte) bool {
	return number >= '0' && number <= '9'
}

func isSymbol(char byte) bool {
	return !isDigit(char) && char != '.'
}

func isAdjacentToSymbol(schematic []string, line, row int) bool {
	for i := line - 1; i <= line+1; i++ {
		if i < 0 || i >= len(schematic) {
			continue
		}
		for j := row - 1; j <= row+1; j++ {
			if j < 0 || j >= len(schematic[i]) {
				continue
			}
			if isSymbol(schematic[i][j]) {
				return true
			}
		}
	}
	return false
}

func findGearRatio(schematic []string, line, row int) [2]int {
	set := make(map[[3]int]int)
	for i := line - 1; i <= line+1; i++ {
		if i < 0 || i >= len(schematic) {
			continue
		}
		for j := row - 1; j <= row+1; j++ {
			if j < 0 || j >= len(schematic[i]) {
				continue
			}
			if isDigit(schematic[i][j]) {
				number, start := findNumber(schematic, i, j)
				set[[3]int{number, line, start}]++ // number can be the same but not its position
			}
		}
	}
	parts := []int{}
	for k := range set {
		parts = append(parts, k[0])
	}
	if len(parts) == 2 {
		return [2]int{parts[0], parts[1]}
	}
	return [2]int{}
}

func findNumber(schematic []string, line, row int) (number, start int) {
	numStr := ""
	j := row - 1
	for j >= 0 && isDigit(schematic[line][j]) {
		j-- // backtrack to find where number starts
	}
	j++ // number starts from next character
	start = j
	for j < len(schematic[line]) && isDigit(schematic[line][j]) {
		numStr += string(schematic[line][j])
		j++
	}
	number, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return
}
