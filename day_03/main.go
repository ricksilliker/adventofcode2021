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
	f, err := os.Open("./day_03/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	bitZeroCount := make([]int, 12)
	totalReports := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		totalReports += 1
		for i, bit := range strings.Split(scanner.Text(), "") {
			if bit == "0" {
				bitZeroCount[i] += 1
			}
		}
	}

	gammaRate := make([]int, 12)
	epsilonRate := make([]int, 12)
	for i, value := range bitZeroCount {
		fmt.Printf("column: %v, count: %v, totalCount: %v\n", i, value, totalReports)
		if value > totalReports / 2 {
			gammaRate[i] = 0
			epsilonRate[i] = 1
		} else {
			gammaRate[i] = 1
			epsilonRate[i] = 0
		}
	}

	gammaRateStr := ""
	epsilonRateStr := ""
	for i, char := range gammaRate {
		gammaRateStr += strconv.Itoa(char)
		epsilonRateStr += strconv.Itoa(epsilonRate[i])
	}

	fmt.Printf("Gamma rate: %s\n", gammaRateStr)
	fmt.Printf("Epsilon rate: %s\n", epsilonRateStr)

	gammaRateDecimal, _ := strconv.ParseInt(gammaRateStr, 2, 64)
	fmt.Printf("Gamma Rate: %v\n", gammaRateDecimal)
	epsilonRateDecimal, _ := strconv.ParseInt(epsilonRateStr, 2, 64)
	fmt.Printf("Epsilon Rate: %v\n", epsilonRateDecimal)
	fmt.Printf("Answer: %v\n", gammaRateDecimal * epsilonRateDecimal)

}

func partTwo() {
	f, err := os.Open("./day_03/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	totalReports := 0

	var oxygenBits []string
	var carbonBits []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		totalReports += 1
		oxygenBits = append(oxygenBits, scanner.Text())
		carbonBits = append(carbonBits, scanner.Text())
	}

	// Get O2 rating.
	for i := 0; i < 12; i++ {
		if len(oxygenBits) == 1 {
			break
		}
		fmt.Println(i)

		var ones int
		var zeroes int
		for _, bit := range oxygenBits {
			bits := strings.Split(bit, "")
			if bits[i] == "0" {
				zeroes ++
			} else {
				ones ++
			}
		}

		if zeroes > ones  {
			var n int
			for _, b := range oxygenBits {
				strSlice := strings.Split(b, "")
				if strSlice[i] == "0" {
					oxygenBits[n] = b
					n++
				}
			}
			oxygenBits = oxygenBits[:n]
		} else {
			var n int
			for _, b := range oxygenBits {
				strSlice := strings.Split(b, "")
				if strSlice[i] == "1" {
					oxygenBits[n] = b
					n++
				}
			}
			oxygenBits = oxygenBits[:n]
		}
	}

	fmt.Println(oxygenBits)

	oxygenRating, _ := strconv.ParseInt(oxygenBits[0], 2, 64)
	fmt.Printf("02 Rating: %v\n", oxygenRating)

	//	Get CO2 Rating.
	for i := 0; i < 12; i++ {
		if len(carbonBits) == 1 {
			break
		}
		var ones int
		var zeroes int
		for _, bit := range carbonBits {
			bits := strings.Split(bit, "")
			if bits[i] == "0" {
				zeroes ++
			} else {
				ones ++
			}
		}

		if ones < zeroes  {
			var n int
			for _, b := range carbonBits {
				strSlice := strings.Split(b, "")
				if strSlice[i] == "1" {
					carbonBits[n] = b
					n++
				}
			}
			carbonBits = carbonBits[:n]
		} else {
			var n int
			for _, b := range carbonBits {
				strSlice := strings.Split(b, "")
				if strSlice[i] == "0" {
					carbonBits[n] = b
					n++
				}
			}
			carbonBits = carbonBits[:n]
		}
	}

	fmt.Println(carbonBits)

	carbonRating, _ := strconv.ParseInt(carbonBits[0], 2, 64)
	fmt.Printf("02 Rating: %v\n", carbonRating)

	fmt.Printf("Answer: %v\n", oxygenRating * carbonRating)
}
