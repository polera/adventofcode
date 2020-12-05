package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {

	seatIds, err := getSeatIds("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	// Part 1: Get the highest seat Id
	sort.Sort(sort.IntSlice(seatIds))
	fmt.Println("Part 1 - Highest seat ID: ", seatIds[len(seatIds)-1])

	// Part 2: Find my seat ID
	mySeatId, err := getMySeatId(seatIds)
	if err != nil {
		log.Fatalf("unable to find seat id: %s", err.Error())
	}
	fmt.Println("Part 2 - My seat ID: ", mySeatId)

}

func getMySeatId(seatIds []int) (int, error) {
	prev := 0
	for index, seat := range seatIds {
		if index == 0 {
			prev = seat
		}

		if seat-2 == prev {
			return seat - 1, nil
		}
		prev = seat
	}
	return -1, fmt.Errorf("seat not found in list of seatIds")
}

func getSeatIds(path string) ([]int, error) {
	var seatIds []int
	boardingPassFile, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load boarding passes: %s", err.Error())
	}
	defer func() {
		_ = boardingPassFile.Close()
	}()

	boardingPassScanner := bufio.NewScanner(boardingPassFile)
	boardingPassScanner.Split(bufio.ScanLines)
	for boardingPassScanner.Scan() {
		pass := boardingPassScanner.Text()
		rowLow := 0.0
		rowHigh := 127.0
		columnLow := 0.0
		columnHigh := 7.0

		splitPass := strings.Split(pass, "")
		for _, splitVal := range splitPass {
			switch splitVal {
			case "F":
				// take the lower half
				newHigh := math.Floor((rowLow + rowHigh) / 2)
				rowHigh = newHigh
			case "B":
				// take the upper half
				newLow := math.Ceil((rowLow + rowHigh) / 2)
				rowLow = newLow
			case "L":
				// take the lower half
				newHigh := math.Floor((columnLow + columnHigh) / 2)
				columnHigh = newHigh
			case "R":
				// take the upper half
				newLow := math.Ceil((columnLow + columnHigh) / 2)
				columnLow = newLow
			}
		}
		seatId := int(rowLow*8 + columnLow)
		seatIds = append(seatIds, seatId)
	}
	return seatIds, boardingPassScanner.Err()
}
