package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	devin := []string{"Devin James Black"}

	for _, name := range devin {
		fmt.Printf("%q", strings.Split(name, ""))
	}
}
