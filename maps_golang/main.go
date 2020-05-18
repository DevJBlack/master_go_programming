package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs:%d\n", len(employees))

	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	//employees["Dan"] = "Programmer"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Mary"] = 60.2
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 32.11,
		"EUR": 43.13,
		// 50: 322.1, //Error, all keys need to be the same
		"CHF": 3243.1,
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 10.27
	balances["DEV"] = 100.27
	fmt.Println(balances)
	fmt.Println(balances["RUN"])

	balances["RUN"] = 27.

	v, ok := balances["RUN"]

	if ok {
		fmt.Println("The RUN of a life time:", v)
	} else {
		fmt.Println("The RUN key is not a good run")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)

}
