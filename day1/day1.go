package main

import (
	"bufio"
	"os"
	"strconv"
)

const filename string = "day1.txt"

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

	var max int = 0
	var max2 int = 0
	var max3 int = 0

	for _, val := range cals {
		if val > max {
			max3 = max2
			max2 = max
			max = val
		}
	}

	println(max)
	println(max2)
	println(max3)

	println(max + max2 + max3)
	
}