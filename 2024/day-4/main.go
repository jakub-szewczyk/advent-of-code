package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() (string, error) {
	data, err := os.ReadFile("2024/day-4/input.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func countXMAS(data string) (int, error) {
	lines := strings.Split(data, "\n")
	count := 0
	principalDiagonal := make(map[int][]byte)
	secondaryDiagonal := make(map[string][]byte)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		row := ""
		col := ""
		for j := 0; j < len(lines[i]); j++ {
			row += string(lines[i][j])
			col += string(lines[j][i])
			principalDiagonal[i+j] = append(principalDiagonal[i+j], lines[i][j])
			secondaryDiagonal[strconv.Itoa(i-j)] = append(secondaryDiagonal[strconv.Itoa(i-j)], lines[i][j])
		}
		count += strings.Count(row, "XMAS")
		count += strings.Count(row, "SAMX")
		count += strings.Count(col, "XMAS")
		count += strings.Count(col, "SAMX")
	}

	for _, bytes := range principalDiagonal {
		value := string(bytes)
		count += strings.Count(value, "XMAS")
		count += strings.Count(value, "SAMX")
	}

	for _, bytes := range secondaryDiagonal {
		value := string(bytes)
		count += strings.Count(value, "XMAS")
		count += strings.Count(value, "SAMX")
	}

	return count, nil
}

func main() {
	fmt.Println("Day 4")
	data, _ := parseInput()

	fmt.Println("Part 1")
	sum, _ := countXMAS(data)
	fmt.Println(sum)
}
