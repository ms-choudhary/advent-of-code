package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Tile struct {
	id                       int
	content                  []string
	sides                    []string
	left, right, top, bottom *Tile
}

var lookup = map[string][]int{}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	tiles := []Tile{}
	tile := []string{}
	var id int
	readTile := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			tiles = append(tiles, Tile{id: id, content: tile, sides: []string{tile[0], rightSide(tile), tile[len(tile)-1], leftSide(tile)}})
			tile = []string{}
			readTile = false
		} else if readTile {
			tile = append(tile, line)
		} else {
			if n, err := fmt.Sscanf(line, "Tile %d:", &id); err != nil || n != 1 {
				log.Fatalf("failed to parse line: %v", err)
			}
			readTile = true
		}
	}

	for _, t := range tiles {
		addSide(lookup, t.sides[0], t.id)
		addSide(lookup, t.sides[1], t.id)
		addSide(lookup, t.sides[2], t.id)
		addSide(lookup, t.sides[3], t.id)
	}

	//fmt.Println(findCorners(tiles))
	for _, t := range findCorners(tiles) {
		fmt.Println(t.id)
	}
}

//func stitchTiles(tiles []Tile, lookup map[string][]int) [][]Tile {

//}

func flipTopBottom(t *Tile) {
	temp := t.sides[0]
	t.sides[0] = t.sides[2]
	t.sides[2] = temp
}

func flipLeftRight(t *Tile) {
	temp := t.sides[1]
	t.sides[1] = t.sides[3]
	t.sides[3] = temp
}

func rotate(t *Tile, times int) {
	newSides := make([]string, 4)
	for i, s := range t.sides {
		newSides[(i+times)%4] = s
	}
	t.sides = newSides
}

func findCorners(tiles []Tile) []Tile {
	res := []Tile{}
	for _, t := range tiles {
		count := 0
		for _, s := range t.sides {
			if adj, ok := lookup[s]; ok && len(adj) > 1 {
				count++
			}
		}

		if count == 2 {
			fmt.Println(count)
			res = append(res, t)
		}
	}
	return res
}

func addSide(lookup map[string][]int, side string, id int) {
	if _, ok := lookup[side]; ok {
		lookup[side] = append(lookup[side], id)
	} else {
		lookup[side] = []int{id}
	}
}

func leftSide(tc []string) string {
	res := ""
	for _, s := range tc {
		res += s[:1]
	}
	return res
}

func rightSide(tc []string) string {
	res := ""
	for _, s := range tc {
		res += s[len(s)-1:]
	}
	return res
}
