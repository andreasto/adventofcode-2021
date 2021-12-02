package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type course struct {
	direction string
	value     int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	courses := []course{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, " ")
		parsed, err := strconv.Atoi(slice[1])
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course{slice[0], parsed})
	}

	result := part1(courses)
	fmt.Printf("Result for part1: %d\n", result)
	fmt.Printf("Result for part2: %d\n", part2(courses))
}

func part2(courses []course) int {
	x := 0
	y := 0
	aim := 0

	for _, c := range courses {
		if c.direction == "forward" {
			x += c.value
			y += aim * c.value
		} else if c.direction == "up" {
			aim -= c.value
		} else if c.direction == "down" {
			aim += c.value
		}
	}

	return x * y
}

func part1(courses []course) int {
	x := 0
	y := 0
	for _, c := range courses {
		if c.direction == "forward" {
			x += c.value
		} else if c.direction == "up" {
			y -= c.value
		} else if c.direction == "down" {
			y += c.value
		}
	}

	return x * y
}
