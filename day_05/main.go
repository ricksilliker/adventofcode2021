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

type Point struct {
	X int
	Y int
}

type Line struct {
	Points [2]Point
}

func (l *Line) X1() int {
	return l.Points[0].X
}

func (l *Line) X2() int {
	return l.Points[1].X
}

func (l *Line) Y1() int {
	return l.Points[0].Y
}

func (l *Line) Y2() int {
	return l.Points[1].Y
}

var points = make(map[int]int)

func main() {
	f, err := os.Open("./day_05/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var numSegments int
	var segments []Line

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(strings.ReplaceAll(line, " ", ""), "->")
		coordinates := strings.Split(pairs[0], ",")
		x1, _ := strconv.Atoi(coordinates[0])
		y1, _ := strconv.Atoi(coordinates[1])

		coordinates = strings.Split(pairs[1], ",")
		x2, _ := strconv.Atoi(coordinates[0])
		y2, _ := strconv.Atoi(coordinates[1])

		if x1 == x2 || y1 == y2 {
			numSegments++
			segments = append(segments, Line{
				Points: [2]Point{
					{x1, y1},
					{x2, y2},
				},
			})

		}
	}

	fmt.Println(segments)

	for {
		currentSegment := segments[0]
		segments = append(segments[:0], segments[1:]...)

		if len(segments) == 0 {
			break
		}

		for _, item := range segments {
			solveIntersection(currentSegment, item)
		}

	}

	//fmt.Println(points)

	var answer int
	for _, elem := range points {
		if elem > 1 {
			answer += elem
		}
	}

	fmt.Printf("Answer: %v\n", answer)
}

func solveIntersection(line1, line2 Line) {
	if line1.X1() == line2.X1() {
		min := int(math.Min(float64(line1.Y2()), float64(line2.Y2())))
		max := int(math.Max(float64(line1.Y1()), float64(line2.Y1())))
		//min := int(math.Min(math.Min(float64(line1.Y1()), float64(line1.Y2())), math.Min(float64(line2.Y1()), float64(line2.Y2()))))
		//max := int(math.Max(math.Max(float64(line1.Y1()), float64(line1.Y2())), math.Max(float64(line2.Y1()), float64(line2.Y2()))))
		for i := min; i < max; i++ {
			if _, ok := points[i]; !ok {
				points[i] = 1
			} else {
				points[i] += 1
			}
		}
	} else if line1.Y1() == line2.Y2() {
		min := int(math.Min(float64(line1.X2()), float64(line2.X2())))
		max := int(math.Max(float64(line1.X1()), float64(line2.X1())))

		//min := int(math.Min(math.Min(float64(line1.X1()), float64(line1.X2())), math.Min(float64(line2.X1()), float64(line2.X2()))))
		//max := int(math.Max(math.Max(float64(line1.X1()), float64(line1.X2())), math.Max(float64(line2.X1()), float64(line2.X2()))))
		for i := min; i < max; i++ {
			if _, ok := points[i]; !ok {
				points[i] = 1
			} else {
				points[i] += 1
			}
		}
	} else {
		math.Abs(float64(line2.Y2() - line2.Y1()))
	}
}