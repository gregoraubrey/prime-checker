package main

import "fmt"

func main() {
	num := 99

	_, msg := isPrime(num)
	fmt.Println(msg)
}

func isPrime(num int) (bool, string) {
	if num == 0 || num == 1 {
		return false, fmt.Sprintf("%d is not prime.", num)
	}

	if num < 0 {
		return false, fmt.Sprintf("%d is not a positive number, so it cannot be prime.", num)
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false, fmt.Sprintf("%d is not prime because it is divisible by %d.", num, i)
		}
	}

	return true, fmt.Sprintf("%d is prime!", num)
}