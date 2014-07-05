package main

/* Problem

An RNA string is a string formed from the alphabet containing 'A', 'C', 'G', and 'U'.

Given a DNA string t corresponding to a coding strand, its transcribed RNA string u is formed by replacing all occurrences of 'T' in t with 'U' in u.

Given: A DNA string t having length at most 1000 nt.

Return: The transcribed RNA string of t. */
import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	EOF byte = 10
	T   byte = 84
	U   byte = 85
)

func main() {
	rna, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	for i, n := range rna {
		switch n {
		case T:
			rna[i] = U
		case EOF:
			fmt.Println(string(rna))
			break
		}
	}

}
