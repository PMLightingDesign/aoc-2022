package main

import (
	"bufio"
	"os"
	"strconv"
	"sort"
)

const filename string = "test.txt"

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var sum int = 0
	var cals []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// If the line is not empty
		if len(line) > 0 {
			// parse the line as an int
			// and add it to the sum
			val, err:= strconv.Atoi(line)
			check(err)
			sum += val
		} else {
			cals = append(cals, sum)
			sum = 0
		}
	}

	// Sort the list of calories
	sort.Ints(cals)

	// Print the top 3
	var final int = 0
	for i := 0; i < 3; i++ {
		final += cals[len(cals)-1-i]
	}

	println(final)
	
}