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

	chunks := chunkSlice(depths, 3)
	previousWindow := 0
	depthHasIncreasedCounter := 0
	for _, chunk := range chunks {
		if len(chunk) < 3 {
			break
		}
		chunkSum := chunk[0] + chunk[1] + chunk[2]
		if previousWindow == 0 {
			previousWindow = chunkSum
			continue
		}

		if chunkSum > previousWindow {
			depthHasIncreasedCounter++
			previousWindow = chunkSum
			continue
		}
		previousWindow = chunkSum
	}

	fmt.Printf("Depth has increased %d times \n", depthHasIncreasedCounter)
}

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += 1 {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
