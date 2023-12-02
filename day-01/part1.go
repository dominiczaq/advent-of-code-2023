package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	fmt.Println("part1")
	file, err := os.Open("input-part1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	re := regexp.MustCompile("[0-9]+?")

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		foundStrings := re.FindAllString(line, -1)
		first := foundStrings[0]
		last := foundStrings[len(foundStrings)-1]

		// fmt.Println(line, first, last)
		combination := first + last
		numCombination, err := strconv.Atoi(combination)
		if err != nil {
			panic(err)
		}
		sum += numCombination
	}

	fmt.Println("Total sum:", sum)

}
