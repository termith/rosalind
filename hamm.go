package main

/*Problem

Given two strings s and t of equal length, the Hamming distance between s and t, denoted dH(s,t), is the number of corresponding symbols that differ in s and t. See Figure 2.

Given: Two DNA strings s and t of equal length (not exceeding 1 kbp).

Return: The Hamming distance dH(s,t).
*/

import (
	"fmt"
	"io/ioutil"
	"os"
)

const NEW_STRING byte = 10

var hammDistance int = 0


func main() {

	firstString := make([]byte, 0)
	secondString := make([]byte, 0)

	dnaArray, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	for i, n := range dnaArray {
		if n == NEW_STRING {
			firstString = dnaArray[0:i]
			secondString = dnaArray[i+1:]
			break
		}
	}

	for j := 0; j < len(firstString); j++ {
		if firstString[j] != secondString[j] {
			hammDistance++
		}
	}

	fmt.Println(hammDistance)
}
