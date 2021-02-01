package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var ingredientToAlergen = map[string]string{}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	ingredientCount := map[string]int{}
	alergenToFood := map[string][][]string{}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, "(")

		ings := strings.Split(fields[0][:len(fields[0])-1], " ")
		alergs := strings.Split(fields[1][9:len(fields[1])-1], ", ")

		for _, v := range ings {
			ingredientCount[v]++
		}

		for _, v := range alergs {
			alergenToFood[v] = append(alergenToFood[v], ings)
		}
	}

	for alerg, foods := range alergenToFood {
		ings := mapIngredient(alerg, foods)
		for _, i := range ings {
			ingredientToAlergen[i] = alerg
		}
	}

	fmt.Println(ingredientToAlergen)

	count := 0
	for ing, c := range ingredientCount {
		if _, ok := ingredientToAlergen[ing]; !ok {
			count += c
		}
	}

	fmt.Println(count)
}

func mapIngredient(alergen string, foods [][]string) []string {
	ingcount := map[string]int{}
	for _, f := range foods {
		for _, ing := range f {
			ingcount[ing]++
		}
	}

	res := []string{}
	for k, v := range ingcount {
		if _, ok := ingredientToAlergen[k]; !ok && v == len(foods) {
			res = append(res, k)
		}
	}

	return res
}
