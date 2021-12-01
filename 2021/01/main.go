package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sweepReport, err := loadSweepReport("puzzle_input.txt")
	if err != nil {
		log.Fatalf("Unable to continue: %s", err.Error())
	}

	// Part 1
	fmt.Println("Part 1")
	increaseCount := getIncreases(sweepReport, 1)
	fmt.Printf("Found %d measurement increases in the report.\n", increaseCount)

	// Part 2
	fmt.Println("Part 2")
	groupIncreaseCount := getIncreases(sweepReport, 3)
	fmt.Printf("Found %d measurement increases in the report.\n", groupIncreaseCount)
}

func loadSweepReport(path string) ([]int, error) {
	entries := []int{}
	reportFile, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("Unable to load report: %s", err.Error())
	}
	defer func() {
		_ = reportFile.Close()
	}()

	reportScanner := bufio.NewScanner(reportFile)
	reportScanner.Split(bufio.ScanWords)
	for reportScanner.Scan() {
		entry, err := strconv.Atoi(reportScanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Invalid data in report: %s (%s)", entry, err.Error())
		}
		entries = append(entries, entry)
	}
	return entries, reportScanner.Err()
}

func getIncreases(sweeps []int, windowSize int) int {
	increaseCount := 0
	for index := range sweeps {
		// positions
		previousEnd := index + windowSize
		currentStart := index + 1
		currentEnd := index + windowSize + 1

		previous := sweeps[index:previousEnd]
		current := sweeps[currentStart:currentEnd]
		if isIncrease(previous, current) {
			increaseCount += 1
		}
	}
	return increaseCount
}

func isIncrease(previous, current []int) bool {
	previousSum := 0
	currentSum := 0
	for _, value := range previous {
		previousSum += value
	}
	for _, value := range current {
		currentSum += value
	}
	return currentSum > previousSum
}
