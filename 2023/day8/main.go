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

	instruction := scanner.Text()
	start, end := "AAA", "ZZZ"
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
	}

	count := 0
	current := start
	i := 0
	for current != end {
		switch instruction[i] {
		case 'L':
			current = maps[current][0]
		case 'R':
			current = maps[current][1]
		}
		count++
		i = count % len(instruction)
	}
	fmt.Println(count)
}
