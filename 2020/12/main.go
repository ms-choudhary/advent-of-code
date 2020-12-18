package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Action rune

const (
	Forward Action = 'F'
	North   Action = 'N'
	South   Action = 'S'
	East    Action = 'E'
	West    Action = 'W'
	Left    Action = 'L'
	Right   Action = 'R'
)

type Nav struct {
	action Action
	value  int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	instructions := []Nav{}
	for scanner.Scan() {
		line := scanner.Text()

		if val, err := strconv.Atoi(line[1:]); err == nil {
			instructions = append(instructions, Nav{action: Action(line[0]), value: val})
		}
	}

	east, north := calcInstructions(instructions)
	fmt.Printf("%v\n", east+north)
}

func calcInstructions(instructions []Nav) (int, int) {
	pos := map[Action]int{
		East:  0,
		North: 0,
		South: 0,
		West:  0,
	}

	curdir := East
	for _, i := range instructions {
		switch i.action {
		case Forward:
			pos[curdir] += i.value
		case North:
			pos[North] += i.value
		case South:
			pos[South] += i.value
		case East:
			pos[East] += i.value
		case West:
			pos[West] += i.value
		case Left:
			leftCircle := []Action{North, West, South, East}
			leftIndex := map[Action]int{North: 0, West: 1, South: 2, East: 3}
			nQuarters := i.value / 90
			curdir = leftCircle[(leftIndex[curdir]+nQuarters)%4]
		case Right:
			rightCircle := []Action{North, East, South, West}
			rightIndex := map[Action]int{North: 0, East: 1, South: 2, West: 3}
			nQuarters := i.value / 90
			curdir = rightCircle[(rightIndex[curdir]+nQuarters)%4]
		}
	}

	ew, ns := 0, 0
	if pos[East] > pos[West] {
		ew = pos[East] - pos[West]
	} else {
		ew = pos[West] - pos[East]
	}

	if pos[North] > pos[South] {
		ns = pos[North] - pos[South]
	} else {
		ns = pos[South] - pos[North]
	}

	return ew, ns
}
