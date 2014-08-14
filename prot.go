/*
Problem

The 20 commonly occurring amino acids are abbreviated by using 20 letters from the English alphabet (all letters except for B, J, O, U, X, and Z). Protein strings are constructed from these 20 symbols. Henceforth, the term genetic string will incorporate protein strings along with DNA strings and RNA strings.

The RNA codon table dictates the details regarding the encoding of specific codons into the amino acid alphabet.

Given: An RNA string s corresponding to a strand of mRNA (of length at most 10 kbp).

Return: The protein string encoded by s.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	const EOF byte = 10

	var (
		rnaCodonTable map[string]string
		proteinString string
	)

	tableFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(tableFile, &rnaCodonTable)
	if err != nil {
		panic(err)
	}

	rnaString, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(rnaString)-3; i += 3 {
		if codon := string(rnaString[i : i+3]); rnaCodonTable[codon] == "X" {
			fmt.Println(proteinString)
			break
		} else {
			proteinString += rnaCodonTable[codon]
		}
	}

}
