package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	targetSum := 2020
	expenseReport, err := loadExpenseReport("puzzle_input.txt")
	if err != nil {
		log.Fatalf("Unable to continue: %s", err.Error())
	}

	// Part 1
	fmt.Println("Part 1")
	expenses1, err := getExpenses(targetSum, expenseReport, 2)
	if err != nil {
		fmt.Println("Not found: ", err.Error())
	}
	printResults(&expenses1)

	// Part 2
	fmt.Println("Part 2")
	expenses2, err := getExpenses(targetSum, expenseReport, 3)
	if err != nil {
		fmt.Println("Not found: ", err.Error())
	}
	printResults(&expenses2)
}

func printResults(expenses *map[int]bool) {
	if len(*expenses) > 0 {
		result := 1
		for expense := range *expenses {
			result *= expense
		}
		fmt.Printf("Expense research result: %d\n", result)
	}
}


func loadExpenseReport(path string) ([]int, error){
	entries := []int{}
	reportFile, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("Unable to load expense report: %s", err.Error())
	}
	defer func() {
		_ = reportFile.Close()
	}()

	reportScanner := bufio.NewScanner(reportFile)
	reportScanner.Split(bufio.ScanWords)
	for reportScanner.Scan() {
		entry, err := strconv.Atoi(reportScanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Invalid data in expense report: %s (%s)", entry, err.Error())
		}
		entries = append(entries, entry)
	}
	return entries, reportScanner.Err()
}

func getExpenses(targetSum int, expenses []int, entries uint) (map[int]bool, error){
	seen := make(map[int]bool)
	results := make(map[int]bool)

	for _, expense := range expenses {
		seen[expense] = true
		searchVal := targetSum - expense

		if  entries > 2{
			subEntries := entries - 1
			subResults, err := getExpenses(searchVal, expenses, subEntries)
			if err != nil {
				continue
			}
			results[expense] = true
			for result := range subResults {
				results[result] = true
			}
			if len(results) == int(entries) {
				return results, nil
			}
		} else{
			match, found := find(searchVal, seen)
			if found {
				results[expense] = true
				results[match] = true
				return results, nil
			}
		}

	}
	return results, fmt.Errorf("Criteria not met.  Could not find %d expenses that add up to %d",
		entries,
		targetSum)
}

func find(needle int, haystack map[int]bool)(int, bool){
	for item := range haystack {
		if item == needle {
			return item, true
		}
	}
	return -1, false
}