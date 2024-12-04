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

	list1 := []int{}
	list2 := []int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		// list 1
		i, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatalf("Failed to parse string-int: %s", err)
		}
		list1 = append(list1, i)

		// list 2
		j, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("Failed to parse string-int: %s", err)
		}
		list2 = append(list2, j)
	}

	/*
		sort.Ints(list1)
		sort.Ints(list2)
		sum := 0
		for i := 0; i < len(list1); i++ {
			distance := list1[i] - list2[i]
			fmt.Printf("%v \n", distance)
			if distance < 0 {
				sum -= int(distance)
			} else {
				sum += int(distance)
			}
		}

		fmt.Printf("\n (day1) Sum: %v \n", sum)*/
	sum := 0
	for i := 0; i < len(list1); i++ {
		val := list1[i]
		count := 0
		for j := 0; j < len(list2); j++ {
			if list2[j] == val {
				count++
			}
		}
		sum += val * count
	}
	fmt.Printf("\n (day1) Sum: %v \n", sum)
}
