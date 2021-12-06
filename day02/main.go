package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	r := regexp.MustCompile(`(forward|up|down) (\d+)`)

	p1_xPos := 0
	p1_yPos := 0

	p2_xPos := 0
	p2_yPos := 0
	p2_aim := 0

	for _, line := range lines {

		submatch := r.FindSubmatch([]byte(line))

		direction := string(submatch[1])
		value, _ := strconv.Atoi(string(submatch[2]))

		switch direction {
		case "forward":
			p1_xPos += value
			p2_xPos += value
			p2_yPos += p2_aim * value
		case "up":
			p1_yPos -= value
			p2_aim -= value
		case "down":
			p1_yPos += value
			p2_aim += value
		default:
			log.Printf("%s is not a valid direction (v=%d)", direction, value)
		}

	}

	//Part One
	p1_answer := p1_xPos * p1_yPos
	p2_answer := p2_xPos * p2_yPos

	fmt.Println("Day 02 ANSWERS")
	fmt.Printf("Part 1: %d\n  Final Position: %d\n  Final Depth: %d\n", p1_answer, p1_xPos, p1_yPos)
	fmt.Printf("Part 2: %d\n  Final Position: %d\n  Final Depth: %d\n  Final Aim: %d\n", p2_answer, p2_xPos, p2_yPos, p2_aim)

}
