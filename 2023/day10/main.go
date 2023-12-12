package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tiles := [][]rune{}
	start := [2]int{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, 'S') {
			start[0] = i
			start[1] = strings.IndexRune(line, 'S')
		}
		tiles = append(tiles, []rune(line))
		i++
	}
	for _, next := range checkPaths(tiles, start[0], start[1]) {
		fmt.Println(step(tiles, start[0], start[1], next[0], next[1]) / 2)
	}
}

func checkPaths(tiles [][]rune, i, j int) [][2]int {
	result := [][2]int{}
	if isValidPath(tiles, i-1, j, "F7|") { // top
		result = append(result, [2]int{i - 1, j})
	}
	if isValidPath(tiles, i+1, j, "JL|") { // bottom
		result = append(result, [2]int{i + 1, j})
	}
	if isValidPath(tiles, i, j-1, "FL-") { // left
		result = append(result, [2]int{i, j - 1})
	}
	if isValidPath(tiles, i, j+1, "J7-") { // right
		result = append(result, [2]int{i, j + 1})
	}
	return result
}

func isValidPath(tiles [][]rune, i, j int, contain string) bool {
	if i < 0 || j < 0 || i >= len(tiles) || j >= len(tiles[i]) {
		return false
	}
	return strings.ContainsRune(contain, tiles[i][j])
}

func step(tiles [][]rune, prevI, prevJ, i, j int) int {
	switch tiles[i][j] {
	case 'J': // North and West
		if i-1 != prevI {
			return 1 + step(tiles, i, j, i-1, j)
		}
		if j-1 != prevJ {
			return 1 + step(tiles, i, j, i, j-1)
		}
	case '7': // South and West
		if i+1 != prevI {
			return 1 + step(tiles, i, j, i+1, j)
		}
		if j-1 != prevJ {
			return 1 + step(tiles, i, j, i, j-1)
		}
	case 'F': // South and East
		if i+1 != prevI {
			return 1 + step(tiles, i, j, i+1, j)
		}
		if j+1 != prevJ {
			return 1 + step(tiles, i, j, i, j+1)
		}
	case 'L': // North and East
		if i-1 != prevI {
			return 1 + step(tiles, i, j, i-1, j)
		}
		if j+1 != prevJ {
			return 1 + step(tiles, i, j, i, j+1)
		}
	case '-': // West and East
		if j-1 != prevJ {
			return 1 + step(tiles, i, j, i, j-1)
		}
		if j+1 != prevJ {
			return 1 + step(tiles, i, j, i, j+1)
		}
	case '|': // North and South
		if i+1 != prevI {
			return 1 + step(tiles, i, j, i+1, j)
		}
		if i-1 != prevI {
			return 1 + step(tiles, i, j, i-1, j)
		}
	default:
	}
	return 1
}
