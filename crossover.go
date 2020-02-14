package ga

import (
	"math/rand"
	"strings"
)

func crossover(name string, parents []typeChromosome) typeChromosome {
	if strings.ToLower(name) == "onepoint" {
		return crossoverOnePoint(parents)
	} else {
		return typeChromosome{nil, -1}
	}
}

func crossoverOnePoint(parents []typeChromosome) typeChromosome {
	var point int = rand.Intn(5)

	var offspring typeChromosome
	offspring.genes = append(parents[0].genes[:point], parents[1].genes[point:]...)
	offspring.fitness = calcFitness(offspring)

	return offspring
}
