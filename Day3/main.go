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
	// gammas := []int{}
	// epsilons := []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := [][]int{}
	for scanner.Scan() {
		row := scanner.Text()
		rows := strings.Split(row, "")

		parsed := []int{}
		for _, r := range rows {
			p, _ := strconv.Atoi(r)
			parsed = append(parsed, p)
		}

		instructions = append(instructions, parsed)
	}

	fmt.Printf("Power consumption: %d\n", GetValues(instructions))
	fmt.Printf("Lifesupport rating: %d\n", GetValues2(instructions))
}

func GetValues(rows [][]int) int64 {
	result := map[int]string{}

	for childIndex := 0; childIndex < len(rows[0]); childIndex++ {
		for parentIndex := 0; parentIndex < len(rows); parentIndex++ {
			result[childIndex] += strconv.Itoa(rows[parentIndex][childIndex])
		}
	}

	gamma := ""
	epsilon := ""
	for index := 0; index < len(result); index++ {

		one := strings.Count(result[index], "1")
		zero := strings.Count(result[index], "0")

		if one > zero {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaConverted, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonConverted, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaConverted * epsilonConverted
}

func GetValues2(rows [][]int) int64 {
	var result, result2 []int

	result = RemoveFromSlice(rows, 0, false)
	result2 = RemoveFromSlice(rows, 0, true)

	a := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), ""), "[]")
	b := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result2)), ""), "[]")

	oxyginRating, _ := strconv.ParseInt(a, 2, 64)
	co2scrubberrating, _ := strconv.ParseInt(b, 2, 64)

	return oxyginRating * co2scrubberrating
}

func RemoveFromSlice(rows [][]int, index int, isCO2Rating bool) []int {
	copy := [][]int{}
	if len(rows) == 1 {
		return rows[0]
	}

	one, zero := 0, 0
	for _, row := range rows {
		if row[index] == 1 {
			one++
		} else {
			zero++
		}
	}

	if !isCO2Rating {
		if one > zero || one == zero {
			for _, row := range rows {
				if row[index] == 1 {
					copy = append(copy, row)
				}
			}
		} else {
			for _, row := range rows {
				if row[index] == 0 {
					copy = append(copy, row)
				}
			}
		}
	} else {
		if zero < one || zero == one {
			for _, row := range rows {
				if row[index] == 0 {
					copy = append(copy, row)
				}
			}
		} else {
			for _, row := range rows {
				if row[index] == 1 {
					copy = append(copy, row)
				}
			}
		}
	}

	return RemoveFromSlice(copy, index+1, isCO2Rating)
}
