package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	validV1, validV2, err := loadValidPwRecords("puzzle_input.txt")
	if err != nil {
		log.Fatalf("Unable to continue: %s", err.Error())
	}

	fmt.Println("Valid V1 password records: ", len(validV1))
	fmt.Println("Valid V2 password records: ", len(validV2))
}

type PasswordRecord struct {
	Requirement Requirement
	Password string
}

type Requirement struct {
	Min int
	Max int
	Character string
}

func (p *PasswordRecord) parse(rawEntry string) {
	parts := strings.Split(rawEntry, " ")

	// get rule criteria
	minMax := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	// get significant character requirement
	char := parts[1][0:1] // strip trailing ":"

	// get password value
	password := parts[2]

	req := Requirement{
		Min:       min,
		Max:       max,
		Character: char,
	}

	p.Requirement = req
	p.Password = password
}

func (p *PasswordRecord) isValidV1() bool {

	occurrences := strings.Count(p.Password, p.Requirement.Character)
	// Must contain character
	if occurrences < 1{
		return false
	}

	if !(p.Requirement.Min <= occurrences && occurrences <= p.Requirement.Max) {
		return false
	}

	return true
}

func (p *PasswordRecord) isValidV2() bool {

	occurrences := strings.Count(p.Password, p.Requirement.Character)

	// Must contain character
	if occurrences < 1{
		return false
	}

	// Invalid - both positions cannot contain character
	if string(p.Password[p.Requirement.Min-1]) == p.Requirement.Character &&
		string(p.Password[p.Requirement.Max-1]) == p.Requirement.Character {
		return false
	}

	// If first position contains character, second may not
	if string(p.Password[p.Requirement.Min-1]) == p.Requirement.Character &&
		string(p.Password[p.Requirement.Max-1]) != p.Requirement.Character {
		return true
	}

	// If second position contains character, first may not
	if len(p.Password) >= p.Requirement.Max -1 &&
		string(p.Password[p.Requirement.Max-1]) == p.Requirement.Character  &&
		string(p.Password[p.Requirement.Min-1]) != p.Requirement.Character{
		return true
	}

	return false
}


func loadValidPwRecords(path string) ([]*PasswordRecord, []*PasswordRecord, error){
	var v1records []*PasswordRecord
	var v2records []*PasswordRecord
	pwDatabase, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to load pw database: %s", err.Error())
	}
	defer func() {
		_ = pwDatabase.Close()
	}()

	pwDatabaseScanner := bufio.NewScanner(pwDatabase)
	pwDatabaseScanner.Split(bufio.ScanLines)
	for pwDatabaseScanner.Scan() {
		entry := pwDatabaseScanner.Text()
		record := &PasswordRecord{}
		record.parse(entry)
		if record.isValidV1() {
			v1records = append(v1records, record)
		}
		if record.isValidV2() {
			v2records = append(v2records, record)
		}
	}
	return v1records, v2records, pwDatabaseScanner.Err()
}
