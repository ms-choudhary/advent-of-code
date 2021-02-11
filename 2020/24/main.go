package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type HexCord struct {
	ew, sene, swnw int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	tilesFlipCount := map[string]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		i := 0
		var x, y float64
		for i < len(line) {
			switch line[i] {
			case 'e':
				x++
				i++

			case 'w':
				x--
				i++

			case 's':
				if line[i:i+2] == "se" {
					x += 0.5
					y -= 0.5
				} else if line[i:i+2] == "sw" {
					x -= 0.5
					y -= 0.5
				} else {
					log.Fatal("could not parse: s")
				}
				i += 2

			case 'n':
				if line[i:i+2] == "ne" {
					x += 0.5
					y += 0.5
				} else if line[i:i+2] == "nw" {
					x -= 0.5
					y += 0.5
				} else {
					log.Fatal("could not parse: n")
				}
				i += 2
			}
		}

		tilesFlipCount[key(x, y)]++
	}

	blackTiles := 0
	for k, v := range tilesFlipCount {
		if v%2 != 0 {
			blackTiles++
		}
		fmt.Printf("%v => %d\n", k, v)
	}

	fmt.Println(blackTiles)
}

func key(x, y float64) string {
	return strconv.FormatFloat(x, 'f', -1, 32) + ":" + strconv.FormatFloat(y, 'f', -1, 32)
}
