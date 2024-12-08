package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	Value    int
	Operands []int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	input := strings.TrimSpace(string(bytes))
	input = strings.ReplaceAll(input, "\r\n", "\n")

	eqs := parseEquations(input)

	part1Result := 0
	part2Result := 0
	for _, eq := range eqs {
		part1Result += getPart1Result(eq.Value, eq.Operands)
		part2Result += getPart2Result(eq.Value, eq.Operands)
	}

	fmt.Printf("The total of all the correct part 1 results is: %d\n", part1Result)
	fmt.Printf("The total of all the correct part 2 results is: %d\n", part2Result)
}

func getPart1Result(target int, operands []int) int {
	return getResultIfEqual(target, operands, false)
}

func getPart2Result(target int, operands []int) int {
	return getResultIfEqual(target, operands, true)
}

func getResultIfEqual(target int, operands []int, concat bool) int {
	if len(operands) < 1 {
		return 0
	}

	if len(operands) == 1 {
		if operands[0] == target {
			return target
		}
		return 0
	}

	addOps := []int{operands[0] + operands[1]}
	if len(operands) > 2 {
		addOps = append(addOps, operands[2:]...)
	}

	res := getResultIfEqual(target, addOps, concat)
	if res != 0 {
		return res
	}

	mulOps := []int{operands[0] * operands[1]}
	if len(operands) > 2 {
		mulOps = append(mulOps, operands[2:]...)
	}

	res = getResultIfEqual(target, mulOps, concat)
	if res != 0 || !concat {
		return res
	}

	op1Str := strconv.Itoa(operands[0])
	op2Str := strconv.Itoa(operands[1])
	opConc, err := strconv.Atoi(op1Str + op2Str)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	concOps := []int{opConc}
	if len(operands) > 2 {
		concOps = append(concOps, operands[2:]...)
	}

	return getResultIfEqual(target, concOps, concat)
}

func parseEquations(input string) []equation {
	lines := strings.Split(input, "\n")

	eqs := make([]equation, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		valString, s := parts[0], parts[1]
		opsString := strings.Split(s, " ")

		val, err := strconv.Atoi(valString)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		ops := make([]int, 0, len(opsString))
		for _, opS := range opsString {
			op, err := strconv.Atoi(opS)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			ops = append(ops, op)
		}

		eqs = append(eqs, equation{Value: val, Operands: ops})
	}

	return eqs
}
