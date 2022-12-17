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

func scoreChar(char byte) int {
	if int(char) < 91{
		return (int(char) - 64) + 26
	} else {
		return int(char) - 96
	}
}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

			if len(lines) >= 3 {

			var l1, l2, l3 [127]bool

			for i := 0; i < len(lines[0]); i++ {
				l1[int(lines[0][i])] = true
			}

			for i := 0; i < len(lines[1]); i++ {
				l2[int(lines[1][i])] = true
			}

			for i := 0; i < len(lines[2]); i++ {
				l3[int(lines[2][i])] = true
			}

			// Check if the same character appears in all three lines
			for i := 0; i < len(l1); i++ {
				if l1[i] && l2[i] && l3[i] {
					total += scoreChar(byte(i))
				}
			}

			lines = []string{}

		}
	}

	println(total)

}