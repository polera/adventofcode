package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var bagRuleParser = regexp.MustCompile(`(\d{1})?\W?(\w+\W\w+) bags?`)

func main() {
	bagRules, err := loadBagRules("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	// Part 1
	bags := make(map[string]bool)
	fmt.Println("Part 1: How many bags contain at least one shiny gold bag? ",
		len(bagExistsCount("shiny gold", bagRules, bags)))

	// Part 2
	fmt.Println("Part 2: How many individual bags in one shiny gold bag? ",
		bagAccumulation("shiny gold", bagRules))
}

func bagExistsCount(query string, rules map[string]map[string]int, bags map[string]bool) map[string]bool {
	for rule := range rules {
		if _, ok := rules[rule][query]; ok {
			bags[rule] = true
			bags = bagExistsCount(rule, rules, bags)
		}
	}
	return bags
}

func bagAccumulation(query string, rules map[string]map[string]int) int {
	accumulation := 0
	for bag, count := range rules[query] {
		accumulation += count + count*bagAccumulation(bag, rules)
	}
	return accumulation
}

func loadBagRules(path string) (map[string]map[string]int, error) {
	bagRules := make(map[string]map[string]int)
	rules, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load bag rules: %s", err.Error())
	}
	defer func() {
		_ = rules.Close()
	}()

	rulesScanner := bufio.NewScanner(rules)
	rulesScanner.Split(bufio.ScanLines)
	for rulesScanner.Scan() {
		entry := rulesScanner.Text()
		ruleMatches := bagRuleParser.FindAllStringSubmatch(entry, -1)
		bagRule := ""
		for index, match := range ruleMatches {
			if index == 0 {
				bagRule = match[len(match)-1]
				bagRules[bagRule] = make(map[string]int)
				continue
			}
			bagCount, err := strconv.Atoi(match[1])
			if err != nil {
				bagCount = 0
			}
			bagRules[bagRule][match[2]] = bagCount
		}
	}
	return bagRules, nil
}
