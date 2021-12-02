package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PlanEntry struct {
	Direction string
	Velocity  int
}

func main() {
	planCourse, err := loadPlan("puzzle_input.txt")
	if err != nil {
		log.Fatalf("Unable to continue: %s", err.Error())
	}

	// Part 1
	fmt.Println("Part 1")
	hPos, depth := runCourse(planCourse)
	fmt.Printf("h: %d d: %d\n", hPos, depth)
	fmt.Printf("Result: %d\n", hPos*depth)

	// Part 1
	fmt.Println("Part 2")
	p2hPos, p2depth := runCourseWithAim(planCourse)
	fmt.Printf("h: %d d: %d\n", p2hPos, p2depth)
	fmt.Printf("Result: %d\n", p2hPos*p2depth)
}

func loadPlan(path string) ([]PlanEntry, error) {
	entries := []PlanEntry{}
	planFile, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("Unable to load plan: %s", err.Error())
	}
	defer func() {
		_ = planFile.Close()
	}()

	reportScanner := bufio.NewScanner(planFile)
	reportScanner.Split(bufio.ScanLines)
	for reportScanner.Scan() {
		entry := strings.Split(reportScanner.Text(), " ")
		if len(entry) < 2 {
			return nil, fmt.Errorf("Invalid data in plan, line length less than 2.")
		}
		direction := entry[0]
		velocity, err := strconv.Atoi(entry[1])
		if err != nil {
			return nil, fmt.Errorf("Invalid data in plan: %s (%s)", entry, err.Error())
		}
		planEntry := PlanEntry{
			Direction: direction,
			Velocity:  velocity,
		}
		entries = append(entries, planEntry)
	}
	return entries, reportScanner.Err()
}

func runCourse(course []PlanEntry) (int, int) {
	horizontalPosition := 0
	depth := 0

	for _, entry := range course {
		horizontalPosition, depth = executeCommand(entry, horizontalPosition, depth)
	}
	return horizontalPosition, depth
}

func executeCommand(entry PlanEntry, horizontalPosition, depth int) (int, int) {
	switch entry.Direction {
	case "forward":
		horizontalPosition += entry.Velocity
	case "up":
		depth -= entry.Velocity
	case "down":
		depth += entry.Velocity
	}

	return horizontalPosition, depth
}

func runCourseWithAim(course []PlanEntry) (int, int) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, entry := range course {
		horizontalPosition, depth, aim = executeCommandWithAim(entry, horizontalPosition, depth, aim)
	}
	return horizontalPosition, depth
}

func executeCommandWithAim(entry PlanEntry, horizontalPosition, depth, aim int) (int, int, int) {
	switch entry.Direction {
	case "forward":
		horizontalPosition += entry.Velocity
		depth += aim * entry.Velocity
	case "up":
		aim -= entry.Velocity
	case "down":
		aim += entry.Velocity
	}

	return horizontalPosition, depth, aim
}
