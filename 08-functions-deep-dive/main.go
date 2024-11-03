package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
