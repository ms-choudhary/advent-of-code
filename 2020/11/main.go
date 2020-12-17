package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	seats := [][]rune{}
	for scanner.Scan() {
		seats = append(seats, []rune(scanner.Text()))
	}

	changed := false
	for {
		changed, seats = fillSeats(seats)
		if !changed {
			break
		}

		changed, seats = emptySeats(seats)
		if !changed {
			break
		}
	}

	occupied := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' {
				occupied++
			}
		}
	}

	fmt.Println(occupied)
}

func emptySeats(seats [][]rune) (bool, [][]rune) {
	newseats := [][]rune{}
	changed := false
	for i := 0; i < len(seats); i++ {
		row := []rune{}
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' && adjSeatsOccupied(i, j, seats) >= 4 {
				changed = true
				row = append(row, 'L')
			} else {
				row = append(row, seats[i][j])
			}
		}
		newseats = append(newseats, row)
	}
	return changed, newseats
}

func adjSeatsOccupied(i, j int, seats [][]rune) int {
	res := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x == i && y == j {
				continue
			}

			if x >= 0 && x < len(seats) && y >= 0 && y < len(seats[0]) {
				if seats[x][y] == '#' {
					res++
				}
			}
		}
	}
	return res
}

func fillSeats(seats [][]rune) (bool, [][]rune) {
	newseats := [][]rune{}
	changed := false
	for i := 0; i < len(seats); i++ {
		row := []rune{}
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == 'L' && adjSeatsOccupied(i, j, seats) == 0 {
				changed = true
				row = append(row, '#')
			} else {
				row = append(row, seats[i][j])
			}
		}
		newseats = append(newseats, row)
	}
	return changed, newseats
}
