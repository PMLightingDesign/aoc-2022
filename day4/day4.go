package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

const filename string = "input.txt"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type numRange struct {
	min int
	max int
}

func (r *numRange) createRange(s string) {
	min, err := strconv.Atoi(strings.Split(s, "-")[0])
	check(err)
	r.min = min

	max, err := strconv.Atoi(strings.Split(s, "-")[1])
	check(err)
	r.max = max
}

func (r *numRange) contains(r2 numRange) bool {
	return r.min <= r2.min && r.max >= r2.max
}

func checkRow(s string) bool {
	strings := strings.Split(s, ",")

	r1 := numRange{}
	r1.createRange(strings[0])
	
	r2 := numRange{}
	r2.createRange(strings[1])

	println(r1.min, "-", r1.max, " | ", r2.min, "-", r2.max)

	if r1.contains(r2) || r2.contains(r1) {
		println("Contained")
		return true
	}

	println("Not contained")
	return false
}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var count int

	for scanner.Scan() {
		if checkRow(scanner.Text()) {
			count++
		}
	}

	println(count)

}