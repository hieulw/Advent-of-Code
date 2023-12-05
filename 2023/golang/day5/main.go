package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	maps := make(map[string][]Map)
	seeds := []SeedRange{}

	// load first line to SeedRange
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		split := strings.Split(scanner.Text(), ": ")
		split = strings.Split(split[1], " ")
		for i := 0; i < len(split); i += 2 {
			seed, _ := strconv.Atoi(split[i])
			length, _ := strconv.Atoi(split[i+1])
			seeds = append(seeds, SeedRange{seed, length})
		}
	}

	// load next lines to maps
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
		converted := []int{}
		for _, s := range split {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			converted = append(converted, val)
		}
		maps[lastMap] = append(maps[lastMap], Map{destination: converted[0], source: converted[1], length: converted[2]})
	}

	lowestLocation := locationToSeed(maps, keyOrder, seeds)
	fmt.Println(lowestLocation)
}

type (
	Map struct {
		destination int
		source      int
		length      int
	}
	SeedRange struct {
		start  int
		length int
	}
)

func (m *Map) lookup(sourceValue int) int {
	if sourceValue < m.source || sourceValue >= m.source+m.length {
		return sourceValue
	}
	return m.destination + (sourceValue - m.source)
}

func (m *Map) reverseLookup(destinationValue int) int {
	if destinationValue < m.destination || destinationValue >= m.destination+m.length {
		return destinationValue
	}
	return m.source + (destinationValue - m.destination)
}

func (s *SeedRange) contains(value int) bool {
	return value >= s.start && value < s.start+s.length
}

func locationToSeed(maps map[string][]Map, keyOrder []string, seeds []SeedRange) int {
	lowestLocation := 0
	slices.Reverse(keyOrder)
	for {
		lastResult := lowestLocation
		for _, key := range keyOrder {
			for _, m := range maps[key] {
				result := m.reverseLookup(lastResult)
				if result != lastResult {
					lastResult = result
					break
				}
				lastResult = result
			}
		}
		for _, s := range seeds {
			if s.contains(lastResult) {
				return lowestLocation
			}
		}
		lowestLocation++
	}
}
