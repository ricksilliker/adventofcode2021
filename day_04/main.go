package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoCell struct {
	Number int
	Marked bool
}

type BingoBoard struct {
	Cells [5][5]BingoCell
}

func (b *BingoBoard) HasWon() bool {
	var rowMarkers int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Cells[i][j].Marked {
				rowMarkers++
			}
		}
		if rowMarkers == 5 {
			return true
		}
		rowMarkers = 0
	}

	return false
}

func (b *BingoBoard) GetMarkedRowOrColumn() bool {
	if b.HasWon() {
		return true
	} else {
		transposedBoard := BingoBoard{}
		for i, rows := range b.Cells {
			for j, cell := range rows {
				transposedBoard.Cells[j][i] = cell
			}
		}
		return transposedBoard.HasWon()
	}
}

func (b *BingoBoard) Equal(other BingoBoard) bool {
	var one string
	var two string
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			one += strconv.Itoa(b.Cells[i][j].Number)
			two += strconv.Itoa(other.Cells[i][j].Number)
		}
	}

	return one == two
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("./day_04/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var bingoCalls []int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	for _, val := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(val)
		bingoCalls = append(bingoCalls, num)
	}

	var boards []BingoBoard
	currentBoard := -1
	currentRow := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, BingoBoard{})
			currentRow = 0
			currentBoard++
			continue
		}

		rowData := scanner.Text()
		values := strings.Fields(rowData)
		//fmt.Println("column values", values)

		var newRow [5]BingoCell
		for col, value := range values {
			//fmt.Printf("on column %v\n", col)
			num, _ := strconv.Atoi(value)
			newRow[col] = BingoCell{
				Number: num,
				Marked: false,
			}
		}
		//fmt.Println(currentRow, newRow)

		boards[currentBoard].Cells[currentRow] = newRow
		currentRow++
	}

	fmt.Println("num boards", len(boards))

	var winner BingoBoard
	var winningNumber int
	out:
	for _, num := range bingoCalls {
		for b := 0; b < len(boards); b++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if boards[b].Cells[i][j].Number == num {
						boards[b].Cells[i][j].Marked = true
					}
				}
			}
		}
		for b := 0; b < len(boards); b++ {
			if boards[b].GetMarkedRowOrColumn() {
				winner = boards[b]
				winningNumber = num
				break out
			}
		}
	}
	//fmt.Println(winner)

	unmarkedSum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !winner.Cells[i][j].Marked {
				unmarkedSum += winner.Cells[i][j].Number
			}
		}
	}

	fmt.Println("Calculation: ", unmarkedSum, winningNumber)

	fmt.Println("Answer: ", unmarkedSum * winningNumber)
}

func partTwo() {
	f, err := os.Open("./day_04/input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err)
		return
	}
	defer f.Close()

	var bingoCalls []int
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	for _, val := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(val)
		bingoCalls = append(bingoCalls, num)
	}

	var boards []BingoBoard
	currentBoard := -1
	currentRow := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			boards = append(boards, BingoBoard{})
			currentRow = 0
			currentBoard++
			continue
		}

		rowData := scanner.Text()
		values := strings.Fields(rowData)
		//fmt.Println("column values", values)

		var newRow [5]BingoCell
		for col, value := range values {
			//fmt.Printf("on column %v\n", col)
			num, _ := strconv.Atoi(value)
			newRow[col] = BingoCell{
				Number: num,
				Marked: false,
			}
		}

		boards[currentBoard].Cells[currentRow] = newRow
		currentRow++
	}

	fmt.Println("num boards", len(boards))

	var winningBoards []BingoBoard
	var winningNumbers []int
	for _, num := range bingoCalls {
		for b := 0; b < len(boards); b++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if boards[b].Cells[i][j].Number == num {
						boards[b].Cells[i][j].Marked = true
					}
				}
			}
			if boards[b].GetMarkedRowOrColumn() {
				if !contains(winningBoards, boards[b]) {
					winningBoards = append(winningBoards, boards[b])
					winningNumbers = append(winningNumbers, num)
				}
			}
		}
		//for b := 0; b < len(boards); b++ {
		//
		//}
	}

	fmt.Println("Winning board count: ", len(winningBoards))

	unmarkedSum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !winningBoards[len(winningBoards)-1].Cells[i][j].Marked {
				unmarkedSum += winningBoards[len(winningBoards)-1].Cells[i][j].Number
			}
		}
	}

	fmt.Println("Calculation: ", unmarkedSum, winningNumbers[len(winningNumbers)-1])
	fmt.Println("Answer: ", unmarkedSum * winningNumbers[len(winningNumbers)-1])
}

func contains(list []BingoBoard, elem BingoBoard) bool {
	for i := 0; i < len(list); i++ {
		if list[i].Equal(elem) {
			return true
		}
	}
	return false
}