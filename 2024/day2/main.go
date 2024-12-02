package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Report = []int

func main() {
	reports := readReports("input")

	numSafeReports := 0
	numSafeReportsV2 := 0
	for _, r := range reports {
		if isSafe(r) {
			numSafeReports += 1
			numSafeReportsV2 += 1
		} else {
			// this seems like a pretty brute-force approach, not sure if there is a better way ...
			for without := 0; without < len(r); without += 1 {
				r2 := Report{}
				r2 = append(r2, r[0:without]...)
				r2 = append(r2, r[without+1:len(r)]...)
				if isSafe(r2) {
					numSafeReportsV2 += 1
					break // break inner loop, continue outer loop
				}
			}
		}
	}

	fmt.Println("Number of safe reports:", numSafeReports)
	fmt.Println("Number of safe reports with problem dampener (v2):", numSafeReportsV2)
}

// a report only counts as safe if both of the following are true:
//
//	The levels are either all increasing or all decreasing.
//	Any two adjacent levels differ by at least one and at most three.
func isSafe(r Report) bool {
	increasing := 0
	decreasing := 0
	for i := 0; i < len(r)-1; i += 1 {
		diff := r[i+1] - r[i]
		if diff > 0 {
			increasing += 1
		}
		if diff < 0 {
			decreasing += 1
		}
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
	}
	if increasing > 0 && decreasing > 0 {
		return false
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func readReports(filename string) []Report {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = ' '          // space-separated values
	reader.FieldsPerRecord = -1 // variable number of fields on each line
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	reports := make([]Report, 0, len(records))
	for _, record := range records {
		report := make(Report, 0, len(record))
		for _, r := range record {
			n, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}

	return reports
}
