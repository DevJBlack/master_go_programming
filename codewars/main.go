package main

func main() {

}

func evenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func opposite(value int) int {
	if value%2 == 0 {
		return 1 * value
	}
	return -1 * value

}

func positiveSum(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			result += numbers[i]
		}
	}
	return result
}

func positiveSums(numbers []int) (sum int) {
	for _, num := range numbers {
		if num > 0 {
			sum = sum + num
		}
	}
	return
}

func ppositiveSum(numbers []int) int {
	var result int
	for _, number := range numbers {
		if number >= 0 {
			result += number
		}
	}
	return result
}

func makeNegative(x int) int {
	if x >= 0 {
		return x * -1
	} else if x <= 0 {
		return x * 1
	} else {
		return x
	}

}

func makeMeNegative(x int) int {
	if x >= 0 {
		return -x
	}
	return x
}
