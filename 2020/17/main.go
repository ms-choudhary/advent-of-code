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
	slice := [][]rune{}
	for scanner.Scan() {
		slice = append(slice, []rune(scanner.Text()))
	}

	cubes := make([][][][]rune, 1)
	cubes[0] = make([][][]rune, 1)
	cubes[0][0] = slice

	for i := 0; i < 6; i++ {
		ws := [][][][]rune{}
		for w := -1; w < len(cubes)+1; w++ {
			zs := [][][]rune{}
			for z := -1; z < len(cubes[0])+1; z++ {
				ys := [][]rune{}
				for y := -1; y < len(cubes[0][0])+1; y++ {
					xs := []rune{}
					for x := -1; x < len(cubes[0][0][0])+1; x++ {
						n := findActNeighbour(x, y, z, w, cubes)
						ch := '.'
						if at(x, y, z, w, cubes) == '#' && (n == 2 || n == 3) {
							ch = '#'
						} else if at(x, y, z, w, cubes) == '.' && n == 3 {
							ch = '#'
						}
						xs = append(xs, ch)
					}
					ys = append(ys, xs)
				}
				zs = append(zs, ys)
			}
			ws = append(ws, zs)
		}
		cubes = ws
	}

	totAct := 0
	for w := range cubes {
		for z := range cubes[w] {
			for y := range cubes[w][z] {
				for x := range cubes[w][z][y] {
					if cubes[w][z][y][x] == '#' {
						totAct++
					}
				}
			}
		}
	}
	fmt.Println(totAct)
}

func printCubes(cubes [][][]rune) {
	for z := range cubes {
		for y := range cubes[z] {
			fmt.Println(string(cubes[z][y]))
		}
		fmt.Println()
	}
}

func findActNeighbour(x, y, z, w int, cubes [][][][]rune) int {
	active := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				for l := w - 1; l <= w+1; l++ {
					if i == x && j == y && k == z && l == w {
						continue
					}
					if at(i, j, k, l, cubes) == '#' {
						active++
					}
				}
			}
		}
	}
	return active
}

func at(x, y, z, w int, cubes [][][][]rune) rune {
	if x < 0 || x >= len(cubes[0][0][0]) {
		return '.'
	}

	if y < 0 || y >= len(cubes[0][0]) {
		return '.'
	}

	if z < 0 || z >= len(cubes[0]) {
		return '.'
	}

	if w < 0 || w >= len(cubes) {
		return '.'
	}
	return cubes[w][z][y][x]
}
