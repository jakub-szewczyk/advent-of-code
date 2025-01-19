package main

import (
	"fmt"
	"os"
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
	secondaryDiagonal := make(map[int][]byte)

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
			secondaryDiagonal[i-j] = append(secondaryDiagonal[i-j], lines[i][j])
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

func countStars(data string) (int, error) {
	lines := strings.Split(data, "\n")
	count := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		for j := 0; j < len(lines[i]); j++ {
			if string(lines[i][j]) == "A" {
				// M M
				//
				//  A
				//
				// S S
				M00 := false
				M02 := false
				S20 := false
				S22 := false

				if i-1 >= 0 && j-1 >= 0 && string(lines[i-1][j-1]) == "M" {
					M00 = true
				}
				if i-1 >= 0 && j+1 < len(lines[i-1]) && string(lines[i-1][j+1]) == "M" {
					M02 = true
				}
				if i+1 < len(lines)-1 && j-1 >= 0 && string(lines[i+1][j-1]) == "S" {
					S20 = true
				}
				if i+1 < len(lines)-1 && j+1 < len(lines[i+1]) && string(lines[i+1][j+1]) == "S" {
					S22 = true
				}

				if M00 && M02 && S20 && S22 {
					count++
				}
			}

			if string(lines[i][j]) == "A" {
				// S S
				//
				//  A
				//
				// M M
				S00 := false
				S02 := false
				M20 := false
				M22 := false

				if i-1 >= 0 && j-1 >= 0 && string(lines[i-1][j-1]) == "S" {
					S00 = true
				}
				if i-1 >= 0 && j+1 < len(lines[i-1]) && string(lines[i-1][j+1]) == "S" {
					S02 = true
				}
				if i+1 < len(lines)-1 && j-1 >= 0 && string(lines[i+1][j-1]) == "M" {
					M20 = true
				}
				if i+1 < len(lines)-1 && j+1 < len(lines[i+1]) && string(lines[i+1][j+1]) == "M" {
					M22 = true
				}

				if S00 && S02 && M20 && M22 {
					count++
				}
			}

			if string(lines[i][j]) == "A" {
				// M S
				//
				//  A
				//
				// M S
				M00 := false
				S02 := false
				M20 := false
				S22 := false

				if i-1 >= 0 && j-1 >= 0 && string(lines[i-1][j-1]) == "M" {
					M00 = true
				}
				if i-1 >= 0 && j+1 < len(lines[i-1]) && string(lines[i-1][j+1]) == "S" {
					S02 = true
				}
				if i+1 < len(lines)-1 && j-1 >= 0 && string(lines[i+1][j-1]) == "M" {
					M20 = true
				}
				if i+1 < len(lines)-1 && j+1 < len(lines[i+1]) && string(lines[i+1][j+1]) == "S" {
					S22 = true
				}

				if M00 && S02 && M20 && S22 {
					count++
				}
			}

			if string(lines[i][j]) == "A" {
				// S M
				//
				//  A
				//
				// S M
				S00 := false
				M02 := false
				S20 := false
				M22 := false

				if i-1 >= 0 && j-1 >= 0 && string(lines[i-1][j-1]) == "S" {
					S00 = true
				}
				if i-1 >= 0 && j+1 < len(lines[i-1]) && string(lines[i-1][j+1]) == "M" {
					M02 = true
				}
				if i+1 < len(lines)-1 && j-1 >= 0 && string(lines[i+1][j-1]) == "S" {
					S20 = true
				}
				if i+1 < len(lines)-1 && j+1 < len(lines[i+1]) && string(lines[i+1][j+1]) == "M" {
					M22 = true
				}

				if S00 && M02 && S20 && M22 {
					count++
				}
			}
		}
	}

	return count, nil
}

func main() {
	fmt.Println("Day 4")
	data, _ := parseInput()

	fmt.Println("Part 1")
	sum, _ := countXMAS(data)
	fmt.Println(sum)

	fmt.Println("Part 2")
	sum, _ = countStars(data)
	fmt.Println(sum)
}
