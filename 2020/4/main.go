package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	validPassportsV1, err := loadValidPassportRecords("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue %s", err.Error())
	}

	fmt.Println("Part 1 - valid passports: ", len(validPassportsV1))

	var validPassportsV2 []*PassportRecord
	for _, record := range validPassportsV1 {
		if record.isValidV2() {
			validPassportsV2 = append(validPassportsV2, record)
		}
	}
	fmt.Println("Part 2 - valid passports: ", len(validPassportsV2))
}

var (
	validYear       = regexp.MustCompile(`^\d{4}$`)
	validHeight     = regexp.MustCompile(`^(\d{2,3})(cm|in)$`)
	validHairColor  = regexp.MustCompile(`^#([0-9a-f]{6})$`)
	validEyeColor   = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	validPassportId = regexp.MustCompile(`^\d{9}$`)
)

type PassportRecord struct {
	Fields map[string]string
}

func (p *PassportRecord) isValidV1() bool {
	// "cid" is optional
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range requiredFields {
		if _, present := p.Fields[field]; !present {
			return false
		}
	}
	return true
}

func (p *PassportRecord) isValidV2() bool {
	// "cid" is optional
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		switch field {
		case "byr": // validate birth year
			if !validYear.MatchString(p.Fields[field]) {
				return false
			}
			byr, err := strconv.Atoi(p.Fields[field])
			if err != nil {
				return false
			}
			// Birth year between 1920 and 2002
			if !(byr >= 1920 && byr <= 2002) {
				return false
			}
		case "iyr": // validate issue year
			if !validYear.MatchString(p.Fields[field]) {
				return false
			}
			iyr, err := strconv.Atoi(p.Fields[field])
			if err != nil {
				return false
			}
			// Issue year between 2010 and 2020
			if !(iyr >= 2010 && iyr <= 2020) {
				return false
			}
		case "eyr": // validate expiration year
			if !validYear.MatchString(p.Fields[field]) {
				return false
			}
			eyr, err := strconv.Atoi(p.Fields[field])
			if err != nil {
				return false
			}
			// Expiration year between 2020 and 2030
			if !(eyr >= 2020 && eyr <= 2030) {
				return false
			}
		case "hgt": // validate height
			hgtVals := validHeight.FindStringSubmatch(p.Fields[field])
			if len(hgtVals) < 1 {
				return false
			}
			hgt, err := strconv.Atoi(hgtVals[1])
			if err != nil {
				return false
			}
			unit := hgtVals[2]
			switch unit {
			case "in":
				// Height between 59 and 76 in
				if !(hgt >= 59 && hgt <= 76) {
					return false
				}
			case "cm":
				// Height between 150 and 193 cm
				if !(hgt >= 150 && hgt <= 193) {
					return false
				}
			}
		case "hcl":
			if !validHairColor.MatchString(p.Fields[field]) {
				return false
			}
		case "ecl":
			if !validEyeColor.MatchString(p.Fields[field]) {
				return false
			}
		case "pid":
			if !validPassportId.MatchString(p.Fields[field]) {
				return false
			}
		}
	}
	return true

}

func (p *PassportRecord) parse(rawRecord string) {
	p.Fields = make(map[string]string)
	fields := strings.Split(rawRecord, " ")
	for _, field := range fields {
		splitField := strings.Split(field, ":")
		p.Fields[splitField[0]] = splitField[1]
	}
}

func loadValidPassportRecords(path string) ([]*PassportRecord, error) {
	var v1records []*PassportRecord
	passportBatch, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load passport batch: %s", err.Error())
	}
	defer func() {
		_ = passportBatch.Close()
	}()

	passportBatchScanner := bufio.NewScanner(passportBatch)
	passportBatchScanner.Split(bufio.ScanLines)
	entries := []string{}
	for passportBatchScanner.Scan() {
		entry := passportBatchScanner.Text()
		if len(entry) > 0 {
			entries = append(entries, entry)
			continue
		}
		record := &PassportRecord{}
		record.parse(strings.Join(entries, " "))
		if record.isValidV1() {
			v1records = append(v1records, record)
		}
		// reset entries
		entries = []string{}
	}
	// Process last record; this is terrible - make it not terrible later
	record := &PassportRecord{}
	record.parse(strings.Join(entries, " "))
	if record.isValidV1() {
		v1records = append(v1records, record)
	}

	return v1records, passportBatchScanner.Err()
}
