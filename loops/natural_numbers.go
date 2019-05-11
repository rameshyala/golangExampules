package main

import (
	"fmt"
)

func main() {
	n := 10
	displayNaturalNumbers(n)
	displayRevNaturalNumbers(n)
	displayEvenNumbers(n)
	displayOddNumbers(n)
	displayNaturalNumSum(n)
	displayAlphabits()
	displayMultiplicationTable(n)

	n = 15851
	result := countNumOfDigits(n)
	fmt.Println("Number of digites in the number: ", len(result))
	fmt.Println("Sum of first and last digit sum is: ", result[0]+result[len(result)-1])

	sum := sumOfDigits(n)
	fmt.Println("Sum of all digits: ", sum)

	rev := reverseNumber(n)
	fmt.Println("reverse of this number is:", rev)

	powv := powerOfValue(2, 5)
	fmt.Println("2 power of 5 is", powv)
}

func displayNaturalNumbers(n int) {
	for index := 1; index < n; index++ {
		fmt.Print(" ", index)
	}
	fmt.Println()
}

func displayRevNaturalNumbers(n int) {
	for index := n; index != 0; index-- {
		fmt.Print(" ", index)
	}
	fmt.Println()
}

func displayEvenNumbers(n int) {
	sum := 0
	for index := 1; index <= n; index++ {
		if index%2 == 0 {
			sum = sum + index
			fmt.Print(" ", index)
		}
	}
	fmt.Println()
	fmt.Println("even numbers sum: ", sum)
}

func displayOddNumbers(n int) {
	sum := 0
	for index := 1; index <= n; index++ {
		if index%2 != 0 {
			sum = sum + index
			fmt.Print(" ", index)
		}
	}
	fmt.Println()
	fmt.Println("odd numbers sum: ", sum)
}

func displayNaturalNumSum(n int) {
	sum := 0
	for index := 1; index <= n; index++ {
		sum = sum + index
	}
	fmt.Println(" Total ", sum)
}

func displayAlphabits() {
	for j := 'a'; j <= 'z'; j++ {
		fmt.Printf("%c ", j)
	}
	fmt.Println()
	fmt.Println()
}

func displayMultiplicationTable(n int) {
	for index := 1; index <= n; index++ {
		fmt.Println(n, "X", index, "=", n*index)
	}
}

func countNumOfDigits(n int) []int {
	var result []int

	if n < 9 {
		return append(result, n)
	}

	for index := 1; n != 0; index++ {
		r := n % 10
		result = append(result, r)
		n = n / 10
	}

	return result
}

func sumOfDigits(n int) int {
	sum := 0
	if n < 9 {
		return n
	}

	for index := 1; n != 0; index++ {
		r := n % 10
		sum = sum + r
		n = n / 10
	}

	return sum
}

func reverseNumber(n int) int {
	if n < 10 {
		return n
	}

	reverse := 0
	for index := 1; n != 0; index++ {
		reverse = (n % 10) + (reverse * 10)
		n = n / 10
	}

	return reverse
}

func powerOfValue(num, pow int) int {
	result := 1
	if pow == 0 {
		return num
	}

	for i := 1; i <= pow; i++ {
		result = result * num
	}

	return result
}
