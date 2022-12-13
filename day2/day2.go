package main

import (
	"bufio"
	"os"
)

const filename string = "input.txt"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// A, X Rock, 1
// B, Y Paper, 2
// C, Z Scissors, 3

// Win = 6
// Draw = 3
// Lose = 0

func play(line string) int {
	var score int = 0

	enemy := line[0]
	you := line[2]

	if (enemy == 'A' && you == 'X') || (enemy == 'B' && you == 'Y') || (enemy == 'C' && you == 'Z') {
		score += 3
	} else if (enemy == 'A' && you == 'Y') || (enemy == 'B' && you == 'Z') || (enemy == 'C' && you == 'X') {
		score += 6
	}

	if you == 'X' {
		score += 1
	} else if you == 'Y' {
		score += 2
	} else if you == 'Z' {
		score += 3
	}

	return score

}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var score int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		result :=  play(line)
		println(result)
		score += result
	}

	println(score)

}