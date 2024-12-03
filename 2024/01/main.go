package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"

	mySlices "github.com/iainvm/AdventOfCode/shared/slices"
)

func main() {
	list1, list2, err := ReadInput()
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// calculate the distance
	distance := CalculateDistance(list1, list2)
	fmt.Println("Total Distance:", distance)

	// calculate the similarity
	similarity := CalculateSimilarity(list1, list2)
	fmt.Println("Total Similarity:", similarity)
}

func CalculateSimilarity(list1, list2 []int) int {
	list2Counts := mySlices.CountSlice(list2)

	similarity := 0
	for i := 0; i < len(list1); i++ {
		count, ok := list2Counts[list1[i]]
		if ok {
			similarity += list1[i] * count
		}
	}

	return similarity
}

func CalculateDistance(list1, list2 []int) int {
	slices.Sort(list1)
	slices.Sort(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return distance
}

// ReadInput reads the input.txt file and returns the two lists of ints
// The file has 2 columns separated by whitespace
func ReadInput() ([]int, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, x)
		list2 = append(list2, y)
	}
	return list1, list2, scanner.Err()
}
