package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ApplyRuleOne(first, second string) int {
	rule := map[string]int{"X": 1, "Y": 2, "Z": 3}
	concat := first + second
	switch concat {
	case "AX", "BY", "CZ": // draw
		return 3 + rule[second]
	case "BX", "CY", "AZ": // lose
		return 0 + rule[second]
	case "CX", "AY", "BZ": // win
		return 6 + rule[second]
	}
	return 0
}

func ApplyRuleTwo(first, second string) int {
	var rule map[string]int
	switch second {
	case "Y": // draw
		rule = map[string]int{"X": 0, "Y": 3, "Z": 6, "A": 1, "B": 2, "C": 3}
	case "X": // lose
		rule = map[string]int{"X": 0, "Y": 3, "Z": 6, "A": 3, "B": 1, "C": 2}
	case "Z": // win
		rule = map[string]int{"X": 0, "Y": 3, "Z": 6, "A": 2, "B": 3, "C": 1}
	}
	return rule[first] + rule[second]
}

func main() {
	/*
	   A - X : rock
	   B - Y : paper
	   C - Z : scissors
	*/
	var score, score2 int = 0, 0

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		first, second, _ := strings.Cut(scanner.Text(), " ")
		score += ApplyRuleOne(first, second)
		score2 += ApplyRuleTwo(first, second)
	}
	fmt.Println("total score", score, score2)
}
