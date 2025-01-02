package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseInput() (string, error) {
	data, err := os.ReadFile("2024/day-3/input.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func part1(data string) (int, error) {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	xs := r.FindAllStringSubmatch(data, -1)
	sum := 0
	for _, x := range xs {
		x1, err := strconv.Atoi(x[1])
		if err != nil {
			return 0, nil
		}
		x2, err := strconv.Atoi(x[2])
		if err != nil {
			return 0, nil
		}
		sum += x1 * x2
	}
	return sum, nil
}

func part2(data string) (int, error) {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	xs := r.FindAllStringSubmatch(data, -1)
	sum := 0
	enabled := true
	for _, x := range xs {
		if x[0] == "do()" {
			enabled = true
			continue
		}
		if x[0] == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		x1, err := strconv.Atoi(x[1])
		if err != nil {
			return 0, nil
		}
		x2, err := strconv.Atoi(x[2])
		if err != nil {
			return 0, nil
		}
		sum += x1 * x2
	}
	return sum, nil
}

func main() {
	fmt.Println("Day 3")
	data, _ := parseInput()

	fmt.Println("Part 1")
	sum, _ := part1(data)
	fmt.Println(sum)

	fmt.Println("Part 2")
	sum, _ = part2(data)
	fmt.Println(sum)
}
