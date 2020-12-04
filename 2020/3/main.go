package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	tree = "#"
)

type TreeSearch struct {
	YStep    int
	XStep    int
	Position int
	Trees    int
}

type TreeSearches struct {
	Searches []*TreeSearch
}

func main() {

	// Part 1
	part1 := &TreeSearches{
		Searches: []*TreeSearch{
			{
				XStep:    3,
				YStep:    1,
				Position: 3,
			},
		},
	}
	numTrees1, err := countTrees("puzzle_input.txt", part1)
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}
	for _, search := range numTrees1.Searches {
		fmt.Println("Part 1 trees encountered: ", search.Trees)
	}

	// Part 2
	part2 := &TreeSearches{
		Searches: []*TreeSearch{
			{
				XStep:    1,
				YStep:    1,
				Position: 1,
			},
			{
				XStep:    3,
				YStep:    1,
				Position: 3,
			},
			{
				XStep:    5,
				YStep:    1,
				Position: 5,
			},
			{
				XStep:    7,
				YStep:    1,
				Position: 7,
			},
			{
				XStep:    1,
				YStep:    2,
				Position: 1,
			},
		},
	}
	numTrees2, err := countTrees("puzzle_input.txt", part2)
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	treesEncountered := 1
	for _, search := range numTrees2.Searches {
		treesEncountered *= search.Trees
	}
	fmt.Println("Part 2 trees encountered product: ", treesEncountered)

}

func countTrees(path string, searches *TreeSearches) (*TreeSearches, error) {

	slopeMap, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return searches, fmt.Errorf("unable to load slope map: %s", err.Error())
	}
	defer func() {
		_ = slopeMap.Close()
	}()

	slopeMapScanner := bufio.NewScanner(slopeMap)
	slopeMapScanner.Split(bufio.ScanLines)
	lineLength := 0
	currentLine := 0
	for slopeMapScanner.Scan() {
		entry := slopeMapScanner.Text()

		// All the lines are the same length, check this once
		if currentLine == 0 {
			lineLength = len(entry)
			currentLine += 1
			// No action is taken on the first line of input
			continue
		}

		for _, search := range searches.Searches {
			if currentLine%search.YStep == 0 {
				linePos := search.Position % lineLength
				if string(entry[linePos]) == tree {
					search.Trees += 1
				}
				search.Position += search.XStep
			}
		}
		currentLine += 1
	}

	return searches, slopeMapScanner.Err()
}
