package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	stacks := []Stack{
		{items: []Item{'W', 'R', 'F'}},
		{items: []Item{'T', 'H', 'M', 'C', 'D', 'V', 'W', 'P'}},
		{items: []Item{'P', 'M', 'Z', 'N', 'L'}},
		{items: []Item{'J', 'C', 'H', 'R'}},
		{items: []Item{'C', 'P', 'G', 'H', 'Q', 'T', 'B'}},
		{items: []Item{'G', 'C', 'W', 'L', 'F', 'Z'}},
		{items: []Item{'W', 'V', 'L', 'Q', 'Z', 'J', 'G', 'C'}},
		{items: []Item{'P', 'N', 'R', 'F', 'W', 'T', 'V', 'C'}},
		{items: []Item{'J', 'W', 'H', 'G', 'R', 'S', 'V'}},
	}

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	tempStack := Stack{}
	for scanner.Scan() {
		instruction := scanner.Text()
		r := regexp.MustCompile(`(\d+)`)
		match := MapInt(r.FindAllString(instruction, -1))
		quantity, fromStack, toStack := match[0], match[1]-1, match[2]-1
		for i := 0; i < quantity; i++ {
			tempStack.Push(stacks[fromStack].Pop())
		}
		for i := 0; i < quantity; i++ {
			stacks[toStack].Push(tempStack.Pop())
		}
	}
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%c", stacks[i].Peek())
	}
}

func MapInt(list []string) []int {
	result := []int{0, 0, 0}
	for i, v := range list {
		v, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result[i] = v
	}
	return result
}

type (
	Item  interface{}
	Stack struct {
		items []Item
		mutex sync.Mutex
	}
)

func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	size := len(stack.items)

	if size == 0 {
		return nil
	}

	lastItem := stack.items[size-1]
	stack.items = stack.items[:size-1]

	return lastItem
}

func (stack *Stack) Peek() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	size := len(stack.items)

	if size == 0 {
		return nil
	}

	return stack.items[size-1]
}
