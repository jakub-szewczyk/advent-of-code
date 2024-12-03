package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"jakubszewczyk.com.pl/utils"
)

func parseCols() ([]int, []int, error) {
	data, err := os.ReadFile("2024/day-1/input.txt")
	if err != nil {
		return nil, nil, err
	}
	ids := strings.Fields(string(data))
	col1 := []int{}
	col2 := []int{}
	for i, id := range ids {
		if id, _ := strconv.Atoi(id); i%2 == 0 {
			col1 = append(col1, id)
		} else {
			col2 = append(col2, id)
		}
	}
	return col1, col2, nil
}

func totalDist() (int, error) {
	col1, col2, err := parseCols()
	if err != nil {
		return 0, err
	}
	slices.Sort(col1)
	slices.Sort(col2)
	dist := 0
	for i := range col1 {
		dist += int(math.Abs(float64(col1[i] - col2[i])))
	}
	return dist, nil
}

func simScore() (int, error) {
	col1, col2, err := parseCols()
	if err != nil {
		return 0, err
	}
	score := 0
	for _, x := range col1 {
		score += x * utils.ElemCount(col2, func(y int) bool {
			return y == x
		})
	}
	return score, nil
}

func main() {
	fmt.Println("Day 1")

	dist, _ := totalDist()
	fmt.Println("Part 1", dist)

	score, _ := simScore()
	fmt.Println("Part 2", score)
}
