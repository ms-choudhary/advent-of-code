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

func printPos(pos map[Action]int) {
	fmt.Printf("N: %v, S: %v, E: %v, W: %v\n", pos[North], pos[South], pos[East], pos[West])
}

func calcInstructions(instructions []Nav) (int, int) {
	shippos := map[Action]int{
		East:  0,
		North: 0,
		South: 0,
		West:  0,
	}

	wayppos := map[Action]int{
		East:  10,
		North: 1,
		South: 0,
		West:  0,
	}
	for _, i := range instructions {
		switch i.action {
		case Forward:
			shippos[North] += i.value * wayppos[North]
			shippos[South] += i.value * wayppos[South]
			shippos[East] += i.value * wayppos[East]
			shippos[West] += i.value * wayppos[West]
		case North:
			wayppos[North] += i.value
		case South:
			wayppos[South] += i.value
		case East:
			wayppos[East] += i.value
		case West:
			wayppos[West] += i.value
		case Left:
			leftCircle := []Action{North, West, South, East}
			nQuarters := i.value / 90

			newwayppos := map[Action]int{}

			newwayppos[leftCircle[nQuarters%4]] = wayppos[North]
			newwayppos[leftCircle[(1+nQuarters)%4]] = wayppos[West]
			newwayppos[leftCircle[(2+nQuarters)%4]] = wayppos[South]
			newwayppos[leftCircle[(3+nQuarters)%4]] = wayppos[East]

			wayppos = newwayppos
		case Right:
			rightCircle := []Action{North, East, South, West}
			nQuarters := i.value / 90

			newwayppos := map[Action]int{}

			newwayppos[rightCircle[nQuarters%4]] = wayppos[North]
			newwayppos[rightCircle[(3+nQuarters)%4]] = wayppos[West]
			newwayppos[rightCircle[(2+nQuarters)%4]] = wayppos[South]
			newwayppos[rightCircle[(1+nQuarters)%4]] = wayppos[East]

			wayppos = newwayppos
		}

		//printPos(shippos)
		//printPos(wayppos)
		//fmt.Println()
	}

	ew, ns := 0, 0
	if shippos[East] > shippos[West] {
		ew = shippos[East] - shippos[West]
	} else {
		ew = shippos[West] - shippos[East]
	}

	if shippos[North] > shippos[South] {
		ns = shippos[North] - shippos[South]
	} else {
		ns = shippos[South] - shippos[North]
	}

	return ew, ns
}
