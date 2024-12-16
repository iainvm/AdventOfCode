package main

import (
	"bufio"
	"fmt"
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
		} else {
			// Try all versions of report removing one element
			for i := 0; i < len(report); i++ {
				// Create a new report without the element at index i
				newReport := make([]int, len(report)-1)
				copy(newReport[:i], report[:i])
				copy(newReport[i:], report[i+1:])
				if IsSafeReport(newReport) {
					safeReportsTally++
					break
				}
			}
		}
	}
	return safeReportsTally
}

// IsSafeReport returns true if the report is safe
func IsSafeReport(report []int) bool {
	differentials := make([]int, len(report)-1)
	positive, negative := 0, 0

	// Calculate differentials
	for i := 0; i < len(report)-1; i++ {
		// Calculate differential
		diff := report[i+1] - report[i]
		differentials[i] = diff

		// Count signs
		if diff > 0 {
			positive++
		} else if diff < 0 {
			negative++
		}
	}

	if DifferentialsWithinLimits(differentials, -3, 3) {
		if positive == 0 || negative == 0 {
			return true
		}
	}

	return false
}

// DifferentialsWithinLimits returns true if all differentials are within the limits
func DifferentialsWithinLimits(differentials []int, lowerLimit, upperLimit int) bool {
	for _, diff := range differentials {
		if diff < lowerLimit || diff > upperLimit {
			return false
		}

		if diff == 0 {
			return false
		}
	}
	return true
}

// ReadInput reads the input.txt file
// The file is a space separates list of reports
// 1 report per line
// Each report is a list of integers
func ReadInput() ([][]int, error) {
	// file, err := os.Open("example_input.txt")
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
