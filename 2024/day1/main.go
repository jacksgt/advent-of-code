package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	listA, listB := readLists("input")

	// order lists
	// TODO: should implement sorting myself
	sort.Ints(listA)
	sort.Ints(listB)

	fmt.Println("Distance:", calculateDistance(listA, listB))
	fmt.Println("Similarity:", calculateSimilarity(listA, listB))
}

func calculateDistance(listA, listB []int) int {
	if len(listA) != len(listB) {
		panic("listA and listB must be the same size!")
	}

	if len(listA) == 0 {
		// end of recursion
		return 0
	}

	// calculate distance between smallest elements
	dist := listA[0] - listB[0]
	if dist < 0 {
		dist = dist * -1
	}

	// recursion with the rest of the list
	return dist + calculateDistance(listA[1:], listB[1:])
}

func calculateSimilarity(listA, listB []int) int {
	score := 0
	for _, a := range listA {
		occurences := countOccurences(listB, a)
		score += a * occurences
	}

	return score
}

// expects a sorted lit
func countOccurences(list []int, n int) int {
	o := 0
	for _, v := range list {
		if n == v {
			o += 1
		}
		// short-circuit requires a sorted list
		if v > n {
			break
		}
	}
	return o
}

func readLists(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	listA := []int{}
	listB := []int{}

	for {
		a := 0
		b := 0
		_, err := fmt.Fscanln(file, &a, &b)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		listA = append(listA, a)
		listB = append(listB, b)
	}

	return listA, listB
}
