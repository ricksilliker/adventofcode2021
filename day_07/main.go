package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("./day_07/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var horizontalPositions []int
	var maxPosition int
	var minPosition int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inputStr := strings.Split(scanner.Text(), ",")
	for _, i := range inputStr {
		pos, _ := strconv.Atoi(i)
		horizontalPositions = append(horizontalPositions, pos)
		if pos > maxPosition {
			maxPosition = pos
		}
		if pos < minPosition {
			minPosition = pos
		}
	}
	fmt.Println(len(horizontalPositions))
	fmt.Printf("Min Position: %v\n", minPosition)
	fmt.Printf("Max Position: %v\n", maxPosition)

	costMap := make(map[int]int)

	var bestPosition int
	var bestFuelCost int
	var fuelCost int
	for i := 0; i < len(horizontalPositions); i++ {
		for j := 0; j < len(horizontalPositions); j++ {
			fuelCost += int(math.Abs(float64(horizontalPositions[j]) - float64(horizontalPositions[i])))
		}

		costMap[horizontalPositions[i]] = fuelCost

		if bestFuelCost == 0 || fuelCost < bestFuelCost {
			bestPosition = horizontalPositions[i]
			bestFuelCost = fuelCost
		}
		fuelCost = 0
	}

	fmt.Printf("Answer: %v\n", bestPosition)
	fmt.Println(costMap[bestPosition])
}

func partTwo() {
	f, err := os.Open("./day_07/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var horizontalPositions []int
	var maxPosition int
	var minPosition int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inputStr := strings.Split(scanner.Text(), ",")
	for _, i := range inputStr {
		pos, _ := strconv.Atoi(i)
		horizontalPositions = append(horizontalPositions, pos)
		if pos > maxPosition {
			maxPosition = pos
		}
		if pos < minPosition {
			minPosition = pos
		}
	}
	fmt.Println(len(horizontalPositions))
	fmt.Printf("Min Position: %v\n", minPosition)
	fmt.Printf("Max Position: %v\n", maxPosition)

	//horizontalPositions = []int{16,1,2,0,4,2,7,1,2,14}

	costMap := make(map[int]int)

	var bestPosition int
	var bestFuelCost int
	var fuelCost int
	for i := 0; i < maxPosition; i++ {
		for j := 0; j < len(horizontalPositions); j++ {
			steps := int(math.Abs(float64(horizontalPositions[j]) - float64(i)))
			for s := 0; s < steps; s++ {
				fuelCost += s + 1
			}
		}
		costMap[i] = fuelCost

		if bestFuelCost == 0 || fuelCost < bestFuelCost {
			bestPosition = i
			bestFuelCost = fuelCost
		}
		fuelCost = 0
	}

	fmt.Println(costMap)
	fmt.Printf("Answer: %v\n", bestPosition)
	fmt.Println(bestFuelCost)
}
