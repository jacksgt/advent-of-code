package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`mul\([0-9]+\,[0-9]+\)|do\(\)|don't\(\)`)

	matches := r.FindAll(data, -1)
	total := 0
	enabled := true
	for _, m := range matches {
		if string(m) == "do()" {
			enabled = true
			continue
		}
		if string(m) == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			total += calculate(string(m))
		}
	}

	fmt.Println("Uncorrupted mul instructions:", total)
}

func calculate(s string) int {
	var a, b int
	_, err := fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
	if err != nil {
		panic(err)
	}

	return a * b
}
