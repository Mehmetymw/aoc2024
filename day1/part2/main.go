package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input")
	var l []int
	r := make(map[int]int)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		fields := strings.Fields(line)
		v1, _ := strconv.Atoi(fields[0])
		v2, _ := strconv.Atoi(fields[1])
		l = append(l, v1)
		r[v2]++
	}

	sum := 0
	for _, v := range l {
		if rVal, ok := r[v]; ok {
			sum += v * rVal
		}
	}

	fmt.Println(sum)
}
