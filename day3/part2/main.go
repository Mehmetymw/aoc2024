package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Option string

const (
	DONT Option = "don't()"
	DO   Option = "do()"
)

func main() {
	file, _ := os.ReadFile("../input")
	total := 0

	splitted := strings.Split(string(file), "mul(")[1:]
	lastIndex := 0
	for _, part := range splitted {
		if i := strings.Index(part, ")"); i != -1 {
			if nums := strings.SplitN(part[:i], ",", 2); len(nums) == 2 {
				v1, _ := strconv.Atoi(nums[0])
				v2, _ := strconv.Atoi(nums[1])
				total += v1 * v2
				lastIndex = i
			}
		}
	}
	fmt.Println(total)
}
