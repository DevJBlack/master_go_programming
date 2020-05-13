package main

import "strings"

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

func repeatStr(repititions int, value string) (result string) {
	for i := 0; i < repititions; i++ {
		result += value
	}
	return
}

func boolToWord(word bool) string {
	if word == true {
		return ("Yes")
	}
	return ("No")
}

func boolToWords(word bool) string {
	return map[bool]string{false: "No", true: "Yes"}[word]
}

func summation(n int) int {

	count := 0

	for i := 0; i <= n; i++ {
		count += i

	}
	return (count)

}

func solution(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}
	return (result)
}

func squareSum(numbers []int) int {
	count := 0

	for _, v := range numbers {
		a := v * v
		count += a
	}
	return (count)
}

func squareSums(numbers []int) int {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum = sum + (numbers[i] * numbers[i])
	}

	return sum
}

func squaresSum(nums []int) (res int) {
	for _, val := range nums {
		res += val * val
	}
	return res
}

func removeChar(word string) string {
	return (word[1 : len(word)-1])
}

func removesChar(word string) string {
	return strings.Join(strings.Split(word, "")[1:len(strings.Split(word, ""))-1], "")
}

func removeChars(word string) string {
	var newWord = []rune(word)
	return string(newWord[1 : len(newWord)-1])
}

func isDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func issDivisible(n, x, y int) bool {
	return n%x+n%y == 0
}

func isDivisibles(n, x, y int) bool {
	var out bool

	if n%x == 0 && n%y == 0 {
		out = true
	}

	return out
}

var greet = func() string { return "hello world!" }

func greets() string {
	b := "hello world!"
	return b
}
