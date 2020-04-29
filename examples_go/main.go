package main

import "fmt"

func main() {

	const (
		distance = 149_600_000
		speed    = 299_792_458 / 1000
	)

	const time = distance / speed
	fmt.Printf("The sunlight reaches Earth in %v seconsd.\n", time)

	const (
		distance1 = 149_600_000 * 1000 // distance from the Sun to Earth in m
		//(it's allowed to use underscores in numbers for a better readability)

		speed1 = 299_792_458 // speed of light in m / s
	)

	const time1 = distance1 / speed1 // time in seconds

	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time1)

}
