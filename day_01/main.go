package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//partOne()
	partTwo()
}

func partOne () {
	f, err := os.Open("./day_01/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var depths []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rawValue := scanner.Text()
		val, err := strconv.Atoi(rawValue)
		if err != nil {
			log.Fatalf("error reading line value: %s", err)
			return
		}
		depths = append(depths, val)
	}

	var answer int
	for index, measurement := range depths {
		if index == 0 {
			continue
		}

		if measurement > depths[index-1] {
			answer += 1
		}
	}

	fmt.Printf("Answer: %v \n", answer)
}

func partTwo() {
	f, err := os.Open("./day_01/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var depths []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rawValue := scanner.Text()
		val, err := strconv.Atoi(rawValue)
		if err != nil {
			log.Fatalf("error reading line value: %s", err)
			return
		}
		depths = append(depths, val)
	}

	var pairs [][]int
	var j int
	for i := 0; i < len(depths); i++ {
		j = i + 3
		if j > len(depths) {
			continue
		}
		pairs = append(pairs, depths[i:j])
	}

	var answer int
	for index, pair := range pairs {
		if index + 1 == len(pairs) {
			break
		}

		sumA := 0
		for _, v := range pair {
			sumA += v
		}

		sumB := 0
		for _, v := range pairs[index+1] {
			sumB += v
		}

		if sumB > sumA {
			answer += 1
		}
	}

	fmt.Printf("Answer: %v \n", answer)
}