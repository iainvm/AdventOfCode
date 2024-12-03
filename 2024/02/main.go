package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := ReadInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Count safe reports
	safeReportsTally := CountSafeReports(reports)

	fmt.Println("Safe Reports:", safeReportsTally)
}

func CountSafeReports(reports [][]int) int {
	safeReportsTally := 0
	for _, report := range reports {
		if IsSafeReport(report) {
			safeReportsTally++
		}
	}
	return safeReportsTally
}

// IsSafeReport returns true if the report is safe
func IsSafeReport(report []int) bool {
	allowedDifferential := 3

	// Determine if the report is increasing or decreasing
	increasing := true
	if report[0] > report[1] {
		increasing = false
	}

	previous := report[0]
	for i := 1; i < len(report); i++ {
		// differential should be
		// positive if increasing
		// negative if decreasing
		differential := report[i] - previous

		// If differential is 0
		if differential == 0 {
			return false
		}

		// If differential is too great
		if int(math.Abs(float64(differential))) > allowedDifferential {
			return false
		}

		if increasing && differential < 0 {
			return false
		}
		if !increasing && differential > 0 {
			return false
		}

		previous = report[i]
	}

	return true
}

// ReadInput reads the input.txt file
// The file is a space separates list of reports
// 1 report per line
// Each report is a list of integers
func ReadInput() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int
		line := scanner.Text()
		for _, num := range strings.Split(line, " ") {
			value, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			report = append(report, value)
		}
		reports = append(reports, report)
	}

	return reports, nil
}
