package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func boooolToWord(word bool) string {

	var maybe string

	switch {
	case word == true:
		return "Yes"
	case word == false:
		return "No"
	}
	return maybe
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

func ssquareSum(numbers []int) int {
	count := 0
	for _, v := range numbers {
		v *= v
		count += v
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

func getCount(strn string) int {
	count := 0

	vowels := []string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		count += strings.Count(strn, vowel)
	}

	return count
}

func getCounts(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func getsCount(str string) int {
	count := 0
	for i := range str {
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func vowel() {
	//a, e, i, o, u
	devin := "My name is Devin and I love Golang!"

	count := 0

	for _, vow := range devin {

		if vow == 'a' || vow == 'e' || vow == 'i' || vow == 'o' || vow == 'u' {
			count++
		}
	}
	fmt.Println("here are the vowels in the string", count)

}

func highAndLow(in string) string {
	d := strings.Fields(in)
	min := 0
	max := 0

	for i, s := range d {
		n, _ := strconv.Atoi(s)
		if i == 0 || n < min {
			min = n
		}
		if i == 0 || n > max {
			max = n
		}
	}

	return strconv.Itoa(max) + " " + strconv.Itoa(min)
}

func otherAngle(a int, b int) int {
	return 180 - (a + b)
}

func otherAngles(a int, b int) int {
	var c int
	c = 180 - a - b

	return (c)
}

func grow(arr []int) int {
	if len(arr) == 0 {
		return 1
	}
	return arr[0] * grow(arr[1:])
}

func grows(arr []int) int {
	x := 1
	for _, v := range arr {
		x *= v
	}
	return x
}
