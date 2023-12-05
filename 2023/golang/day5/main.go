package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	maps := make(map[string][]Map)
	seeds := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		split = strings.Split(split[1], " ")
		for _, s := range split {
			seed, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			seeds = append(seeds, seed)
		}
	}
	lastMap := ""
	keyOrder := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "map") {
			split := strings.Split(line, " map:")
			lastMap = split[0]
			keyOrder = append(keyOrder, lastMap)
			continue
		}
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		split2 := []int{}
		for _, s := range split {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			split2 = append(split2, val)
		}
		maps[lastMap] = append(maps[lastMap], Map{destination: split2[0], source: split2[1], length: split2[2]})
	}

	locations := []int{}
	for _, seed := range seeds {
		lastResult := seed
		for _, key := range keyOrder {
			for _, m := range maps[key] {
				result := m.lookup(lastResult)
				if result != lastResult {
					lastResult = result
					break
				}
				lastResult = result
			}
		}
		locations = append(locations, lastResult)
	}

	lowestLocation := math.MaxInt
	for _, location := range locations {
		lowestLocation = min(lowestLocation, location)
	}
	fmt.Println(lowestLocation)
}

type Map struct {
	destination int
	source      int
	length      int
}

func (m *Map) lookup(value int) int {
	if value < m.source || value >= m.source+m.length {
		return value
	}
	return m.destination + (value - m.source)
}
