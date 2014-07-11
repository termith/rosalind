package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

const (
	ID_START byte = 62
	C        byte = 67
	G        byte = 71
)

func main() {

	//Read file and move content to the map stringID -> dnaString

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
	// Go throw map and search max gcRate

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
