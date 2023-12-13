package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	visited = map[[2]int]bool{}
	tiles   = [][]rune{}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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
	checkPath(start[0], start[1])
	fmt.Println(len(visited) / 2)

	countInside := 0
	for i := range tiles {
		for j := range tiles[i] {
			if isInsideTile(i, j) {
				// fmt.Println(string(tiles[i][j]), i, j)
				countInside++
			}
		}
	}
	fmt.Println(countInside)
}

func isVisited(i, j int) bool {
	_, ok := visited[[2]int{i, j}]
	return ok
}

func isInsideTile(i, j int) bool {
	if isVisited(i, j) {
		return false
	}
	countCross := 0
	for k := 0; k < j; k++ {
		if strings.ContainsRune("F7|", tiles[i][k]) && isVisited(i, k) { // Ray-casting algorithm
			countCross++
		}
	}
	return countCross%2 == 1
}

func checkPath(i, j int) {
	border := ""
	if isValidPath(i-1, j, "F7|") { // North
		border += "N"
		step(i, j, i-1, j)
	}
	if isValidPath(i+1, j, "JL|") { // South
		border += "S"
		step(i, j, i+1, j)
	}
	if isValidPath(i, j-1, "FL-") { // West
		border += "W"
		step(i, j, i, j-1)
	}
	if isValidPath(i, j+1, "J7-") { // East
		border += "E"
		step(i, j, i, j+1)
	}
	switch border {
	case "NS":
		tiles[i][j] = '|'
	case "WE":
		tiles[i][j] = '-'
	case "NW":
		tiles[i][j] = 'J'
	case "NE":
		tiles[i][j] = 'L'
	case "SE":
		tiles[i][j] = 'F'
	case "SW":
		tiles[i][j] = '7'
	}
	fmt.Println(string(tiles[i][j]), i, j, border)
}

func isValidPath(i, j int, contain string) bool {
	if i < 0 || j < 0 || i >= len(tiles) || j >= len(tiles[i]) {
		return false
	}
	return strings.ContainsRune(contain, tiles[i][j])
}

func step(prevI, prevJ, i, j int) int {
	visited[[2]int{i, j}] = true
	switch tiles[i][j] {
	case 'J': // North and West
		if i-1 != prevI {
			return 1 + step(i, j, i-1, j)
		}
		if j-1 != prevJ {
			return 1 + step(i, j, i, j-1)
		}
	case '7': // South and West
		if i+1 != prevI {
			return 1 + step(i, j, i+1, j)
		}
		if j-1 != prevJ {
			return 1 + step(i, j, i, j-1)
		}
	case 'F': // South and East
		if i+1 != prevI {
			return 1 + step(i, j, i+1, j)
		}
		if j+1 != prevJ {
			return 1 + step(i, j, i, j+1)
		}
	case 'L': // North and East
		if i-1 != prevI {
			return 1 + step(i, j, i-1, j)
		}
		if j+1 != prevJ {
			return 1 + step(i, j, i, j+1)
		}
	case '-': // West and East
		if j-1 != prevJ {
			return 1 + step(i, j, i, j-1)
		}
		if j+1 != prevJ {
			return 1 + step(i, j, i, j+1)
		}
	case '|': // North and South
		if i+1 != prevI {
			return 1 + step(i, j, i+1, j)
		}
		if i-1 != prevI {
			return 1 + step(i, j, i-1, j)
		}
	}
	return 1
}
