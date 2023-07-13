package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var currentLocation directionVector

var numberOfInstructions int

var visitedXY = make(map[string]bool) // map of visited locations

var countVisitedSpaces int // start at 1 because we start at the first location

func main() {
	start := time.Now()
	// read file
	scanner, file := openFile()
	defer file.Close()
	fetchNumberOfInstructions(scanner)
	readStartingLocation(scanner)
	processInstruction(scanner)
	// print result
	fmt.Printf("Cleaned:  %d\n\n", countVisitedSpaces)

	elapsed := time.Since(start)
	fmt.Printf("Processing took %s\n", elapsed)
}

func readStartingLocation(scanner *bufio.Scanner) {
	countVisitedSpaces = 1
	scanner.Scan()
	secondLine := scanner.Text()
	locationStrings := strings.Split(secondLine, " ") // will be in the form "x y"
	currentLocation.X, _ = strconv.Atoi(locationStrings[0])
	currentLocation.Y, _ = strconv.Atoi(locationStrings[1])
}

func processInstruction(scanner *bufio.Scanner) {
	for i := 0; i < numberOfInstructions; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		processDir(dir, steps)
	}

}

func processDir(dir string, steps int) {
	directionVector := makeDirectionsVector(dir)

	for i := 0; i < steps; i++ {
		currentLocation.add(directionVector)
		stringLocation := getXYString()

		// check if we've been here before
		if !visitedXY[stringLocation] {
			countVisitedSpaces++
			visitedXY[stringLocation] = true
		}
	}
}

func getXYString() string {
	return strconv.Itoa(currentLocation.X) + " " + strconv.Itoa(currentLocation.Y)
}

func openFile() (*bufio.Scanner, *os.File) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	return scanner, file
}

func fetchNumberOfInstructions(scanner *bufio.Scanner) {
	scanner.Scan()
	numberOfInstructions, _ = strconv.Atoi(scanner.Text())
}
