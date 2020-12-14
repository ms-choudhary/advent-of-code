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
	pc, acc  int
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

func runInstructions(prog []Instruction) int {
	pc := 0
	acc := 0
	stack := []*Instruction{}

	skipstore := false
	for pc < len(prog) {

		if prog[pc].executed {
			lastinst := stack[len(stack)-1]

			pc, acc = lastinst.pc, lastinst.acc

			if lastinst.operator == JMP {
				lastinst.operator = NOP
			} else if lastinst.operator == NOP {
				lastinst.operator = JMP
			}
			lastinst.executed = false

			stack = stack[:len(stack)-1]

			skipstore = true
		}

		prog[pc].executed = true

		//fmt.Printf("%v : %v\n", prog[pc].operator, prog[pc].operand)

		switch prog[pc].operator {
		case ACC:
			acc += prog[pc].operand
			pc++

		case JMP:
			if !skipstore {
				prog[pc].acc = acc
				prog[pc].pc = pc
				stack = append(stack, &prog[pc])
				skipstore = false
			}
			pc += prog[pc].operand

		case NOP:
			if !skipstore {
				prog[pc].acc = acc
				prog[pc].pc = pc
				stack = append(stack, &prog[pc])
				skipstore = false
			}
			pc++

		default:
			pc++
		}
	}

	return acc
}
