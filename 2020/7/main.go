package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	contains []*Bag
	count    []int
	partof   []*Bag
}

var bagIndex = map[string]*Bag{}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		addBag(parseBag(scanner.Text()))
	}

	fmt.Println(countAllBagsInside("shiny gold"))
}

func parseBag(s string) (string, []string, []int) {
	fields := strings.Split(s, "contain")

	var typ, color string
	fmt.Sscanf(fields[0], "%s %s bags", &typ, &color)

	if fields[1] == "no other bags." {
		return typ + " " + color, []string{}, []int{}
	}
	containStr := fields[1][:len(fields[1])-1]

	contains := []string{}
	count := []int{}
	for _, v := range strings.Split(containStr, ",") {
		var t, c string
		var i int
		if strings.HasSuffix(v, "bags") {
			fmt.Sscanf(v, "%d %s %s bags", &i, &t, &c)
		} else if strings.HasSuffix(v, "bag") {
			fmt.Sscanf(v, "%d %s %s bag", &i, &t, &c)
		}

		count = append(count, i)
		contains = append(contains, t+" "+c)
	}
	return typ + " " + color, contains, count
}

func printBag(b *Bag) {
	contains := ""
	for i, v := range b.contains {
		contains += strconv.Itoa(b.count[i]) + "." + v.color + ", "
	}

	partof := ""
	for _, v := range b.partof {
		partof += v.color + ", "
	}

	fmt.Printf("c: %s| cnts: %s| prtof: %s\n", b.color, contains, partof)
}

func addBag(color string, contains []string, count []int) *Bag {

	var bag *Bag
	bag, ok := bagIndex[color]
	if !ok {
		bag = &Bag{color: color}
		bagIndex[color] = bag
	}

	containBags := []*Bag{}
	bagCounts := []int{}
	for i, c := range contains {
		if b, ok := bagIndex[c]; ok {
			b.partof = append(b.partof, bag)
			containBags = append(containBags, b)
		} else {
			b = &Bag{color: c}
			b.partof = append(b.partof, bag)
			bagIndex[c] = b
			containBags = append(containBags, b)
		}

		bagCounts = append(bagCounts, count[i])
	}

	bag.contains = containBags
	bag.count = bagCounts

	return bag
}

func countAllBagsInside(color string) int {
	bag := bagIndex[color]

	count := 0
	for i, v := range bag.contains {
		count += bag.count[i] + bag.count[i]*countAllBagsInside(v.color)
	}
	return count
}

func findUniqueParents(color string) map[string]bool {
	bag := bagIndex[color]

	parents := map[string]bool{}
	for _, b := range bag.partof {
		parents[b.color] = true
		m := findUniqueParents(b.color)

		for k, _ := range m {
			parents[k] = true
		}
	}

	return parents
}
