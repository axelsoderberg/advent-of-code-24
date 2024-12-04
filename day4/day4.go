package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}

var wordLenght = 4

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	input := [][]string{}

	row := 0
	for scanner.Scan() {
		rowInput := scanner.Text()

		rowSlice := make([]string, len(rowInput))
		for col := 0; col < len(rowInput); col++ {
			rowSlice[col] = string(rowInput[col])
		}
		input = append(input, rowSlice)
		row++
	}

	// loop through each position in the matrix
	sum := 0
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			// get tot starting from this
			sum += countWordsFromPos(input, r, c)
		}
	}

	fmt.Printf("\n (day4) (part1) Sum: %v \n", sum)

	// Part 2
	sum = 0
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			// get tot starting from this
			sum += countCrossWordFromPos(input, r, c)
		}
	}
	fmt.Printf("\n (day4) (part2) Sum: %v \n", sum)
}

func countWordsFromPos(input [][]string, r int, c int) int {
	count := 0
	//check left-right
	count += checkLeftRight(input, r, c)
	// check up-down
	count += checkUpDown(input, r, c)
	// check all 4 diagonals
	count += checkDiagonals(input, r, c)
	if count > 0 {
		fmt.Printf("\n Found %v words on %v,%v", count, r, c)
	}
	return count
}

func countCrossWordFromPos(input [][]string, r int, c int) int {
	count := 0
	//check left-right
	count += checkCross(input, r, c)
	if count > 0 {
		fmt.Printf("\n Found cross on %v,%v", r, c)
	}
	return count
}

func checkCross(input [][]string, r, c int) int {
	// Check from upper left
	crossWords := []string{}
	reading := ""
	for i := 0; i < 3; i++ {
		x := c + i - 1
		y := r + i - 1
		if !inBounds(y, x, len(input), len(input[0])) {
			break
		}
		reading += input[y][x]
	}
	crossWords = append(crossWords, reading)

	// Check from upper right
	reading = ""
	for i := 0; i < 3; i++ {
		x := c + i - 1
		y := r - i + 1
		if !inBounds(y, x, len(input), len(input[0])) {
			break
		}
		reading += input[y][x]
	}
	crossWords = append(crossWords, reading)
	if correctCross(crossWords) {
		return 1
	}
	return 0

}

func checkDiagonals(input [][]string, r, c int) int {
	count := 0

	directions := [][2]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	// Upper left
	for _, dir := range directions {
		reading := ""
		for i := 0; i < wordLenght; i++ {
			x := c + i*dir[0]
			y := r + i*dir[1]
			if !inBounds(y, x, len(input), len(input[0])) {
				break
			}
			reading += input[y][x]
		}
		if correctWord(reading) {
			count++
		}
	}
	return count
}

func checkLeftRight(input [][]string, r, c int) int {
	count := 0

	// left
	reading := ""
	for i := c - wordLenght + 1; i <= c; i++ {
		if !inBounds(r, i, len(input), len(input[0])) {
			break
		} else {
			// add letter to reading
			reading += input[r][i]
		}
	}
	if correctWord(reading) {
		count++
	}

	// Right
	reading = ""
	for i := c; i < wordLenght+c; i++ {
		if !inBounds(r, i, len(input), len(input[0])) {
			break
		} else {
			// add letter to reading
			reading += input[r][i]
		}
	}
	if correctWord(reading) {
		count++
	}
	return count
}

func checkUpDown(input [][]string, r, c int) int {
	count := 0

	// Up
	reading := ""
	for i := r - wordLenght + 1; i <= r; i++ {
		if !inBounds(i, c, len(input), len(input[0])) {
			break
		} else {
			// add letter to reading
			reading += input[i][c]
		}
	}
	if correctWord(reading) {
		count++
	}

	// Down
	reading = ""
	for i := r; i < wordLenght+r; i++ {
		if !inBounds(i, c, len(input), len(input[0])) {
			break
		} else {
			// add letter to reading
			reading += input[i][c]
		}
	}
	if correctWord(reading) {
		count++
	}
	return count
}

// TODO Should break out the correct word
func correctWord(reading string) bool {
	if reading == "XMAS" {
		return true
	}
	return false
}

func correctCross(words []string) bool {
	for i := 0; i < len(words); i++ {
		if words[i] != "MAS" && words[i] != "SAM" {
			return false
		}
	}
	return true
}

func inBounds(y, x, sizeY, sizeX int) bool {
	return x >= 0 && x < sizeX && y >= 0 && y < sizeY
}
