package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var result int64 = 0

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
		re := regexp.MustCompile(`\d{1}`)
		digitString := re.FindAllString(scanner.Text(), -1)
		calibration, err := strconv.ParseInt(digitString[0]+digitString[len(digitString)-1], 10, 0)
		if err != nil {
			panic(err)
		}
		result += calibration
	}
	fmt.Println(result)
}
