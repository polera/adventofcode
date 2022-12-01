package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	seatIds, err := getSeatIds("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	// Part 1: Get the highest seat Id
	sort.Ints(seatIds)
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
		row := pass[0:7]
		col := pass[7:]

		// Oh, hey, this is binary
		row = strings.ReplaceAll(strings.ReplaceAll(row, "F", "0"), "B", "1")
		col = strings.ReplaceAll(strings.ReplaceAll(col, "L", "0"), "R", "1")

		rowVal, _ := strconv.ParseInt(row, 2, 64)
		colVal, _ := strconv.ParseInt(col, 2, 64)

		seatId := int(rowVal*8 + colVal)
		seatIds = append(seatIds, seatId)
	}
	return seatIds, boardingPassScanner.Err()
}
