package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  int = 20
	height int = 20
)

type organism struct {
	neighbors [8]*organism
	val       bool
}

func main() {
	population := initPopulation()
	sem := make(chan bool)

	printPopulation(population) // print first time manually to avoid initial delay

	for {
		go func() {
			agePopulation(population)
			sem <- true
		}()

		time.Sleep(time.Millisecond * 100)

		<-sem
		printPopulation(population)
	}
}

func agePopulation(population [width][height]*organism) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			ageOrganism(population[i][j])
		}
	}
}

func printPopulation(population [width][height]*organism) {
	fmt.Print("\033[H\033[2J") // clean screen

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := " "
			if population[i][j].val {
				c = "ˆ"
			}

			fmt.Printf(" %s ", c)
		}

		fmt.Println("")
	}
}

func initPopulation() [width][height]*organism {
	var population [width][height]*organism

	// Add values
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			population[i][j] = &organism{
				val: rand.Int()%2 == 0,
			}
		}
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			population[i][j].neighbors = [8]*organism{
				getNeighbor(population, i-1, j-1),
				getNeighbor(population, i-1, j),
				getNeighbor(population, i-1, j+1),
				getNeighbor(population, i, j-1),
				getNeighbor(population, i, j+1),
				getNeighbor(population, i+1, j-1),
				getNeighbor(population, i+1, j),
				getNeighbor(population, i+1, j+1),
			}
		}
	}

	return population
}

func getNeighbor(population [width][height]*organism, i int, j int) *organism {
	if i < 0 || j < 0 || i >= width || j >= height {
		return nil
	}

	return population[i][j]
}

func ageOrganism(o *organism) {
	neighborCount := countNeighbors(o)

	if neighborCount < 2 || neighborCount > 3 {
		o.val = false
		return
	}

	if neighborCount == 3 {
		o.val = true
		return
	}
}

func countNeighbors(o *organism) int {
	acc := 0
	for _, n := range o.neighbors {
		if n == nil {
			continue
		}

		if n.val {
			acc++
		}
	}

	return acc
}
