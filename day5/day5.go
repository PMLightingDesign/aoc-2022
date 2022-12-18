package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

const filename string = "test.txt"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func processHeader(header []string) [][]int {
	var stack [][]int
	
	// Determine the number of columns. The last line of the header is the column numbers
	var colCount int = 0
	var iRow string = header[len(header)-1]

	for _, c := range iRow {
		// If the character is not a space, convert it to an int and add it to the columns
		if c != ' ' {
			// If strconv.Atoi(c) is greater than colCount, set colCount to that value
			if i, err := strconv.Atoi(string(c)); err == nil && i > colCount {
				colCount = i
			}
		}
	}

	// Add a blank slice to the stack for each column
	for i := 0; i < colCount; i++ {
		stack = append(stack, []int{})
	}

	// Create the stack
	for i := len(header)-2; i >= 0; i-- {
		for j, c := range header[i] {
			// If the character is a '[', convert the next character to an int and add it to the row
			if c == '[' {
				stack[int(j/4)] = append(stack[int(j/4)], int(header[i][j+1]))
			}
		}
	}
	
	return stack
}

func moveStack(stack *[][]int, instruction string) {
	// Split the string using the space as a delimiter
	instructionParts := strings.Split(instruction, " ")
	count, err := strconv.Atoi(instructionParts[1])
	check(err)
	source, err := strconv.Atoi(instructionParts[3])
	check(err)
	source--
	destination, err := strconv.Atoi(instructionParts[5])
	check(err)
	destination--

	// Move count elements from source to destination which are indexes of the stack
	(*stack)[destination] = append((*stack)[destination], (*stack)[source][:count]...)
	(*stack)[source] = (*stack)[source][count:]
	
}

func printStack(stack [][]int) {
	println("------------------")
	// Stack is a slice of slices. Format and print it
	for i, row := range stack {
		// Print each value in the row as a character
		fmt.Printf("Row %d: ", i)
		for _, val := range row {
			fmt.Printf("%c", val)
			
		}
		fmt.Println()
	}
}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var header []string
	var headerFound bool = false
	var instructions []string

	for scanner.Scan() {
		// If the string is empty
		if len(scanner.Text()) == 0 {
			// This is the end of the header
			headerFound = true
			continue
		} else {
			if headerFound {
				instructions = append(instructions, scanner.Text())
			} else {
				header = append(header, scanner.Text())
			}
		}
	}

	stack := processHeader(header)
	printStack(stack)

	for _, instruction := range instructions {
		moveStack(&stack, instruction)
	}

	printStack(stack)

}