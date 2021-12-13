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
	f, err := os.Open("./day_08/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	digit1 := 2
	digit4 := 4
	digit7 := 3
	digit8 := 7
	answer := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		entry := scanner.Text()
		entryParts := strings.Split(entry, "|")
		outputSegments := strings.Split(strings.TrimSpace(entryParts[1]), " ")
		fmt.Println(outputSegments)
		for _, seg := range outputSegments {
			if len(seg) == digit1 || len(seg) == digit4 || len(seg) == digit7 || len(seg) == digit8 {
				answer++
			}
		}
	}

	fmt.Printf("Answer: %v\n", answer)
}

type Pattern struct {
	Zero  Signal
	One   Signal
	Two   Signal
	Three Signal
	Four  Signal
	Five  Signal
	Six   Signal
	Seven Signal
	Eight Signal
	Nine  Signal
}

func (p *Pattern) IsZero(s Signal) bool {
	missing := p.One.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Seven.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Four.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Eight.GetMissing(s)
	if missing != 1 {
		return false
	}

	return true
}

func (p *Pattern) IsOne(s Signal) bool {
	return s.Length() == 2
}

func (p *Pattern) IsTwo(s Signal) bool {
	missing := p.One.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Seven.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Four.GetMissing(s)
	if missing != 2 {
		return false
	}

	missing = p.Eight.GetMissing(s)
	if missing != 2 {
		return false
	}

	return true
}

func (p *Pattern) IsThree(s Signal) bool {
	missing := p.One.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Seven.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Four.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Eight.GetMissing(s)
	if missing != 2 {
		return false
	}

	return true
}

func (p *Pattern) IsFour(s Signal) bool {
	return s.Length() == 4
}

func (p *Pattern) IsFive(s Signal) bool {
	missing := p.One.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Seven.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Four.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Eight.GetMissing(s)
	if missing != 2 {
		return false
	}

	return true
}

func (p *Pattern) IsSix(s Signal) bool {
	missing := p.One.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Seven.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Four.GetMissing(s)
	if missing != 1 {
		return false
	}

	missing = p.Eight.GetMissing(s)
	if missing != 1 {
		return false
	}

	return true
}

func (p *Pattern) IsSeven(s Signal) bool {
	return s.Length() == 3
}

func (p *Pattern) IsEight(s Signal) bool {
	return s.Length() == 7
}

func (p *Pattern) IsNine(s Signal) bool {
	missing := p.One.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Seven.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Four.GetMissing(s)
	if missing != 0 {
		return false
	}

	missing = p.Eight.GetMissing(s)
	if missing != 1 {
		return false
	}

	return true
}

type Signal struct {
	A bool
	B bool
	C bool
	D bool
	E bool
	F bool
	G bool
}

func (s *Signal) Encoding() []bool {
	return []bool{s.A, s.B, s.C, s.D, s.E, s.F, s.G}
}

func (s *Signal) Length() int {
	var result int
	for _, b := range s.Encoding() {
		if b {
			result++
		}
	}
	return result
}

func (s *Signal) GetMissing(other Signal) int {
	var missing int

	for i, val := range s.Encoding() {
		if val && !other.Encoding()[i] {
			missing++
		}
	}

	return missing
}

func (s *Signal) fill(chars []string) {
	for _, c := range chars {
		switch c {
		case "a":
			s.A = true
		case "b":
			s.B = true
		case "c":
			s.C = true
		case "d":
			s.D = true
		case "e":
			s.E = true
		case "f":
			s.F = true
		case "g":
			s.G = true
		default:
			log.Fatalf("Unknown letter found filling signal.")
		}
	}
}

func partTwo() {
	f, err := os.Open("./day_08/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	digit1 := 2
	digit4 := 4
	digit7 := 3
	digit8 := 7

	var answer int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		entry := scanner.Text()
		entryParts := strings.Split(entry, "|")
		signalWires := strings.Split(strings.TrimSpace(entryParts[0]), " ")
		pattern := Pattern{}
		var unkownSignals []Signal
		for _, wire := range signalWires {
			var signal Signal
			signal.fill(strings.Split(wire, ""))
			switch len(wire) {
			case digit1:
				pattern.One = signal
			case digit4:
				pattern.Four = signal
			case digit7:
				pattern.Seven = signal
			case digit8:
				pattern.Eight = signal
			default:
				unkownSignals = append(unkownSignals, signal)
			}
		}

		for _, s := range unkownSignals {
			if pattern.IsZero(s) {
				pattern.Zero = s
			} else if pattern.IsTwo(s) {
				pattern.Two = s
			} else if pattern.IsThree(s) {
				pattern.Three = s
			} else if pattern.IsFive(s) {
				pattern.Five = s
			} else if pattern.IsSix(s) {
				pattern.Six = s
			} else if pattern.IsNine(s) {
				pattern.Nine = s
			} else {
				log.Fatalf("failed to find matching digit")
			}
		}

		outputSegments := strings.Split(strings.TrimSpace(entryParts[1]), " ")
		fmt.Println(len(signalWires))
		fmt.Println(outputSegments)

		var solvedValue string
		for _, seg := range outputSegments {
			var s Signal
			s.fill(strings.Split(seg, ""))
			if pattern.IsZero(s) {
				solvedValue += "0"
			} else if pattern.IsOne(s) {
				solvedValue += "1"
			} else if pattern.IsTwo(s) {
				solvedValue += "2"
			} else if pattern.IsThree(s) {
				solvedValue += "3"
			} else if pattern.IsFour(s) {
				solvedValue += "4"
			} else if pattern.IsFive(s) {
				solvedValue += "5"
			} else if pattern.IsSix(s) {
				solvedValue += "6"
			} else if pattern.IsSeven(s) {
				solvedValue += "7"
			} else if pattern.IsEight(s) {
				solvedValue += "8"
			} else if pattern.IsNine(s) {
				solvedValue += "9"
			} else {
				log.Fatalf("failed to find matching digit")
			}
		}

		solvedInt, _ := strconv.Atoi(solvedValue)
		answer += solvedInt
	}

	fmt.Printf("Answer: %v\n", answer)
}
