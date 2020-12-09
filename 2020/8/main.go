package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var accumulator = 0

func main() {
	bootCode, err := loadBootCode("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	// Part 1
	runBoot(bootCode, false)
	fmt.Println("Part 1", accumulator)

	// Part 2
	accumulator = 0
	runBoot(bootCode, true)
	fmt.Println("Part 2", accumulator)

}

func runBoot(program []*string, attemptPatch bool) bool {
	executedInstructions := make(map[int]*string)
	index := 0
	var executionOrder []int

	patchAttempted := false
	for {
		if _, ok := executedInstructions[index]; ok {
			if !patchAttempted && attemptPatch {
				// Attempt patch
				for _, k := range executionOrder {
					accumulator = 0
					// Swap instruction
					program[k] = swapInstruction(executedInstructions[k])
					success := runBoot(program, false)
					// Swap back
					program[k] = swapInstruction(program[k])
					if success {
						return true
					}
				}
				patchAttempted = true
				return false
			}
			break
		}
		if index >= len(program) {
			return true
		}
		executedInstructions[index] = program[index]
		executionOrder = append(executionOrder, index)
		index = index + executeInstruction(*program[index])
	}
	return false
}

func swapInstruction(instruction *string) *string {
	parseInstruction := strings.Split(*instruction, " ")
	operation := parseInstruction[0]
	argument := parseInstruction[1]

	switch operation {
	case "nop":
		swapped := fmt.Sprintf("jmp %s", argument)
		return &swapped
	case "jmp":
		swapped := fmt.Sprintf("nop %s", argument)
		return &swapped
	}

	return instruction
}

func executeInstruction(instruction string) int {
	parseInstruction := strings.Split(instruction, " ")
	operation := parseInstruction[0]
	argument, _ := strconv.Atoi(parseInstruction[1])
	switch operation {
	case "nop":
		return 1
	case "acc":
		accumulator += argument
		return 1
	case "jmp":
		return argument
	}
	return 0
}

func loadBootCode(path string) ([]*string, error) {
	var bootCodeInstructions []*string
	bootCode, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load boot code: %s", err.Error())
	}
	defer func() {
		_ = bootCode.Close()
	}()

	bootCodeScanner := bufio.NewScanner(bootCode)
	bootCodeScanner.Split(bufio.ScanLines)
	for bootCodeScanner.Scan() {
		entry := bootCodeScanner.Text()
		bootCodeInstructions = append(bootCodeInstructions, &entry)
	}
	return bootCodeInstructions, nil
}
