package main

/*Problem

A string is simply an ordered collection of symbols selected from some alphabet and formed into a word; the length of a string is the number of symbols that it contains.

An example of a length 21 DNA string (whose alphabet contains the symbols 'A', 'C', 'G', and 'T') is "ATGCTTCAGAAAGGTCTTACG."

Given: A DNA string s of length at most 1000 nt.

Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.*/

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	A_counter int = 0
	C_counter int = 0
	G_counter int = 0
	T_counter int = 0
)

const (
	A   byte = 65
	C   byte = 67
	G   byte = 71
	T   byte = 84
	EOF byte = 10
)

func main() {

	dna, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, n := range dna {
		switch n {
		case A:
			A_counter++
		case C:
			C_counter++
		case G:
			G_counter++
		case T:
			T_counter++
		case EOF:
			break
		}
	}

	if A_counter+C_counter+G_counter+T_counter != len(dna)-1 {
		panic(errors.New("Wrong answer!"))
	} else {
		fmt.Println(A_counter, C_counter, G_counter, T_counter)
	}
}
