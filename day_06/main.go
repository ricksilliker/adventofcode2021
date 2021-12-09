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
	//partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("./day_06/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	//lanterfish := []int{3, 3, 3}
	var lanterfish []int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inputStr := strings.Split(scanner.Text(), ",")
	for _, i := range inputStr {
		age, _ := strconv.Atoi(i)
		lanterfish = append(lanterfish, age)
	}

	fmt.Printf("Initial State %v\n", len(lanterfish))

	var days int
	var newFishToAdd int
	for {
		days++

		if days == 81 {
			break
		}

		for i := 0; i < len(lanterfish); i++ {
			if lanterfish[i] == 0 {
				lanterfish[i] = 6
				newFishToAdd++
			} else {
				lanterfish[i] -= 1
			}
		}

		for i := 0; i < newFishToAdd; i++ {
			lanterfish = append(lanterfish, 8)
		}
		//fmt.Printf("Number added: %v\n", newFishToAdd)
		//fmt.Printf("Total Day %v: %v\n", days, len(lanterfish))

		newFishToAdd = 0
	}

	fmt.Printf("Answer: %v\n", len(lanterfish))
}

func partTwo() {
	f, err := os.Open("./day_06/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var lanterfish [9]int64
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	inputStr := strings.Split(scanner.Text(), ",")
	for _, i := range inputStr {
		age, _ := strconv.Atoi(i)
		lanterfish[age] += 1
	}

	fmt.Printf("Initial State %v\n", lanterfish)

	var days int
	for {
		days++

		if days == 257 {
			break
		}

		var copiedSlice [9]int64
		for i := len(lanterfish) - 1; i >= 0; i-- {
			if i == 0 {
				copiedSlice[6] += lanterfish[0]
				copiedSlice[8] = lanterfish[0]
			} else {
				copiedSlice[i-1] = lanterfish[i]
			}
		}
		lanterfish = copiedSlice

		//fmt.Printf("Number added: %v\n", newFishToAdd)
		fmt.Printf("Total Day %v: %v\n", days, lanterfish)
	}

	var total int64
	for _, item := range lanterfish {
		total += item
	}
	fmt.Printf("Answer: %v\n", total)
}