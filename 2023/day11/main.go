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

	x, y := expandUniverse(universe)
	galaxies := findGalaxy(universe)
	galaxies = galaxyAfterExpand(galaxies, x, y, 1e6)
	sum := 0.0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += shortestPath(galaxies[i], galaxies[j])
		}
	}
	fmt.Printf("%.0f\n", sum)
}

func galaxyAfterExpand(galaxies []Galaxy, x, y []int, expandTime int) []Galaxy {
	expandTime -= 1
	for i := range galaxies {
		expandX, expandY := 0, 0
		for j := 0; j < len(x) && galaxies[i].x > x[j]; j++ {
			expandX += expandTime
		}
		for j := 0; j < len(y) && galaxies[i].y > y[j]; j++ {
			expandY += expandTime
		}
		galaxies[i].x += expandX
		galaxies[i].y += expandY
	}
	return galaxies
}

func shortestPath(a, b Galaxy) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
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

func expandUniverse(universe [][]rune) (xPoints, yPoints []int) {
	for i := 0; i < len(universe); i++ { // expand horizontally
		if !slices.Contains(universe[i], '#') {
			yPoints = append(yPoints, i)
		}
	}
	for j := 0; j < len(universe[0]); j++ { // expand vertically
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
		xPoints = append(xPoints, j)
	}
	return
}
