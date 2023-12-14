package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
 * Given an array of numbers of length `N`,
 * find both the minimum and maximum using less than `2 * (N - 2)` comparisons.
 */

func main() {
	var ary []int
	for _, str := range os.Args[1:] {
		m, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ary = append(ary, m)
	}

	N := len(ary)

	fmt.Printf("Array N %d: %v\n", N, ary)

	comparisonCount := 0
	minimum := ary[0]
	maximum := ary[0]
	for i := 1; i < N; i++ {
		comparisonCount++
		if ary[i] < minimum {
			minimum = ary[i]
		}
		comparisonCount++
		if ary[i] > maximum {
			maximum = ary[i]
		}
	}

	phrase := "did not pass criterion"
	if comparisonCount < 2*(N-2) {
		phrase = "did pass criterion"
	}

	fmt.Printf("2(N - 2) = %d, %s\n", 2*(N-2), phrase)
	fmt.Printf("made %d comparison, 2*(N-1) = %d\n", comparisonCount, 2*N-1)
	fmt.Printf("min\t%d\nmax\t%d\n", minimum, maximum)
}
