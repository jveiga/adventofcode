package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const SUMTOTAL = 2020

type Expense struct {
	a, b int
}

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("failed to open input file %w", err)
	}
	bts := bytes.NewBuffer(f)
	var (
		nums     []int
		expenses []Expense
		seen     map[Expense]struct{}
	)
	read := bufio.NewReader(bts)
	for {
		line, _, err := read.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("wtf was that ", err)
		}
		if string(line) == "" {
			continue
		}
		num, err := strconv.Atoi(strings.ReplaceAll(string(line), "\n", ""))
		if err != nil {
			log.Fatal("wtf was that number ", err)
		}
		nums = append(nums, num)
	}
	seen = make(map[Expense]struct{})
	for _, i := range nums {
		for _, j := range nums {
			if i+j == SUMTOTAL {
				exp := Expense{i, j}
				if _, ok := seen[exp]; !ok {
					expenses = append(expenses, exp)
					seen[exp] = struct{}{}
					fmt.Println(exp)
				}
			}
			for _, k := range nums {
				if i+j+k == SUMTOTAL {
					fmt.Println(i * j * k)
				}
			}
		}
	}
	// for
}
