package ga

import (
	"math/rand"
	"sort"
	"strings"
)

func replacement(name string, population []typeChromosome, offspring typeChromosome) []typeChromosome {
	if strings.ToLower(name) == "worst" {
		return replacementWorst(population, offspring)
	} else if strings.ToLower(name) == "best" {
		return replacementBest(population, offspring)
	} else {
		var prob int = rand.Intn(100)
		if prob <= 80 {
			return replacementWorst(population, offspring)
		} else {
			return replacementBest(population, offspring)
		}
	}
}

func replacementWorst(population []typeChromosome, offspring typeChromosome) []typeChromosome {
	if offspring.fitness >= population[0].fitness {
		population = append(population[1:], offspring)

		sort.Slice(population, func(i, j int) bool {
			return population[i].fitness < population[j].fitness
		})
	}

	return population
}

func replacementBest(population []typeChromosome, offspring typeChromosome) []typeChromosome {
	if offspring.fitness <= population[len(population)-1].fitness {
		population = append(population[1:], offspring)

		sort.Slice(population, func(i, j int) bool {
			return population[i].fitness < population[j].fitness
		})
	}

	return population
}
