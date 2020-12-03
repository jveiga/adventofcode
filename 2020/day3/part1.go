package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Coord struct{ x, y int }

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("failed to read input", err)
	}
	lines := strings.Split(string(data), "\n")

	position := Coord{3, 1}
	numberOfTrees := 0
	log.Println(lines[0])
	for _, line := range lines[1:] {
		length := len(line)
		if len(line) == 0 {
			continue
		}
		x := position.x % (length)
		location := line[x]
		if location == '#' {
			numberOfTrees++
		}
		position.x += 3
		// position.y++
	}
	fmt.Println(numberOfTrees)
}
