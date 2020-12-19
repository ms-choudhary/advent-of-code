package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	arvTs := 0
	if scanner.Scan() {
		line := scanner.Text()
		arvTs, _ = strconv.Atoi(line)
	}

	busIDs := []int{}
	if scanner.Scan() {
		line := scanner.Text()
		buses := strings.Split(line, ",")

		for _, b := range buses {
			if b == "x" {
				continue
			}
			id, _ := strconv.Atoi(b)
			busIDs = append(busIDs, id)
		}
	}

	id, time := earliestBus(busIDs, arvTs)
	fmt.Println(id * time)
}

func earliestBus(busIDs []int, arvTs int) (int, int) {
	minDiff := math.MaxInt32
	minBusId := 0

	for _, b := range busIDs {
		if arvTs%b == 0 {
			minDiff = 0
			minBusId = b
		} else if (b*(arvTs/b)+b)-arvTs < minDiff {
			minDiff = (b*(arvTs/b) + b) - arvTs
			minBusId = b
		}
	}
	return minBusId, minDiff
}
