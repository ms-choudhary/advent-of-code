package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Bus struct {
	id, queue int
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	buses := []Bus{}
	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		for i, b := range fields {
			if b == "x" {
				continue
			}
			id, _ := strconv.Atoi(b)
			buses = append(buses, Bus{id: id, queue: i})
		}
	}

	sort.Slice(buses, func(i, j int) bool {
		return buses[i].id > buses[j].id
	})

	fmt.Println(earliestTs(buses))
}

func earliestTs(buses []Bus) int {
	// 99999999999558
	for ts := 99999999999558; ; ts += buses[0].id {
		//fmt.Println(ts)
		satisfies := true
		for _, b := range buses {
			if (ts+b.queue-buses[0].queue)%b.id != 0 {
				satisfies = false
			}
		}

		if satisfies {
			return ts - buses[0].queue
		}
	}
}
