package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type InsType string

const (
	ACC InsType = "acc"
	JMP         = "jmp"
	NOP         = "nop"
)

type Instruction struct {
	operator InsType
	operand  int
	executed bool
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	instructions := []Instruction{}
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")

		inst := fields[0]
		var val int
		if fields[1][0] == '+' {
			val, _ = strconv.Atoi(fields[1][1:])
		} else {
			val, _ = strconv.Atoi(fields[1])
		}

		instructions = append(instructions, Instruction{operator: InsType(inst), operand: val})
	}

	//for _, v := range instructions {
	//fmt.Printf("%v : %v\n", v.operator, v.operand)
	//}

	fmt.Println(runInstructions(instructions))
}

func runInstructions(instruction []Instruction) int {
	pc := 0
	acc := 0

	for {
		if instruction[pc].executed {
			return acc
		}

		instruction[pc].executed = true

		switch instruction[pc].operator {
		case ACC:
			acc += instruction[pc].operand
			pc++
		case JMP:
			pc += instruction[pc].operand
		default:
			pc++
		}
	}

	return -1
}
