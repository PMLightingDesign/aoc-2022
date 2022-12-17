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

	return -1
}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var bucket1 [127]bool
		var bucket2 [127]bool

		tSplit := len(scanner.Text()) / 2

		for i, char := range scanner.Text() {
			if i < tSplit {
				bucket1[int(char)] = true
			} else {
				bucket2[int(char)] = true
			}
		}

		for i := 0; i < 127; i++ {
			if bucket1[i] && bucket2[i] {
				total += scoreChar(byte(i))
			}
		}

	}

	println(total)

}