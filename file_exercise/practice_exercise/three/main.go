package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Opening the file
	openFile, err := os.Open("info.txt")
	// Error Handling
	if err != nil {
		log.Fatal(err)
	}

	// Using bufio for NewScanner of file
	scan := bufio.NewScanner(openFile)

	// Creating a for loop to read the entire file
	for scan.Scan() {
		fmt.Println(scan.Text())
	}

	// Error handling
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}
