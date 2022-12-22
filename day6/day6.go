package main

import (
	"bufio"
	"os"
	// "strings"
	// "strconv"
	// "fmt"
)

const filename string = "input.txt"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewReader(f)

	var i int
	var bytes []byte

	const groupSize int = 14

	for {
		b, err := scanner.ReadByte()
		if err != nil {
			println("End of file")
			break
		}

		bytes = append(bytes, b)
		i++

		if i < groupSize {
			continue
		}

		var dupe bool
		var charList [127]bool

		// Loop through the last groupSize bytes
		for j := i - groupSize; j < i; j++ {
			if charList[bytes[j]] {
				dupe = true
				break
			}
			charList[bytes[j]] = true
		}

		if !dupe {
			println("We had to read ", i, " bytes to find string wihtout duplicates")
			break
		}

	}

}