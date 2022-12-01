package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	instructions, err := loadInstructions("puzzle_input.txt")
	if err != nil {
		log.Fatalf("unable to continue: %s", err.Error())
	}

	// Part 1
	nav := &Navigation{
		Heading:     "E",
		Coordinates: []complex128{complex(0, 0)}, // set initial waypoint
	}
	fmt.Println("Part 1: ", navigate(nav, instructions))

	// Part 2
	nav2 := &Navigation{
		Heading:  "E",
		Waypoint: complex(10, 1),
	}
	fmt.Println("Part 2: ", navigateV2(nav2, instructions))
}

var instructionParser = regexp.MustCompile(`^(\w)(\d+)$`)

type Navigation struct {
	Heading     string
	Coordinates []complex128
	Waypoint    complex128
}

func (n *Navigation) moveV2(direction string, units int) {

	directions := map[string]complex128{
		"N": complex(0.0, 1.0),
		"S": complex(0.0, -1.0),
		"E": complex(1.0, 0.0),
		"W": complex(-1.0, 0.0),
	}

	switch direction {
	case "L":
		degrees := float64(units)
		n.rotateWaypoint(direction, degrees)
	case "R":
		degrees := float64(units)
		n.rotateWaypoint(direction, degrees)
	case "F":
		coordinate := complex(float64(units), 0) * n.Waypoint
		n.Coordinates = append(n.Coordinates, coordinate)
	default:
		n.Waypoint += directions[direction] * complex(float64(units), 0)
	}
}

func (n *Navigation) rotateWaypoint(direction string, degrees float64) {
	switch direction {
	case "L":
		// counterclockwise
		switch degrees {
		case 90.0:
			n.Waypoint = complex(imag(n.Waypoint)*-1, real(n.Waypoint)) // (-y, x)
		case 180.0:
			n.Waypoint = complex(real(n.Waypoint)*-1, imag(n.Waypoint)*-1) // (-x, -y)
		case 270.0:
			n.Waypoint = complex(imag(n.Waypoint), real(n.Waypoint)*-1) // (y, -x)
		}
	case "R":
		// clockwise
		switch degrees {
		case 90.0:
			n.Waypoint = complex(imag(n.Waypoint), real(n.Waypoint)*-1) // (y, -x)
		case 180.0:
			n.Waypoint = complex(real(n.Waypoint)*-1, imag(n.Waypoint)*-1) // (-x, -y)
		case 270.0:
			n.Waypoint = complex(imag(n.Waypoint)*-1, real(n.Waypoint)) // (-y, x)
		}
	}
}

func (n *Navigation) moveV1(direction string, units int) {

	directions := map[string]complex128{
		"N": complex(0.0, 1.0),
		"S": complex(0.0, -1.0),
		"E": complex(1.0, 0.0),
		"W": complex(-1.0, 0.0),
	}

	switch direction {
	case "L":
		degrees := float64(units) / 90.0
		n.steer(degrees, direction)
	case "R":
		degrees := float64(units) / 90.0
		n.steer(degrees, direction)
	case "F":
		coordinate := directions[n.Heading] * complex(float64(units), 0)
		n.Coordinates = append(n.Coordinates, coordinate)
	default:
		coordinate := directions[direction] * complex(float64(units), 0)
		n.Coordinates = append(n.Coordinates, coordinate)
	}
}

func (n *Navigation) steer(degrees float64, direction string) {
	switch direction {
	case "L":
		switch degrees {
		case 1.0:
			switch n.Heading {
			case "E":
				n.Heading = "N"
			case "S":
				n.Heading = "E"
			case "W":
				n.Heading = "S"
			case "N":
				n.Heading = "W"
			}
		case 2.0:
			switch n.Heading {
			case "E":
				n.Heading = "W"
			case "S":
				n.Heading = "N"
			case "W":
				n.Heading = "E"
			case "N":
				n.Heading = "S"
			}
		case 3.0:
			switch n.Heading {
			case "E":
				n.Heading = "S"
			case "S":
				n.Heading = "W"
			case "W":
				n.Heading = "N"
			case "N":
				n.Heading = "E"
			}
		}
	case "R":
		switch degrees {
		case 1.0:
			switch n.Heading {
			case "E":
				n.Heading = "S"
			case "S":
				n.Heading = "W"
			case "W":
				n.Heading = "N"
			case "N":
				n.Heading = "E"
			}
		case 2.0:
			switch n.Heading {
			case "E":
				n.Heading = "W"
			case "S":
				n.Heading = "N"
			case "W":
				n.Heading = "E"
			case "N":
				n.Heading = "S"
			}
		case 3.0:
			switch n.Heading {
			case "E":
				n.Heading = "N"
			case "S":
				n.Heading = "E"
			case "W":
				n.Heading = "S"
			case "N":
				n.Heading = "W"
			}
		}
	}
}

func (n *Navigation) CalculateManhattanDistance() float64 {
	sum := complex(float64(0), 0)
	for _, coordinate := range n.Coordinates {
		sum += coordinate
	}

	return math.Abs(real(sum)) + math.Abs(imag(sum))
}

func navigateV2(navigation *Navigation, instructions []string) float64 {

	for _, instruction := range instructions {
		parts := instructionParser.FindAllStringSubmatch(instruction, -1)
		direction := parts[0][1]
		units := parts[0][2]
		intUnits, err := strconv.Atoi(units)
		if err != nil {
			log.Fatalf("invalid units: %s", err.Error())
		}
		navigation.moveV2(direction, intUnits)
	}

	return navigation.CalculateManhattanDistance()
}

func navigate(navigation *Navigation, instructions []string) float64 {

	for _, instruction := range instructions {
		parts := instructionParser.FindAllStringSubmatch(instruction, -1)
		direction := parts[0][1]
		units := parts[0][2]
		intUnits, err := strconv.Atoi(units)
		if err != nil {
			log.Fatalf("invalid units: %s", err.Error())
		}
		navigation.moveV1(direction, intUnits)
	}

	return navigation.CalculateManhattanDistance()
}

func loadInstructions(path string) ([]string, error) {
	var instructionSet []string
	instructions, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, fmt.Errorf("unable to load instructions: %s", err.Error())
	}
	defer func() {
		_ = instructions.Close()
	}()

	instructionScanner := bufio.NewScanner(instructions)
	instructionScanner.Split(bufio.ScanLines)
	for instructionScanner.Scan() {
		entry := instructionScanner.Text()
		instructionSet = append(instructionSet, entry)
	}

	return instructionSet, nil
}
