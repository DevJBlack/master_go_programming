package main

import (
	"fmt"
)

func si(slice []string, myStr string) bool {

	for _, v := range slice {
		if v == myStr {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"cat", "bear", "horse", "lion", "snake"}
	result := si(animals, "kitty") //false //"bear" == true

	fmt.Println(result)
}
