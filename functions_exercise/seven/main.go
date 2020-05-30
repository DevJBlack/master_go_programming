package main

import (
	"fmt"
	"strings"
)

func searchI(pets []string, one string) bool {

	for _, v := range pets {
		if strings.EqualFold(one, v) {
			return true
		}
	}
	return false
}

func main() {
	newPet := []string{"cat", "dog", "horse", "liama", "alpacha"}
	result := searchI(newPet, "hOrSe")

	fmt.Println(result)
}
