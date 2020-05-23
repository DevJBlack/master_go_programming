package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// First Attempt at opening file with bufio
	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // defer file.Close()

	// scan := bufio.NewScanner(file)

	// for scan.Scan() {
	// 	fmt.Println(scan.Text())
	// }

	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("info.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	// reading the whole file line by line:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// checking for any possible errors:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
