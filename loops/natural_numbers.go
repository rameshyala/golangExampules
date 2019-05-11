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
}
