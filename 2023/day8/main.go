package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var start []string
	instruction := scanner.Text()
	maps := map[string][2]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		split := strings.Split(line, " = ")
		current := split[0]
		split = strings.Split(split[1], ", ")
		left := strings.TrimLeft(split[0], "(")
		right := strings.TrimRight(split[1], ")")
		maps[current] = [2]string{left, right}
		if strings.HasSuffix(current, "A") {
			start = append(start, current)
		}
	}

	steps := []int{}
	for i := range start {
		steps = append(steps, countSteps(maps, instruction, start[i]))
	}
	// finding least common multiple of all steps
	lcmSteps := steps[0]
	for i := range steps {
		lcmSteps = lcm(lcmSteps, steps[i])
	}
	fmt.Println(lcmSteps)
}

func countSteps(maps map[string][2]string, instruction, start string) int {
	current := start
	count := 0
	i := 0
	for !strings.HasSuffix(current, "Z") {
		switch instruction[i] {
		case 'L':
			current = maps[current][0]
		case 'R':
			current = maps[current][1]
		}
		count++
		i = count % len(instruction)
	}
	return count
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
