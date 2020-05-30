// Naked return
package main

import (
	"fmt"
)

func all(sum ...int) (value int) {
	for _, v := range sum {
		value += v
	}

	return
}

func main() {
	a := all(1, 2, 3, 4, 5)
	fmt.Println(a)
}
