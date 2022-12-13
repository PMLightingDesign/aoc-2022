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

// A Rock 1
// B Paper, 2
// C Scissors 3

// Z Win = 6
// Y Draw = 3
// X Lose = 0

func play(line string) int {
	var score int = 0

	result := line[2]
	throw := line[0]

	if result == 'Z' {
		score = 6
		
		if(throw == 'A'){
			// We need paper to win
			score += 2
		} else if(throw == 'B'){
			// We need scissors to win
			score += 3
		} else if(throw == 'C'){
			// We need rock to win
			score += 1
		}

	} else if result == 'Y' {
		score = 3

		if(throw == 'A'){
			// We need rock to draw
			score += 1
		} else if(throw == 'B'){
			// We need paper to draw
			score += 2
		} else if(throw == 'C'){
			// We need scissors to draw
			score += 3
		}

	} else {
		score = 0

		if(throw == 'A'){
			// We need scissors to lose
			score += 3
		} else if(throw == 'B'){
			// We need rock to lose
			score += 1
		} else if(throw == 'C'){
			// We need paper to lose
			score += 2
		}
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