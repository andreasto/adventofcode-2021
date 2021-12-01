package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depths := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, parsed)
	}

	increased := 0
	previous := -1
	for i, d := range depths {
		if i == 0 {
			previous = d
			continue
		}

		if d > previous {
			increased++
			previous = d
		}
	}

	fmt.Printf("The depth increased: %d times", increased)
}
