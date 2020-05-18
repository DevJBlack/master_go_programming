package main

import "fmt"

func main() {
	// 1.
	var m1 map[float64]bool
	fmt.Printf("m1 type: %T, m1 value: %#v\n", m1, m1)

	// 2.
	m2 := map[int]string{1: "Sting", 2: "Queen"}

	// 3.
	m2[10] = "Abba"

	// 4
	fmt.Println(m2[2])   // existing key
	fmt.Println(m2[100]) // non-existing key

	var m11 map[int]bool

	// ERROR -> panic: assignment to entry in nil map
	// m1[5] = true

	m22 := map[int]int{3: 10, 4: 40}
	m33 := map[int]int{3: 10, 4: 40}

	// ERROR -> invalid operation: m2 == m3 (map can only be compared to nil)
	// fmt.Println(m2 == m3)

	_, _, _ = m11, m22, m33

	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 1)

	for k, v := range m {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}
}
