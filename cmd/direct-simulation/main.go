// Take a coin. Throw it 30 times. Is it a fair coin
// when you get 21 heads or tails?
package main

import (
	"fmt"
	"math/rand"
)

const treshold = 0.05

func main() {
	rand.Seed(42)

	minHeadOrTails := 21
	iterations := 1000
	counFlipAmount := 30

	isFair := isItAFairCoin(minHeadOrTails, iterations, counFlipAmount)

	fmt.Printf(`
Is a coint that you throw %d times and 
%d of them are heads / tails, a fair coin? 
is fair == %t`,
		counFlipAmount, minHeadOrTails, isFair)
}

func isItAFairCoin(minHeads, iterations, amount int) bool {
	counter := 0

	for i := 0; i < iterations; i++ {
		headsCount := sumRndInt(2, amount)

		if headsCount >= minHeads {
			counter++
		}
	}

	return (float64(counter) / float64(iterations)) > treshold
}

func sumRndInt(rnd, times int) int {
	sum := 0

	for i := 0; i < times; i++ {
		sum += rand.Intn(rnd)
	}

	return sum
}
