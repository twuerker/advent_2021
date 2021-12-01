package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readNumberLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("skipping line: %s", line)
		} else {
			numbers = append(numbers, number)
		}
	}
	return numbers, scanner.Err()
}

func main() {

	numbers, err := readNumberLines("input.txt")
	if err != nil {
		log.Fatalf("readNumberLines: %s", err)
	}

	answer1 := 0
	answer2 := 0

	for i, number := range numbers {
		//Part One
		if i >= 1 && number > numbers[i-1] {
			answer1++
		}
		//Part Two
		// (n[i] + n[i-1] + n[i-2]) > (n[i-1] + n[i-2] + n[i-3])
		// n[i] > n[i-3]
		if i >= 3 && number > numbers[i-3] {
			answer2++
		}
	}

	fmt.Printf("Day 01 ANSWERS\n  Part 1: %d\n  Part 2: %d\n", answer1, answer2)

}
