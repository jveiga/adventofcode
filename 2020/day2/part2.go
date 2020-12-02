package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("failed to read input", err)
	}

	lines := strings.Split(string(data), "\n")
	validPassword := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, ":")
		if len(split) != 2 {
			log.Println("invalid line?", line)
			continue
		}
		rule := strings.Split(split[0], " ")
		if len(rule) != 2 {
			log.Fatal("invalid rule section", rule)
		}
		minPart := strings.Split(rule[0], "-")
		letter := rule[1]
		// minPart := rule[0]
		min, err := strconv.Atoi(minPart[0])
		if err != nil {
			log.Println("invalid minimum", minPart[0], err)
			continue
		}
		max, err := strconv.Atoi(minPart[1])
		if err != nil {
			log.Println("invalid minimum", minPart[1], err)
			continue
		}
		password := split[1]
		if (byte(password[min]) == byte(letter[0])) != (password[max] == byte(letter[0])) {
			validPassword++
		}
	}
	fmt.Println(validPassword)
}
