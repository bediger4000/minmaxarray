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

	min := ary[0]
	max := ary[1]
	comparisonCount := 1
	if min > max {
		min, max = max, min
	}

	for i := 2; i < N; i++ {
		comparisonCount++
		if ary[i] < min {
			min = ary[i]
			continue
		}
		comparisonCount++
		if ary[i] > max {
			max = ary[i]
		}
	}

	fmt.Printf("2(N - 2) = %d\n", 2*(N-2))
	phrase := "does not meet criterion"
	if comparisonCount < 2*(N-2) {
		phrase = "meets criterion"
	}
	fmt.Printf("made %d comparison, 2(N-2)+1 = %d, %s\n", comparisonCount, 2*(N-2)+1, phrase)
	fmt.Printf("min\t%d\nmax\t%d\n", min, max)
}