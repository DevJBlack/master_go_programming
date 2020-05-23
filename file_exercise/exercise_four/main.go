package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("Gophers is the Golang Mascot")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := ioutil.ReadFile("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Words in info: %s", newFile)
}
