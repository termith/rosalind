package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	const (
		A   byte = 65
		C   byte = 67
		G   byte = 71
		T   byte = 84
		EOF byte = 10
	)

	dnaString, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	length := len(dnaString) - 1

	reverseCompliment := make([]byte, length+1)

	for i, n := range dnaString {
		switch n {
		case A:
			reverseCompliment[length-i] = T
		case C:
			reverseCompliment[length-i] = G
		case G:
			reverseCompliment[length-i] = C
		case T:
			reverseCompliment[length-i] = A
		case EOF:
			break
		}
	}

	fmt.Println(string(reverseCompliment))

}
