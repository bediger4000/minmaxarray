package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

/*
 * Given an array of numbers of length `N`,
 * find both the minimum and maximum using less than `2 * (N - 2)` comparisons.
 *
 * Extremely naive version.
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
	max := math.MinInt
	min := math.MaxInt

	for i := 0; i < N; i++ {
		comparisonCount++
		if ary[i] < min {
			min = ary[i]
		}
		comparisonCount++
		if ary[i] > max {
			max = ary[i]
		}
	}

	fmt.Printf("2(N - 2) = %d\n", 2*(N-2))
	fmt.Printf("made %d comparison\n", comparisonCount)
	fmt.Printf("min\t%d\nmax\t%d\n", min, max)
}
