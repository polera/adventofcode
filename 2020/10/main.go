package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapters, err := getAdapters("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	part1 := calculateJoltDifferences(adapters)
	fmt.Println("Part 1 - 1-jolt diffs * 3-jolt diffs ", part1)

	sliceAdapters := []int{}
	for adapter := range adapters {
		sliceAdapters = append(sliceAdapters, adapter)
	}
	sort.Ints(sliceAdapters)
	part2 := adapterArrangements(sliceAdapters)
	fmt.Println("Part 2: ", part2)
}

// THIS TOOK A REALLY LONG TIME TO FIGURE OUT!!!
// Simple solution was to sort adapters and accumulate connections a we go
// Last adapter will have accumulated the number of connection combinations
func adapterArrangements(adapters []int) int {
	connections := map[int]int{
		0: 1, // set 0 key to 1 in order to capture 1, 2 or 3 connected to 0 as a valid connection
	}
	for _, adapter := range adapters {
		// accumulate connections for each increasing adapter
		connections[adapter] = connections[adapter-3] + connections[adapter-2] + connections[adapter-1]
	}
	return connections[adapters[len(adapters)-1]]
}

func calculateJoltDifferences(adapters map[int]bool) int {
	differences := map[int]int{
		3: 1, // start out with one here since the built in adapter is always a diff of 3
	}
	sourceRating := 0
	joltageRatings := []int{1, 2, 3}
	iter := 0
	numAdapters := len(adapters)
	for iter < numAdapters {
		for _, rating := range joltageRatings {
			if adapters[sourceRating+rating] {
				// Store jolt diff
				differences[rating] += 1
				// Move to next source rating
				sourceRating += rating
				break
			}
		}
		iter += 1
	}
	return differences[3] * differences[1]
}

func getAdapters(path string) (map[int]bool, error) {
	adapters := make(map[int]bool)
	adapterBag, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load adapter bag: %s", err.Error())
	}
	defer func() {
		_ = adapterBag.Close()
	}()

	adapterScanner := bufio.NewScanner(adapterBag)
	adapterScanner.Split(bufio.ScanLines)
	for adapterScanner.Scan() {
		entry := adapterScanner.Text()
		intEntry, err := strconv.Atoi(entry)
		if err != nil {
			return nil, fmt.Errorf("bad data in input file: %s", err.Error())
		}
		adapters[intEntry] = true
	}
	return adapters, nil
}
