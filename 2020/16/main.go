package main

import (
	"bufio"
	"fmt"
	"log"
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

	ranges := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var r1low, r1high int
		var r2low, r2high int

		n, err := fmt.Sscanf(strings.Split(line, ":")[1][1:], "%d-%d or %d-%d", &r1low, &r1high, &r2low, &r2high)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("couldn't parse complete line")
		}

		ranges = append(ranges, []int{r1low, r1high})
		ranges = append(ranges, []int{r2low, r2high})
	}

	myticket := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else if line == "your ticket:" {
			continue
		}
		myticket = parseTicket(line)
	}

	fmt.Println(myticket)

	tickets := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "nearby tickets:" {
			continue
		}
		tickets = append(tickets, parseTicket(line))
	}

	invalidSum := 0
	for _, t := range tickets {
		for _, f := range t {
			found := false
			for _, r := range ranges {
				if f >= r[0] && f <= r[1] {
					found = true
				}
			}
			if !found {
				fmt.Println(f)
				invalidSum += f
			}
		}
	}

	fmt.Println(invalidSum)
}

func parseTicket(s string) []int {
	fields := strings.Split(s, ",")
	ticket := []int{}
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		ticket = append(ticket, n)
	}
	return ticket
}
