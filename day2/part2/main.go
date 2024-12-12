package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advanced checks if removing one element makes the report safe
func advanced(split []string) bool {
	for i := 0; i < len(split); i++ {
		// Create a copy of the array without the i-th element
		L := append(split[:i], split[i+1:]...)
		R := make([]int, len(L)-1)

		for j := 1; j < len(L); j++ {
			val1, _ := strconv.Atoi(L[j])
			val2, _ := strconv.Atoi(L[j-1])
			R[j-1] = val1 - val2
		}

		firstDiff, _ := strconv.Atoi(L[1])
		secondDiff, _ := strconv.Atoi(L[0])
		sign := firstDiff - secondDiff

		if sign < 0 {
			// Decreasing
			isValid := true
			for _, diff := range R {
				if diff > -1 || diff < -3 {
					isValid = false
					break
				}
			}
			if isValid {
				return true
			}
		} else if sign > 0 {
			// Increasing
			isValid := true
			for _, diff := range R {
				if diff < 1 || diff > 3 {
					isValid = false
					break
				}
			}
			if isValid {
				return true
			}
		}
	}
	return false
}

// Main function to process the input
func main() {
	filePath := "../input" // Update this to your input file path
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		split := strings.Fields(line)
		L := make([]int, len(split)-1)

		for i := 1; i < len(split); i++ {
			val1, _ := strconv.Atoi(split[i])
			val2, _ := strconv.Atoi(split[i-1])
			L[i-1] = val1 - val2
		}

		firstDiff, _ := strconv.Atoi(split[1])
		secondDiff, _ := strconv.Atoi(split[0])
		sign := firstDiff - secondDiff

		if sign < 0 {
			// Decreasing
			isValid := true
			for _, diff := range L {
				if diff > -1 || diff < -3 {
					isValid = false
					break
				}
			}
			if !isValid {
				if advanced(split) {
					safeCount++
				}
			} else {
				safeCount++
			}
		} else if sign > 0 {
			// Increasing
			isValid := true
			for _, diff := range L {
				if diff < 1 || diff > 3 {
					isValid = false
					break
				}
			}
			if !isValid {
				if advanced(split) {
					safeCount++
				}
			} else {
				safeCount++
			}
		} else {
			if advanced(split) {
				safeCount++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println(safeCount)
}
