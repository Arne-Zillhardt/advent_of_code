package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Arne-Zillhardt/advent_of_code/2025/day01/01"
	"github.com/Arne-Zillhardt/advent_of_code/2025/fileInput"
)

type Puzzle interface {
	Execute(fileInput.FileInput)
}

var puzzles = map[string]Puzzle {
	"day0101": day0101.NewDay0101(),
}

func main() {
	fmt.Print("Enter key: ")
	reader := bufio.NewReader(os.Stdin)

	key, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}

	key = key[:len(key)-1]
	puzzles[key].Execute(fileInput.NewFileInput())
}
