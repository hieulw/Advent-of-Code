package main

import (
	"fmt"
	"os"
)

func main() {
	stream_buffer, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	expected_len := 14
	for i := 0; i < len(stream_buffer)-expected_len; i++ {
		mark := Set{}
		for j := i; j < i+expected_len; j++ {
			mark[stream_buffer[j]] = struct{}{}
		}
		if len(mark) == expected_len {
			fmt.Println(i + expected_len)
			break
		}
	}
}

type Set map[interface{}]struct{}
