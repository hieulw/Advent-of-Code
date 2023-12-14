package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

type Galaxy struct {
	x, y int
}

func main() {
	universe := [][]rune{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		universe = append(universe, []rune(scanner.Text()))
	}

	universe = expandUniverse(universe)
	galaxies := findGalaxy(universe)
	sum := 0.0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += math.Abs(float64(galaxies[i].x-galaxies[j].x)) + math.Abs(float64(galaxies[i].y-galaxies[j].y))
		}
	}
	fmt.Printf("%.0f\n", sum)
}

func findGalaxy(universe [][]rune) []Galaxy {
	galaxies := []Galaxy{}
	for i := range universe {
		for j := range universe[i] {
			if universe[i][j] == '#' {
				galaxy := Galaxy{x: j, y: i}
				galaxies = append(galaxies, galaxy)
			}
		}
	}
	return galaxies
}

func expandUniverse(universe [][]rune) [][]rune {
	for i := 0; i < len(universe); i++ {
		if !slices.Contains(universe[i], '#') {
			universe = slices.Insert(universe, i, universe[i])
			i++
		}
	}
	for j := 0; j < len(universe[0]); j++ {
		foundGalaxy := false
		for i := range universe {
			if universe[i][j] == '#' {
				foundGalaxy = true
				break
			}
		}
		if foundGalaxy {
			continue
		}
		for i := range universe {
			universe[i] = slices.Insert(universe[i], j, '.')
		}
		j++
	}
	return universe
}
