package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	earliestDeparture, busIds, err := loadNotes("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	fmt.Println("Part 1: ", getWaitTime(earliestDeparture, busIds))

}

func getWaitTime(earliestDepature int, busIds map[int]int) int {
	possibleBuses := make(map[int]int)
	var busDeltas []int
	for bus := range busIds {
		nextDeparture := math.Round(float64(earliestDepature)/float64(busIds[bus])) * float64(busIds[bus])
		delta := int(nextDeparture) - earliestDepature
		if delta > 0 {
			possibleBuses[delta] = busIds[bus]
			busDeltas = append(busDeltas, delta)
		}
	}
	sort.Ints(busDeltas)
	return busDeltas[0] * possibleBuses[busDeltas[0]]
}

func loadNotes(path string) (int, map[int]int, error) {
	var earliestTimestamp int
	busIds := make(map[int]int)

	notes, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return -1, nil, fmt.Errorf("unable to load notes :%s", err.Error())
	}
	defer func() {
		_ = notes.Close()
	}()

	notesScanner := bufio.NewScanner(notes)
	notesScanner.Split(bufio.ScanLines)
	index := 0
	for notesScanner.Scan() {
		entry := notesScanner.Text()
		if index == 0 {
			earliestTimestamp, err = strconv.Atoi(entry)
			if err != nil {
				return -1, nil, fmt.Errorf("bad data in notes: %s", err.Error())
			}
			index += 1
			continue
		}
		buses := strings.Split(entry, ",")
		for busIndex, bus := range buses {
			busId, err := strconv.Atoi(bus)
			if err != nil {
				continue
			}
			busIds[busIndex] = busId
		}
	}

	return earliestTimestamp, busIds, nil

}
