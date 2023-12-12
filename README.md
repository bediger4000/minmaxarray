# Daily Coding Problem: Problem #1578 [Hard]

This problem was asked by Facebook.

Given an array of numbers of length `N`,
find both the minimum and maximum using less than `2 * (N - 2)` comparisons.

## Algorithm analysis

These interview problems often have a catch for length zero or one inputs.
This one seems to have that catch:

| N  | 2(N - 2) |
|----|----------|
| 0 | -4 |
| 1 | -2 |
| 2 | 0 |
| 3 | 2 |
| 4 | 4 |
| 5 | 6 |

What is a program going to do, give back some comparisons for 0- or 1-length arrays?
For arrays of length 2, zero comparisons doesn't make sense.
Should the interview candidate point this out?

The [obvious algorithm](a0.go) sets variables `minimum` and `maximum` to the smallest number and largest number
representable by whatever type the variables possess,
then plows through all of the array one element at a time comparing variables `minimum` and
`maximum` to each element.
That would require `2*N` comparisons,
one to check if array element is less than current value of `minimum`
one to check if array element is greater than current value of `maximum`.

The [first change](a1.go) to make to get to the goal is to observe
that you can set `minimum` and `maximum` to the value of the first
element of the array instead of smallest and largest representable values.
You can start compariing with the *second* element of the array.
You've avoided 2 comparisons, so exactly `2*(N - 2)` comparisons.

## Interview analysis
