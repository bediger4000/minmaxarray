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
 * Algorithm from stackexchange:
 * https://stackoverflow.com/questions/13544476/how-to-find-max-and-min-in-array-using-minimum-comparisons
 * 1. Pick 2 elements(a, b), compare them. (say a > b)
 * 2. Update min by comparing (min, b)
 * 3. Update max by comparing (max, a)
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

	minimum := ary[0]
	maximum := ary[1]
	if minimum > maximum {
		minimum, maximum = maximum, minimum
	}
	comparisonCount := 1

	var i int
	for i = 2; i < N-1; i += 2 {
		a := ary[i]
		b := ary[i+1]
		comparisonCount++
		if b > a {
			a, b = b, a
		}
		comparisonCount++
		if b < minimum {
			minimum = b
		}
		comparisonCount++
		if a > maximum {
			maximum = a
		}
	}

	if i == (N - 1) {
		comparisonCount++
		if ary[i] < minimum {
			minimum = ary[i]
		} else if ary[i] > maximum {
			comparisonCount++
			maximum = ary[i]
		}
	}

	fmt.Printf("2(N - 2) = %d\n", 2*(N-2))
	phrase := "does not meet criterion"
	if comparisonCount < 2*(N-2) {
		phrase = "meets criterion"
	}
	fmt.Printf("made %d comparison, 3N/2-2 = %d, %s\n", comparisonCount, 3*N/2-2, phrase)
	fmt.Printf("min\t%d\nmax\t%d\n", minimum, maximum)
}
