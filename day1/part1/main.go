package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	data, _ := os.ReadFile("../input")
	var l, r []int
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		f := strings.Fields(line)
		a, _ := strconv.Atoi(f[0])
		b, _ := strconv.Atoi(f[1])
		l, r = append(l, a), append(r, b)
	}
	sort.Ints(l)
	sort.Ints(r)
	sum := 0
	for i := range l {
		sum += abs(r[i] - l[i])
	}
	fmt.Println(sum)
	elapsed := time.Since(start)
	fmt.Println("Çalışma süresi:", elapsed)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
