package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digitWord = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func repaceRuneAtIndex(str string, replace rune, index int) string {
	return str[:index] + string(replace) + str[index+1:]
}

func replaceWordToNumber(str string) string {
	for index := range str {
		curr := str[index:]
		for number, word := range digitWord {
			if strings.HasPrefix(curr, word) {
				str = repaceRuneAtIndex(str, rune(strconv.Itoa(number + 1)[0]), index)
			}
		}
	}
	return str
}

func main() {
	var result int

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := replaceWordToNumber(scanner.Text())
		re := regexp.MustCompile(`\d{1}`)
		digitString := re.FindAllString(text, -1)
		calibration, err := strconv.Atoi(digitString[0] + digitString[len(digitString)-1])
		if err != nil {
			panic(err)
		}
		result += calibration
	}
	fmt.Println(result)
}
