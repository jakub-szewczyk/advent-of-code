package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseReports() ([][]int, error) {
	data, err := os.ReadFile("2024/day-2/input.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	reports := [][]int{}
	for _, line := range lines {
		levels := strings.Fields(line)
		intLevels := []int{}
		for i := range levels {
			intLevel, _ := strconv.Atoi(levels[i])
			intLevels = append(intLevels, intLevel)
		}
		if len(intLevels) > 0 {
			reports = append(reports, intLevels)
		}
	}
	return reports, nil
}

func adjacentInRange(levels []int) []bool {
	safeLevels := []bool{}
	for i, level := range levels {
		if i+1 == len(levels) {
			break
		}
		nextLevel := levels[i+1]
		diff := math.Abs(float64(level - nextLevel))
		safeLevels = append(safeLevels, diff >= 1 && diff <= 3)
	}
	return safeLevels
}

func parseSafeReports() ([][]int, error) {
	reports, err := parseReports()
	if err != nil {
		return nil, err
	}
	safeReports := [][]int{}
	for _, levels := range reports {
		revLevels := make([]int, len(levels))
		copy(revLevels, levels)
		slices.Reverse(revLevels)
		if (slices.IsSorted(levels) || slices.IsSorted(revLevels)) && !slices.Contains(adjacentInRange(levels), false) {
			safeReports = append(safeReports, levels)
		}
	}
	return safeReports, nil
}

func parseDampenerSafeReports() ([][]int, error) {
	reports, err := parseReports()
	if err != nil {
		return nil, err
	}
	safeReports := [][]int{}
	for _, levels := range reports {
		revLevels := make([]int, len(levels))
		copy(revLevels, levels)
		slices.Reverse(revLevels)
		if (slices.IsSorted(levels) || slices.IsSorted(revLevels)) && !slices.Contains(adjacentInRange(levels), false) {
			safeReports = append(safeReports, levels)
		} else {
			for i := range levels {
				lvls := make([]int, len(levels))
				copy(lvls, levels)
				lvls = slices.Delete(lvls, i, i+1)
				revLvls := make([]int, len(lvls))
				copy(revLvls, lvls)
				slices.Reverse(revLvls)
				if (slices.IsSorted(lvls) || slices.IsSorted(revLvls)) && !slices.Contains(adjacentInRange(lvls), false) {
					safeReports = append(safeReports, lvls)
					break
				}
			}
		}
	}
	return safeReports, nil
}

func main() {
	fmt.Println("Day 2")

	fmt.Println("Part 1")
	reports, _ := parseSafeReports()
	fmt.Println(len(reports))

	fmt.Println("Part 2")
	reports, _ = parseDampenerSafeReports()
	fmt.Println(len(reports))
}
