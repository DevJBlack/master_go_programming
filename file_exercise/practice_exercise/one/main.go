package main

import (
	"log"
	"os"
)

func main() {
	// Step one, create a file

	// file, err := os.Create("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// Step Two, rename the file

	// _, err := os.Stat("info.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("The file does not exist!")
	// 	}
	// }

	// err = os.Rename("info.txt", "data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Step three, remove the file
	err := os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}
