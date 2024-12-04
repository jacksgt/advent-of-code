package main

import (
	"bufio"
	"fmt"
	"os"
)

var searchStr = []byte("XMAS")

var searchBack = []byte("SAMX")

func main() {
	// file, err := os.Open("inputsmall.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := [][]byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Bytes())
	}

	var numMatches uint
	for i := range data {
		for j := range data[i] {
			numMatches = numMatches +
				try0d(data, i, j) +
				try45d(data, i, j) +
				try90d(data, i, j) +
				try135d(data, i, j) +
				try180d(data, i, j) +
				try225d(data, i, j) +
				try270d(data, i, j) +
				try315d(data, i, j)
		}
	}

	fmt.Printf("Found %d matches in the crossword puzzle\n", numMatches)

}

// top
func try0d(d [][]byte, i, j int) uint {
	if i-3 < 0 {
		return 0
	}
	buf := []byte{d[i][j], d[i-1][j], d[i-2][j], d[i-3][j]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

// top-right
func try45d(d [][]byte, i, j int) uint {
	if i-3 < 0 || j+3 >= len(d[i-3]) {
		return 0
	}
	buf := []byte{d[i][j], d[i-1][j+1], d[i-2][j+2], d[i-3][j+3]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

// right
func try90d(d [][]byte, i, j int) uint {
	if j+3 >= len(d[i]) {
		return 0
	}
	buf := []byte{d[i][j], d[i][j+1], d[i][j+2], d[i][j+3]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

// bottom-right
func try135d(d [][]byte, i, j int) uint {
	if i+3 >= len(d) || j+3 >= len(d[i+3]) {
		return 0
	}
	buf := []byte{d[i][j], d[i+1][j+1], d[i+2][j+2], d[i+3][j+3]}
	return equals(buf, searchStr) ///+ equals(buf, searchBack)
}

// bottom
func try180d(d [][]byte, i, j int) uint {
	if i+3 >= len(d) {
		return 0
	}
	buf := []byte{d[i][j], d[i+1][j], d[i+2][j], d[i+3][j]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

// bottom-left
func try225d(d [][]byte, i, j int) uint {
	if i+3 >= len(d) || j-3 < 0 {
		return 0
	}
	buf := []byte{d[i][j], d[i+1][j-1], d[i+2][j-2], d[i+3][j-3]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

// left
func try270d(d [][]byte, i, j int) uint {
	if j-3 < 0 {
		return 0
	}
	buf := []byte{d[i][j], d[i][j-1], d[i][j-2], d[i][j-3]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

// top-left
func try315d(d [][]byte, i, j int) uint {
	if i-3 < 0 || j-3 < 0 {
		return 0
	}
	buf := []byte{d[i][j], d[i-1][j-1], d[i-2][j-2], d[i-3][j-3]}
	return equals(buf, searchStr) //+ equals(buf, searchBack)
}

func equals(a []byte, b []byte) uint {
	if len(a) != len(b) {
		return 0
	}
	for i := range a {
		if a[i] != b[i] {
			return 0
		}
	}
	return 1 // found a match
}
