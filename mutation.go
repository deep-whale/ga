package ga

import (
	"math/rand"
	"strings"
)

func mutation(name string, offspring typeChromosome, rate float32) typeChromosome {
	if strings.ToLower(name) == "flip" {
		return mutationFlip(offspring, rate)
	} else {
		return typeChromosome{nil, -1}
	}
}

// mutationFlip is function to perform bit flip.
func mutationFlip(offspring typeChromosome, rate float32) typeChromosome {
	if rate <= 0.00001 {
		rate = rate * 10
	}

	if rand.Intn(1000) <= int(rate*1000) {
		var point int = rand.Intn(5)

		if offspring.genes[point] == 0 {
			offspring.genes[point] = 1
		} else {
			offspring.genes[point] = 0
		}

		offspring.fitness = calcFitness(offspring)
	}

	return offspring
}
