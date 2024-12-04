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

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	//56 57 54 51 51 49 45
	sum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		fmt.Println(scanner.Text())

		allLists := getAllPossible(line)

		for _, r := range allLists {
			if checkValid(r) {
				sum++
				break
			}
		}
	}
	fmt.Printf("day2 sum: %v", sum)
}

func getAllPossible(line []string) [][]string {
	possible := [][]string{}
	for i := 0; i < len(line); i++ {
		temp := []string{}
		for j := 0; j < len(line); j++ {
			if !(i == j) {
				temp = append(temp, line[j])
			}
		}
		possible = append(possible, temp)
	}
	return possible
}

func checkValid(report []string) bool {
	first, _ := strconv.Atoi(report[0])
	second, _ := strconv.Atoi(report[1])

	// positive if decending
	direction := first - second

	if direction == 0 {
		return false
	}

	for i := 1; i < len(report); i++ {
		x, err := strconv.Atoi(report[i-1])
		if err != nil {
			log.Fatalf("Failed to parse string-int: %s", err)
		}
		y, err := strconv.Atoi(report[i])
		if err != nil {
			log.Fatalf("Failed to parse string-int: %s", err)
		}
		// if decending
		if direction > 0 {
			if !(x > y) {
				return false
			}
		}

		// if accending
		if direction < 0 {
			if !(x < y) {
				return false
			}
		}

		if intAbs(x-y) > 3 || x-y == 0 {
			return false
		}
	}
	return true
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		log.Fatalf("Failed to read file: %s", err)
// 	}

// 	scanner := bufio.NewScanner(file)
// 	//56 57 54 51 51 49 45
// 	sum := 0

// 	for scanner.Scan() {
// 		line := strings.Split(scanner.Text(), " ")
// 		fmt.Println(scanner.Text())

// 		valid := true
// 		skipError := true

// 		first, _ := strconv.Atoi(line[0])
// 		second, _ := strconv.Atoi(line[1])

// 		// positive if decending
// 		direction := first - second

// 		if direction == 0 {
// 			valid = false
// 		}
// 		// 634 IS CORRECT

// 		for i := 1; i < len(line); i++ {
// 			x := -1
// 			y := -1
// 			if valid == false {
// 				break
// 			}
// 			if line[i] == "X" {
// 				if i+1 >= len(line) {
// 					break
// 				}
// 				x, err = strconv.Atoi(line[i-1])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}
// 				y, err = strconv.Atoi(line[i+1])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}
// 			} else if line[i-1] == "X" {
// 				x, err = strconv.Atoi(line[i-2])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}
// 				y, err = strconv.Atoi(line[i])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}

// 			} else {
// 				x, err = strconv.Atoi(line[i-1])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}
// 				y, err = strconv.Atoi(line[i])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}
// 			}

// 			// if decending
// 			if direction > 0 {
// 				if !(x > y) {
// 					if !skipError {
// 						valid = false
// 						break
// 					} else {
// 						skipError = false
// 						line[i] = "X"
// 						i = 0
// 						continue
// 					}
// 				}
// 			}

// 			// if accending
// 			if direction < 0 {
// 				if !(x < y) {
// 					if !skipError {
// 						valid = false
// 						break
// 					} else {
// 						skipError = false
// 						line[i] = "X"
// 						i = 0
// 						continue
// 					}
// 				}
// 			}

// 			if intAbs(x-y) > 3 || x-y == 0 {
// 				if !skipError {
// 					valid = false
// 					break
// 				} else {
// 					skipError = false
// 					line[i] = "X"
// 					i = 0
// 					continue
// 				}
// 			}

// 		}
// 		if valid {
// 			fmt.Printf(" yeah \n")
// 			sum++
// 			continue
// 		} else {
// 			valid = true
// 			line := strings.Split(scanner.Text(), " ")
// 			first, err = strconv.Atoi(line[1])
// 			if err != nil {
// 				log.Fatalf("Failed to parse string-int: %s", err)
// 			}
// 			second, err = strconv.Atoi(line[2])
// 			if err != nil {
// 				log.Fatalf("Failed to parse string-int: %s", err)
// 			}
// 			direction = first - second

// 			if direction == 0 {
// 				valid = false
// 				continue
// 			}

// 			// Check all but first in line
// 			for i := 2; i < len(line); i++ {
// 				x, err := strconv.Atoi(line[i-1])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}
// 				y, err := strconv.Atoi(line[i])
// 				if err != nil {
// 					log.Fatalf("Failed to parse string-int: %s", err)
// 				}

// 				// if decending
// 				if direction > 0 {
// 					if !(x > y) {
// 						valid = false
// 						break
// 					}
// 				}

// 				// if accending
// 				if direction < 0 {
// 					if !(x < y) {
// 						valid = false
// 						break
// 					}
// 				}

// 				if intAbs(x-y) > 3 || x-y == 0 {
// 					valid = false
// 					break
// 				}
// 			}
// 			if valid {
// 				fmt.Printf(" yeah \n")
// 				sum++
// 				continue
// 			}
// 		}
// 	}
// 	fmt.Printf("(day2) Sum: %v \n", sum)
// }

// func intAbs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x

// }
