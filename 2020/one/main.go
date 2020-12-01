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

	expenses, err := getExpesnes(targetSum, expenseReport)
	if err != nil {
		fmt.Println("Unable to cmplete")
		return
	}

	result := 1
	for _, expense := range expenses{
		result *= expense
	}


	fmt.Printf("Expense research result: %d\n", result)
}

func loadExpenseReport(path string) ([]int, error){
	entries := []int{}
	reportFile, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("Unable to load expense report: %s", err.Error())
	}

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


func getExpesnes(targetSum int, expenses []int) ([]int, error){

	seen := []int{}
	results := []int{}
	for _, expense := range expenses {
		seen = append(seen, expense)
		searchVal := targetSum - expense
		match, found := find(searchVal, seen)
		if found {
			results = append(results, expense, match)
			return results, nil
		}
	}
	return results, fmt.Errorf("Pair not found")
}

func find(needle int, haystack []int)(int, bool){
	for _, item := range haystack {
		if item == needle {
			return item, true
		}
	}
	return -1, false
}