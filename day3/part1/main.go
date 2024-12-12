package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("../input")
	total := 0

	for _, part := range strings.Split(string(file), "mul(")[1:] {
		if i := strings.Index(part, ")"); i != -1 {
			if nums := strings.SplitN(part[:i], ",", 2); len(nums) == 2 {
				v1, _ := strconv.Atoi(nums[0])
				v2, _ := strconv.Atoi(nums[1])
				total += v1 * v2
			}
		}
	}
	fmt.Println(total)
}
