package main

/*Problem

The GC-content of a DNA string is given by the percentage of symbols in the string that are 'C' or 'G'. For example, the GC-content of "AGCTATAG" is 37.5%. Note that the reverse complement of any DNA string has the same GC-content.

DNA strings must be labeled when they are consolidated into a database. A commonly used method of string labeling is called FASTA format. In this format, the string is introduced by a line that begins with '>', followed by some labeling information. Subsequent lines contain the string itself; the first line to begin with '>' indicates the label of the next string.

In Rosalind's implementation, a string in FASTA format will be labeled by the ID "Rosalind_xxxx", where "xxxx" denotes a four-digit code between 0000 and 9999.

Given: At most 10 DNA strings in FASTA format (of length at most 1 kbp each).

Return: The ID of the string having the highest GC-content, followed by the GC-content of that string. Rosalind allows for a default error of 0.001 in all decimal answers unless otherwise stated; please see the note on absolute error below.
*/

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ID_START byte = 62
	C        byte = 67
	G        byte = 71
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	scanFile := bufio.NewScanner(file)

	dnaMap := make(map[string]string)

	var currentKey string

	for scanFile.Scan() {
		line := scanFile.Text()
		fmt.Println([]byte(line)[0])
		if []byte(line)[0] == ID_START {
			currentKey = line
			fmt.Println(currentKey)
		} else {
			dnaMap[currentKey] += line
		}
	}

	var maxPercent float64 = 0.0
	var maxId string

	for id, dna := range dnaMap {
		currentCounter := 0.0
		for _, n := range []byte(dna) {
			if n == C || n == G {
				currentCounter++
			}

			currentPercent := (currentCounter / float64(len(dna))) * 100.0
			if currentPercent > maxPercent {
				maxPercent = currentPercent
				maxId = id
			}

		}
	}
	fmt.Println(maxId, maxPercent)
}
