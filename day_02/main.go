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
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("./day_02/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var step []string
	var position int
	var depth int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rawValue := scanner.Text()
		step = strings.Fields(rawValue)

		amount, err := strconv.Atoi(step[1])
		if err != nil {
			log.Fatalf("error reading amount: %s", err)
			return
		}

		switch step[0] {
		case "forward":
			position += amount
		case "up":
			depth -= amount
			// Maybe need this.
			if depth < 0 {
				depth = 0
			}
		case "down":
			depth += amount
		default:
			fmt.Printf("Unknown direction: %s", step[0])
		}
	}

	fmt.Printf("Answer: %v \n", position * depth)
}

func partTwo() {
	f, err := os.Open("./day_02/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var step []string
	var position int
	var depth int
	var aim int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rawValue := scanner.Text()
		step = strings.Fields(rawValue)

		amount, err := strconv.Atoi(step[1])
		if err != nil {
			log.Fatalf("error reading amount: %s", err)
			return
		}

		switch step[0] {
		case "forward":
			position += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		default:
			fmt.Printf("Unknown direction: %s", step[0])
		}
	}

	fmt.Printf("Answer: %v \n", position * depth)
}