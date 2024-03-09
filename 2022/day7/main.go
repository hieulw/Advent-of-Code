package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	cwd := "/"
	tree := make(map[string]int)
	maxSize := 100000
	// var dirs, files map[string]int

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := strings.Split(line, " ")
		switch matches[0] {
		case "$":
			if matches[1] == "cd" {
				cwd = path.Join(cwd, matches[2])
			}
		default:
			currentPath := path.Join(cwd, matches[1])
			_, ok := tree[currentPath]
			size, err := strconv.Atoi(matches[0])
			if err != nil {
				size = 0
			}
			if !ok {
				if size < maxSize && size != 0 {
					tree[currentPath] = size
				}
			}
		}
	}
	for i, v := range tree {
		fmt.Println(i, v)
	}
}
