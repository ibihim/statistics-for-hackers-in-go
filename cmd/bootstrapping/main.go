// How high can the tower grow?
//  - What is the mean of the height of the stack?
//  - What is the uncertainty on this estimate?

package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/stat"
)

const iterations = 10000

func main() {
	data := []float64{
		48, 24, 32, 61, 51, 12, 32, 18, 19, 24,
		21, 41, 29, 21, 25, 23, 42, 18, 23, 13,
	}

	avgData := make([]float64, iterations)
	for i := 0; i < iterations; i++ {
		sample := generateRndWithReplacement(data)
		avgData[i] = stat.Mean(sample, nil)
	}

	avg := stat.Mean(avgData, nil)
	std := stat.StdDev(avgData, nil)

	fmt.Printf("avg: %f\nstd: %f\n", avg, std)
}

func generateRndWithReplacement(data []float64) []float64 {
	dataLen := len(data)
	rndData := make([]float64, dataLen)

	for i := 0; i < dataLen; i++ {
		pos := rand.Intn(dataLen)
		rndData[i] = data[pos]
	}

	return rndData
}
