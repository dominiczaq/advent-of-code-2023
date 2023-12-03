package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	fmt.Println("day2 part2")
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
		if err != nil {
			panic(err)
		}
		records = strings.TrimSpace(records)
		// fmt.Println(game, records)
		individualRecors := strings.Split(records, ";")
		// fmt.Println(game)
		gameMaxRecord := Record{
			Red:   0,
			Green: 0,
			Blue:  0,
		}
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
			if recordStruct.Blue > gameMaxRecord.Blue {
				gameMaxRecord.Blue = recordStruct.Blue
			}
			if recordStruct.Green > gameMaxRecord.Green {
				gameMaxRecord.Green = recordStruct.Green
			}
			if recordStruct.Red > gameMaxRecord.Red {
				gameMaxRecord.Red = recordStruct.Red
			}
		}
		power := gameMaxRecord.Blue * gameMaxRecord.Green * gameMaxRecord.Red
		sum += power
	}

	println("Day2 part 2: Sum of powers", sum)

}
