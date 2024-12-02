package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type floatHeap []float64

func (h floatHeap) Len() int {
	return len(h)
}

func (h floatHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h floatHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *floatHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *floatHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var result1, result2 float64
	h1 := &floatHeap{}
	h2 := &floatHeap{}
	heap.Init(h1)
	heap.Init(h2)
	m1 := make(map[float64]int)
	m2 := make(map[float64]int)

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		n1, err := strconv.ParseFloat(fields[0], 64)
		if err != nil {
			panic(err)
		}
		n2, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			panic(err)
		}
		heap.Push(h1, n1)
		heap.Push(h2, n2)
		m1[n1]++
		m2[n2]++
	}
	for h1.Len() > 0 {
		n1 := heap.Pop(h1).(float64)
		n2 := heap.Pop(h2).(float64)
		result1 += math.Abs(n1 - n2)
	}
	for k, v := range m1 {
		result2 += k * float64(v) * float64(m2[k])
	}
	fmt.Printf("result: %.f\n", result1)
	fmt.Printf("result 2: %.f\n", result2)
}
