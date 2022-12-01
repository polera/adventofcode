package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
)

const (
	empty    = "L"
	occupied = "#"
	floor    = "."
)

func main() {
	initialLayout, err := loadSeatLayout("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	finalLayout := runSimulation(initialLayout, 1)
	occupiedSeats := 0
	for seat := range finalLayout {
		if finalLayout[seat] == occupied {
			occupiedSeats += 1
		}
	}
	fmt.Println("Part 1 - Occupied seats: ", occupiedSeats)

	finalLayout2 := runSimulation(initialLayout, 2)
	occupiedSeats2 := 0
	for seat := range finalLayout2 {
		if finalLayout2[seat] == occupied {
			occupiedSeats2 += 1
		}
	}
	fmt.Println("Part 2 - Occupied seats: ", occupiedSeats2)
}

func runSimulation(seats map[complex64]string, ruleSet int) map[complex64]string {
	newLayout := updateLayout(seats, ruleSet)
	for {
		if reflect.DeepEqual(newLayout, seats) {
			break
		}
		seats = newLayout
		newLayout = updateLayout(seats, ruleSet)
	}
	return newLayout
}

func updateLayout(seats map[complex64]string, ruleSet int) map[complex64]string {
	updatedLayout := make(map[complex64]string)
	for seat := range seats {
		updatedLayout[seat] = seatStatus(seats, seat, ruleSet)
	}
	return updatedLayout
}

func seatStatus(seats map[complex64]string, seat complex64, ruleSet int) string {

	// do not mutate floor positions
	if seats[seat] == floor {
		return seats[seat]
	}

	neighbors := []complex64{
		complex(-1.0, 0.0),  // above
		complex(-1.0, 1.0),  // above-right
		complex(0.0, 1.0),   // right
		complex(1.0, 1.0),   // below-right
		complex(1.0, 0.0),   // below
		complex(1.0, -1.0),  // below-left
		complex(0., -1.0),   // left
		complex(-1.0, -1.0), // above-left
	}

	// Part 1
	if ruleSet == 1 {
		occupiedSeats := 0
		for _, neighbor := range neighbors {
			if seats[seat+neighbor] == occupied {
				occupiedSeats += 1
			}
		}

		switch seats[seat] {
		case empty:
			if occupiedSeats == 0 {
				return occupied
			}
		case occupied:
			if occupiedSeats > 3 {
				return empty
			}
		}
	}

	// Part 2
	if ruleSet == 2 {
		occupiedSeats := 0
		for _, neighbor := range neighbors {
			for iter := 1.0; iter < math.Sqrt(float64(len(seats))); iter++ {
				checkPos := seat + (neighbor * complex(float32(iter), 0))
				if _, ok := seats[checkPos]; !ok {
					// invalid position, break out
					break
				}

				if seats[checkPos] == floor {
					// no seats in sight, keep looking
					continue
				}

				if seats[checkPos] == occupied {
					// record an occupied seat
					occupiedSeats += 1
				}
				// position is either an empty or occupied seat, break out
				break
			}
		}

		switch seats[seat] {
		case empty:
			if occupiedSeats == 0 {
				return occupied
			}
		case occupied:
			if occupiedSeats > 4 {
				return empty
			}
		}
	}

	// no change
	return seats[seat]
}

func loadSeatLayout(path string) (map[complex64]string, error) {
	seatLayout := make(map[complex64]string)
	layout, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load seat layout: %s", err.Error())
	}
	defer func() {
		_ = layout.Close()
	}()

	layoutScanner := bufio.NewScanner(layout)
	layoutScanner.Split(bufio.ScanLines)
	rowIndex := 0
	for layoutScanner.Scan() {
		row := layoutScanner.Text()
		positions := strings.Split(row, "")
		for index, position := range positions {
			seatLayout[complex(float32(rowIndex), float32(index))] = position
		}
		rowIndex += 1
	}
	return seatLayout, nil
}
