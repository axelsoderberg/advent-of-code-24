package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

type Update struct {
	pages []int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	rules := []Rule{}
	updates := []Update{}

	parsingRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = false
			continue
		}
		if parsingRules {
			rules = append(rules, parseRule(line))
		} else {
			updates = append(updates, parseUpdate(line))
		}
	}

	// fmt.Printf("\n Rule: %v \n", rules[0])
	// fmt.Printf("\n Update: %v \n", updates[0])

	// Part 1
	sum := 0
	for _, u := range updates {
		if correctUpdate(u, rules) {
			sum += middlePageNum(u)
		}
	}

	fmt.Printf("\n (day5) (part1) Sum: %v \n", sum)

	// Part 2
	// Could have reused the ones not correct from part 1, but to make it clear I recalculate those
	sum = 0
	for _, u := range updates {
		if !correctUpdate(u, rules) {
			sum += middlePageNum(correctUpdateWithRules(u, rules))
		}
	}

	fmt.Printf("\n (day5) (part2) Sum: %v \n", sum)
}

func correctUpdateWithRules(u Update, rules []Rule) Update {
	rRules := relevantRules(u, rules)
	for _, r := range rRules {
		// Validates, applies rule if broken
		u = applyRule(u, r)
	}
	if correctUpdate(u, rRules) {
		return u
	} else {
		fmt.Printf("\n had to try again \n")
		correctUpdateWithRules(u, rRules)
	}
	return u
}

func applyRule(u Update, r Rule) Update {
	a := firstInstanceIndex(u.pages, r.after)
	b := firstInstanceIndex(u.pages, r.before)
	if a < b {
		u.pages = swap(u.pages, a, b)
	}
	return u
}

func swap(pages []int, i1, i2 int) []int {
	temp := pages[i1]
	pages[i1] = pages[i2]
	pages[i2] = temp
	return pages
}

func middlePageNum(u Update) int {
	return u.pages[len(u.pages)/2]
}

func correctUpdate(u Update, rules []Rule) bool {
	rRules := relevantRules(u, rules)
	for _, r := range rRules {
		if !validUpdateWithRule(u, r) {
			return false
		}
	}
	return true
}

func validUpdateWithRule(u Update, r Rule) bool {
	if firstInstanceIndex(u.pages, r.after) < firstInstanceIndex(u.pages, r.before) {
		return false
	}
	return true

}

func firstInstanceIndex(pages []int, val int) int {
	for i, p := range pages {
		if p == val {
			return i
		}
	}
	log.Fatalf("Warning: value %d not found in pages %v \n", val, pages)
	return -1

}

func relevantRules(u Update, rules []Rule) []Rule {
	rRules := []Rule{}
	for _, r := range rules {
		if contains(u.pages, r.before) && contains(u.pages, r.after) {
			rRules = append(rRules, r)
		}
	}
	return rRules
}

func contains(pages []int, i int) bool {
	for _, v := range pages {
		if v == i {
			return true
		}
	}
	return false
}

func parseUpdate(line string) Update {
	u := strings.Split(line, ",")
	pages := []int{}
	for _, u := range u {
		p, _ := strconv.Atoi(u)
		pages = append(pages, p)
	}
	return Update{
		pages: pages,
	}
}

func parseRule(line string) Rule {
	r := strings.Split(line, "|")
	b, _ := strconv.Atoi(r[0])
	a, _ := strconv.Atoi(r[1])
	return Rule{
		before: b,
		after:  a,
	}
}
