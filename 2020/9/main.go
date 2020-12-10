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

	input, err := loadInput("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	hit := attackXmasData(input, 25, 25)
	fmt.Println("Part 1: ", hit)

	part2 := locateWeakness(input, hit)
	sort.Ints(part2)
	fmt.Println("Part 2: ", part2[0]+part2[len(part2)-1])

}

func locateWeakness(input []int, target int) []int {
	start := 0
	end := 1
	accumulated := input[start] + input[end]
	inputLen := len(input)
	for {
		if end >= inputLen {
			break
		}
		if accumulated == target {
			// done
			return input[start : end+1]
		}
		if accumulated < target {
			// add next index value to our accumulated value
			end += 1
			accumulated += input[end]
		}
		if accumulated > target {
			// slide up one index, subtract oldest index from accumulated value
			accumulated -= input[start]
			start += 1
		}
	}
	return nil
}

func attackXmasData(input []int, preamble int, lookBack int) int {
	start := 0
	end := preamble
	curr := end
	inputLen := len(input)
	for {
		if end >= inputLen {
			break
		}
		searchPop := make(map[int]bool)
		set := input[start:end]
		found := false
		for _, number := range set[len(set)-lookBack:] {
			searchPop[number] = true
			complement := searchPop[input[curr]-number]
			if complement {
				if float64(number) != float64(input[curr])/2.0 {
					// found complement, not our invalid number
					found = true
					break
				}
			}
		}
		if !found {
			// done
			return input[curr]
		}

		start += 1
		curr += 1
		end += 1
	}
	return -1
}

func loadInput(path string) ([]int, error) {
	var numbers []int
	xmasData, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load XMAS data: %s", err.Error())
	}
	defer func() {
		_ = xmasData.Close()
	}()

	xmasScanner := bufio.NewScanner(xmasData)
	xmasScanner.Split(bufio.ScanLines)
	for xmasScanner.Scan() {
		entry := xmasScanner.Text()
		intEntry, err := strconv.Atoi(entry)
		if err != nil {
			return nil, fmt.Errorf("bad data in input file: %s", err.Error())
		}
		numbers = append(numbers, intEntry)
	}
	return numbers, nil
}
