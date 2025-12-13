package day0402

import (
	"log"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0402 struct{
}

func NewDay0402() Day0402 {
	return Day0402{}
}

func (d Day0402)Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day04/01/input")
	numberOfRollsAccessible := rollsAccessible(input)

	log.Printf("The result is: %d", numberOfRollsAccessible)
}

func rollsAccessible(input []string) int {
	numberOfRollsAccessible := 0
	numberOfRollsRemoved := 1

	for numberOfRollsRemoved != 0 {
		indexToRemove := make([][]int, 0)

		for i, v1 := range(input) {
			for j, v2 := range(v1) {
				if v2 == '.' {
					continue
				}

				numberOfRolls := 0
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if y == 0 && x == 0 {
							continue
						}

						if fieldContainsPaper(input, j + x, i + y) {
							numberOfRolls++
						}
					}
				}

				if numberOfRolls < 4 {
					numberOfRollsAccessible++
					indexToRemove = append(indexToRemove, []int{j, i})
				}
			}
		}

		numberOfRollsRemoved = len(indexToRemove)
		input = removeRoll(input, indexToRemove)
	}

	return numberOfRollsAccessible
}

func removeRoll(input []string, toRemove [][]int) []string {
	for _, v := range(toRemove){
		x := v[0]
		y := v[1]
 
		lineToTransform := input[y]
		transformedLine := lineToTransform 
		if x == 0 {
			transformedLine = "." + lineToTransform[1:]
		} else if x >= len(lineToTransform) - 1 {
			transformedLine = lineToTransform[:len(lineToTransform) - 1] + "."
		} else {
			transformedLine = lineToTransform[:x] + "." + lineToTransform[x + 1:]
		}

		input[y] = transformedLine
	}

	return input
}

func fieldContainsPaper(input []string, x int, y int) bool {
	if y >= len(input) || y < 0{
		return false
	}

	if x >= len(input[y]) || x < 0{
		return false
	}

	return input[y][x] == '@'
}
