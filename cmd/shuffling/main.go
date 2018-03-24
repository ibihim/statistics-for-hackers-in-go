// There are two groups A and B. Is the group A
// a small elite? Their average in scores is higher
// than group B's average.
// To verify that, it is necessary to shuffle the results
// with constant group sizes. While doing so, the
// difference is logged. So there will be differences
// of different sizes. How many of them have the same
// difference as the current result? If less than 5% of
// the group have the current difference, the difference
// seems to be statistical significant.
package main

import (
	"fmt"

	"github.com/ibihim/statistics-for-hackers-in-go/pkg/statutils"

	"gonum.org/v1/gonum/stat"
)

const treshold = 0.05

func main() {
	// analysis of the following results
	resultA := []float64{84, 72, 57, 46, 63, 76, 99, 91}
	resultB := []float64{81, 69, 74, 61, 56, 87, 69, 65, 66, 44, 62, 69}
	discussedDiff := stat.Mean(resultA, nil) - stat.Mean(resultB, nil)
	fmt.Printf("diff of avg A and avg B: %v\n", discussedDiff)

	// shuffle setup
	iterations := 1000
	delimiter := len(resultA)
	// is being mutated!!
	dataSet := append(resultA, resultB...)

	// shuffle and validate
	isValid := isGroupABetter(iterations, delimiter, discussedDiff, dataSet)
	fmt.Printf("is the discrepancy statistically significant: %v\n", isValid)
}

func isGroupABetter(iterations, delimiter int, match float64, dataSet []float64) bool {
	// shuffle and and count diffs of avgs that are equal or higher
	counter := 1 // starts at one as discussed diff is a match
	for i := 1; i < iterations; i++ {
		statutils.Shuffle(dataSet)
		if diff := splitAndCalcDiffAvg(dataSet, delimiter); diff >= match {
			counter++
		}
	}

	// calculate percentage and check if the discussed diff is within the treshold
	p := float64(counter) / float64(iterations)
	return p < treshold
}

func splitAndCalcDiffAvg(completeSet []float64, delimiter int) float64 {
	avgA := stat.Mean(completeSet[:delimiter], nil)
	avgB := stat.Mean(completeSet[delimiter:], nil)

	return avgA - avgB
}
