package pimontecarlo

import (
"math/rand"
"math"
)

// GetPi calculates the value of Pi using the Monte Carlo method.
//
// iterations: the number of iterations to perform.
//
// Returns the calculated value of Pi.
func GetPi(iterations int) float64{
	var result,valorPi,mediaPi,currFloat,x,y,z float64
	var currValue int64
	
	for i := 0; i < iterations; i++ {
		currValue = 0
		for j := 0; j < iterations; j++ {
			x = rand.Float64()
			y = rand.Float64()
			z = math.Sqrt(x*x + y*y)
			if (z <= 1) {
				currValue += 1
			}
		}
		currFloat = float64(currValue)
		valorPi = currFloat * 4 / float64(iterations)
		mediaPi += valorPi
	}
	result = mediaPi / float64(iterations)
	return result
}
