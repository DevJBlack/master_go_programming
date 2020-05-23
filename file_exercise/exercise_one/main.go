package main

import (
	"log"
	"os"
)

func main() {

	// Step 1
	// Created a file and then closed it
	// file, err := os.Create("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()

	// Step 2
	// checking if file exists
	// _, err := os.Stat("info.txt")
	// // error handling
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("The file does not exist!")
	// 	}
	// }

	// // renaming the file
	// err = os.Rename("info.txt", "data.txt")
	// // error handling
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Step 3
	// Remove the file
	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}

}
