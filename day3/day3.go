package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Op struct {
	operation string
	x         int
	y         int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	input := ""

	for scanner.Scan() {
		input += scanner.Text()
	}
	//fmt.Printf("\n Input: \n %v", input)

	// Regex for extracting the mul(x,y)
	re, err := regexp.Compile(`mul\(\d{0,3},\d{0,3}\)|do\(\)|don['’]t\(\)`)
	if err != nil {
		log.Fatal("Error compiling regex:", err)
	}

	matches := re.FindAllString(input, -1)
	//fmt.Printf("\n Matches: \n %v", matches)

	// Regex for extracting the op, x, and y
	re2, err := regexp.Compile(`(\w+)\((\d{0,3}),(\d{0,3})\)`)
	if err != nil {
		log.Fatal("Error compiling regex:", err)
	}

	// Parse all the Ops
	calculations := []Op{}
	active := true
	for _, m := range matches {
		if m == "do()" {
			active = true
			continue
		} else if m == "don't()" {
			active = false
			continue
		}
		if active {
			fmt.Printf("\n Parse m: \n %v", m)
			d := re2.FindStringSubmatch(m)
			op := d[1]
			x, _ := strconv.Atoi(d[2])
			y, _ := strconv.Atoi(d[3])
			calculations = append(calculations, Op{
				op,
				x,
				y,
			})
		}
	}

	sum := 0
	for _, c := range calculations {
		if c.operation == "mul" {
			sum += c.x * c.y
		} else {
			log.Fatalf("\n %v is not a valid operation \n", c.operation)
		}
	}

	fmt.Printf("\n Total: %v", sum)
}
