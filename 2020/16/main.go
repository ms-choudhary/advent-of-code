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

	ranges := map[string][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var r1low, r1high int
		var r2low, r2high int

		fields := strings.Split(line, ":")
		n, err := fmt.Sscanf(fields[1][1:], "%d-%d or %d-%d", &r1low, &r1high, &r2low, &r2high)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("couldn't parse complete line")
		}

		ranges[fields[0]] = []int{r1low, r1high, r2low, r2high}
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

	tickets := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "nearby tickets:" {
			continue
		}
		tickets = append(tickets, parseTicket(line))
	}

	validTickets := [][]int{}
	for _, t := range tickets {
		if validTicket(t, ranges) {
			validTickets = append(validTickets, t)
		}
	}

	tempInd := make([]map[string]int, len(validTickets[0]))
	for i := range tempInd {
		tempInd[i] = map[string]int{}
	}

	for _, t := range validTickets {
		for i, f := range t {
			for k, r := range ranges {
				if (f >= r[0] && f <= r[1]) || (f >= r[2] && f <= r[3]) {
					tempInd[i][k]++
				}
			}
		}
	}

	index := make([]string, len(tempInd))

	remInd := []int{}
	for i := 0; i < len(tempInd); i++ {
		remInd = append(remInd, i)
	}

	for {
		newRemInd := []int{}
		for _, i := range remInd {
			interm := []string{}
			for k, v := range tempInd[i] {
				if v == len(validTickets) && !contains(index, k) {
					interm = append(interm, k)
				}
			}

			if len(interm) == 1 {
				index[i] = interm[0]
			} else {
				newRemInd = append(newRemInd, i)
			}
		}
		if len(newRemInd) == 0 {
			break
		}
		remInd = newRemInd
	}

	fmt.Println(strings.Join(index, ","))

	res := 1
	for i, s := range index {
		if strings.HasPrefix(s, "departure") {
			res *= myticket[i]
		}
	}

	fmt.Println(res)
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func validTicket(t []int, ranges map[string][]int) bool {
	for _, f := range t {
		found := false
		for _, r := range ranges {
			if (f >= r[0] && f <= r[1]) || (f >= r[2] && f <= r[3]) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
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
