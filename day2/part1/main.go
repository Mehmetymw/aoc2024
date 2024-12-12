package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Check if a report is safe without any removal
func isSafe(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] >= report[i-1] {
			return false
		}
	}
	return true
}

// Check if removing one level makes the report safe
func isSafeWithDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		// Create a new slice without the i-th element
		temp := append(report[:i], report[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}

// Parse a line into an integer slice
func parseReport(line string) ([]int, error) {
	parts := strings.Fields(line)
	report := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		report[i] = num
	}
	return report, nil
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		report, err := parseReport(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			return
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Analyze reports
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) || isSafeWithDampener(report) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}
