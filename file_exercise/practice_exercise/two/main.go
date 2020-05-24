package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// One way, open the file and then use the ioutil readAll function

	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal()
	// }
	// defer file.Close()

	// b, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inside of info.txt: %s", b)

	// The other way, use ioutil ReadFile function

	content, err := ioutil.ReadFile("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Content inside of info.txt: %s", content)

}
