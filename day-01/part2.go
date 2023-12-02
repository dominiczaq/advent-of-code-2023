package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2(filename string) int {
	fmt.Println("part2")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = ReplaceAllStringNumericNumbers(line)
		foundStrings := GetAllStringAndNumericNumbersFromString(line)

		first := ParseNumberString(foundStrings[0])
		last := ParseNumberString(foundStrings[len(foundStrings)-1])
		combination := first + last
		numCombination, err := strconv.Atoi(combination)

		if err != nil {
			fmt.Println("combination", combination, line, "first", first, "last", last, foundStrings[0])
			panic(err)
		}
		// fmt.Println(line, foundStrings, first, last, combination, numCombination)

		sum += numCombination
	}

	fmt.Println("Total sum part2:", sum)
	return sum
}

func GetAllStringAndNumericNumbersFromString(someString string) []string {
	re := regexp.MustCompile("(([0-9])+?|(two)+?|(three)+?|(four)+?|(five)+?|(six)+?|(seven)+?|(eight)+?|(nine)+?|(one)+?)")

	return re.FindAllString(someString, -1)
}

func ReplaceAllStringNumericNumbers(someString string) string {
	re := regexp.MustCompile("((one)+?|(two)+?|(three)+?|(four)+?|(five)+?|(six)+?|(seven)+?|(eight)+?|(nine)+?)")
	result := re.ReplaceAllStringFunc(someString, func(match string) string {
		switch match {
		case "one":
			return "1e"
		case "two":
			return "2o"
		case "three":
			return "3e"
		case "four":
			return "4r"
		case "five":
			return "5e"
		case "six":
			return "6x"
		case "seven":
			return "7n"
		case "eight":
			return "8t"
		case "nine":
			return "9e"
		default:
			return match
		}
	})
	return result
}

func ParseNumberString(stringNumber string) string {
	switch stringNumber {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return stringNumber
	}
}
