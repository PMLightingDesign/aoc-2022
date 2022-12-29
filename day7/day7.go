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

type File struct {
	name string
	size int
}

type Folder struct {
	children []Folder
	files	[]File
	size	int
}

func (f *Folder) addFile(file File) {
	f.files = append(f.files, file)
}

func (f *Folder) calculateSize() {
	f.size = 0

	for _, file := range f.files {
		f.size += file.size
	}

	for _, folder := range f.children {
		f.calculateSize()
		f.size += folder.size
	}
}

func main() {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// var root Folder

	for scanner.Scan() {
		line := scanner.Text()
		println(line)
	}
}


