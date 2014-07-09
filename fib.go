package main

/* Problem

Given: Positive integers n ≤ 40 and k ≤ 5.

Return: The total number of rabbit pairs that will be present after n months if we begin with 1 pair and in each generation, every pair of reproduction-age rabbits produces a litter of k rabbit pairs (instead of only 1 pair).*/

import (
	"fmt"
	"os"
	"strconv"
)

func KthFibonacci(k uint64) func() uint64 {
	var (
		a uint64 = 0
		b uint64 = 1
	)
	return func() uint64 {
		a, b = b, a*k+b
		return a
	}
}
func main() {

	n, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.ParseUint(os.Args[2], 10, 0)

	fibGenerator := KthFibonacci(k)

	for ; n > 0; n-- {
		fmt.Println(fibGenerator())
	}

}
