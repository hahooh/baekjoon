package main

import (
	"fmt"
)

type position struct {
	x int
	y int
}

func main() {
	var numTest int
	var results []int

	fmt.Scanf("%d", &numTest)

	var preStopDists []int
	preStopDists = append(preStopDists, 0)
	var x int
	var y int
	for i := 0; i < numTest; i++ {
		fmt.Scanf("%d", &x)
		fmt.Scanf("%d", &y)
		distance := y - x
		steps := 0
		currentSpeed := 0
		for distance != 0 {
			currentSpeed = getNextSpeed(distance, currentSpeed, &preStopDists)
			distance -= currentSpeed
			steps++
		}
		results = append(results, steps)
	}

	for _, step := range results {
		fmt.Println(step)
	}
}

func getStopDist(speed int, prevStopDist int) int {
	return speed + prevStopDist
}

func getNextSpeed(distance int, currentSpeed int, preStopDists *[]int) int {
	acceleration := currentSpeed + 1
	accStopDist := getStopDist(acceleration, (*preStopDists)[currentSpeed])
	if accStopDist <= distance {
		if len(*preStopDists) <= acceleration {
			*preStopDists = append(*preStopDists, accStopDist)
		}
		return acceleration
	}

	currStopDist := (*preStopDists)[currentSpeed]
	if currStopDist <= distance {
		return currentSpeed
	}

	return currentSpeed - 1
}
