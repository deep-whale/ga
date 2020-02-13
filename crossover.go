package ga

import (
	"math/rand"
	"strings"
)

func crossover(name string, parents []typeChromosome) typeChromosome {
	if strings.ToLower(name) == "onepoint" {
		return crossoverOnePoint(parents)
	} else {
		var offspring typeChromosome
		offspring.size = -1
		offspring.genes = nil
		offspring.fitness = -1

		return offspring
	}
}

func crossoverOnePoint(parents []typeChromosome) typeChromosome {
	var point int = rand.Intn(parents[0].size)

	var offspring typeChromosome
	offspring.genes = append(parents[0].genes[:point], parents[1].genes[point:]...)
	offspring.size = len(offspring.genes)
	offspring.fitness = calcFitness(offspring)

	return offspring
}
