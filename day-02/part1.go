package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Red   int
	Green int
	Blue  int
}

func Part1() {
	fmt.Println("day2 part1")
	file, err := os.Open("input-part1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		game, records, found := strings.Cut(line, ":")
		if !found {
			panic("Not found any recors")
		}
		game = strings.TrimSpace(game)
		_, gameIndexString, found := strings.Cut(game, " ")
		if !found {
			panic("game index not found")
		}
		gameIndex, err := strconv.Atoi(gameIndexString)
		if err != nil {
			panic(err)
		}
		records = strings.TrimSpace(records)
		// fmt.Println(game, records)
		individualRecors := strings.Split(records, ";")
		// fmt.Println(game)
		possible := true
		for _, record := range individualRecors {
			record = strings.TrimSpace(record)
			// fmt.Println(record)
			cubeColors := strings.Split(record, ",")
			recordStruct := new(Record)
			for _, cubeColor := range cubeColors {
				cubeColor = strings.TrimSpace(cubeColor)
				// fmt.Println(cubeColor)
				value, color, found := strings.Cut(cubeColor, " ")
				if !found {
					panic("values not found in color")
				}
				numValue, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}
				if strings.Contains(color, "red") {
					recordStruct.Red = numValue
				}
				if strings.Contains(color, "green") {
					recordStruct.Green = numValue
				}
				if strings.Contains(color, "blue") {
					recordStruct.Blue = numValue
				}
			}
			if recordStruct.Blue > 14 || recordStruct.Green > 13 || recordStruct.Red > 12 {
				possible = false
				break
			}
		}
		if possible {
			sum += gameIndex
		}
	}

	println("Day2 part 1: Sum of possible game indexes", sum)

}
