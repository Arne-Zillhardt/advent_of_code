package day0401

import (
	"log"

	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Day0401 struct{
}

func NewDay0401() Day0401 {
	return Day0401{}
}

func (d Day0401)Execute(fileInput fileInput.FileInput) {
	input := fileInput.GetContentOfFile("./day04/01/input")
	numberOfRollsAccessible := rollsAccessible(input)

	log.Printf("The result is: %d", numberOfRollsAccessible)
}

func rollsAccessible(input []string) int {
	numberOfRollsAccessible := 0

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
			}
		}
	}

	return numberOfRollsAccessible
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
