package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	collectedForms, err := readFormsV1("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	yesCount := 0
	for _, form := range collectedForms {
		yesCount += len(form)
	}

	fmt.Println("Part 1 - yes counts: ", yesCount)

	formScores, err := readFormsV2("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	fmt.Println("Part 2 - yes counts: ", formScores)
}

func readFormsV2(path string) (int, error) {
	yesAnswers := make(map[string]int)
	formScores := 0
	forms, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return -1, fmt.Errorf("unable to load customs forms: %s", err.Error())
	}
	defer func() {
		_ = forms.Close()
	}()

	formScanner := bufio.NewScanner(forms)
	formScanner.Split(bufio.ScanLines)
	groupCount := 0
	for formScanner.Scan() {
		entry := formScanner.Text()
		if len(entry) > 0 {
			entrySplit := strings.Split(entry, "")
			for _, entry := range entrySplit {
				yesAnswers[entry] += 1
			}
			groupCount += 1
			continue
		}
		for key := range yesAnswers {
			if yesAnswers[key] == groupCount {
				formScores += 1
			}
		}
		yesAnswers = make(map[string]int)
		groupCount = 0
	}
	for key := range yesAnswers {
		if yesAnswers[key] == groupCount {
			formScores += 1
		}
	}
	return formScores, nil
}

func readFormsV1(path string) ([]map[string]int, error) {
	yesAnswers := make(map[string]int)
	collectedForms := make([]map[string]int, 1, 1)
	forms, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load customs forms: %s", err.Error())
	}
	defer func() {
		_ = forms.Close()
	}()

	formScanner := bufio.NewScanner(forms)
	formScanner.Split(bufio.ScanLines)
	for formScanner.Scan() {
		entry := formScanner.Text()
		if len(entry) > 0 {
			entrySplit := strings.Split(entry, "")
			for _, entry := range entrySplit {
				yesAnswers[entry] = 1
			}
			continue
		}
		collectedForms = append(collectedForms, yesAnswers)
		yesAnswers = make(map[string]int)
	}
	collectedForms = append(collectedForms, yesAnswers)
	return collectedForms, nil
}
