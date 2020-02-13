package ga

import (
	"fmt"
	"math/rand"
	"sort"
)

// typeGenes is type of genes in chromosome.
type typeGenes []int

// typeFitness is type of fitness value for chromosome.
type typeFitness int

// typeChromosome is a structure of chromosome.
type typeChromosome struct {
	size    int
	genes   typeGenes
	fitness typeFitness
}

// calcFitness is function to calculate fitness of chromosome.
func calcFitness(chromosome typeChromosome) typeFitness {
	if chromosome.size == -1 {
		return -1
	}

	var fitness typeFitness = 0
	for i := 0; i < chromosome.size; i++ {
		if chromosome.genes[i] == 1 {
			fitness = fitness + 1
		}
	}

	return fitness
}

// initChromosome is function to create and initialize chromosome.
func initChromosome(sizeChromosome int) typeChromosome {
	var chromosome typeChromosome
	chromosome.size = sizeChromosome
	chromosome.genes = make([]int, sizeChromosome)
	chromosome.fitness = -1

	for i := 0; i < sizeChromosome; i++ {
		chromosome.genes[i] = rand.Intn(2)
	}
	chromosome.fitness = calcFitness(chromosome)

	return chromosome
}

// initPopulation is function to create population.
func initPopulation(sizeChromosome int, sizePopulation int) []typeChromosome {
	var population []typeChromosome = make([]typeChromosome, sizePopulation)
	for i := 0; i < sizePopulation; i++ {
		population[i] = initChromosome(sizeChromosome)
	}

	sort.Slice(population, func(i, j int) bool {
		return population[i].fitness < population[j].fitness
	})

	return population
}

func Run(sizeChromosome int, sizePopulation int, numGeneration int) {
	const selectionName string = "roulette"
	const crossoverName string = "onepoint"
	const mutationName string = "flip"
	const replacementName string = "worst"

	var population []typeChromosome = initPopulation(sizeChromosome, sizePopulation)

	fmt.Println("Worst: ", population[0].fitness, " Best: ", population[sizePopulation-1].fitness)

	for generation := 0; generation < numGeneration; generation++ {
		var parents []typeChromosome = selection(selectionName, population)
		if (len(parents) == 0) || (parents == nil) {
			fmt.Println(selectionName, " is not valid option.")
			return
		}

		var offspring typeChromosome = crossover(crossoverName, parents)
		if offspring.size == -1 {
			fmt.Println(crossoverName, " is not valid option.")
			return
		}

		offspring = mutation(mutationName, offspring, 0.0001)
		if offspring.size == -1 {
			fmt.Println(mutationName, " is not valid option.")
			return
		}

		population = replacement(replacementName, population, offspring)
		if (len(population) == 0) || (population == nil) {
			fmt.Println(replacementName, " is not valid option.")
			return
		}

		if population[0].fitness == typeFitness(sizeChromosome) {
			break
		}
	}

	fmt.Println("Worst: ", population[0].fitness, " Best: ", population[sizePopulation-1].fitness)
}
