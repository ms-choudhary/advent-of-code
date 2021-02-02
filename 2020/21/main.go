package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var ingredientToAlergen = map[string]string{}
var alergenToIngredient = map[string]string{}

type AlergenFoodMapping struct {
	alergen string
	foods   [][]string
}

func main() {

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	ingredientCount := map[string]int{}
	alergenMapping := []AlergenFoodMapping{}
	alergenIndex := map[string]int{}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, "(")

		ings := strings.Split(fields[0][:len(fields[0])-1], " ")
		alergs := strings.Split(fields[1][9:len(fields[1])-1], ", ")

		for _, v := range ings {
			ingredientCount[v]++
		}

		for _, v := range alergs {
			if i, ok := alergenIndex[v]; ok {
				alergenMapping[i].foods = append(alergenMapping[i].foods, ings)
			} else {
				alergenMapping = append(alergenMapping, AlergenFoodMapping{v, [][]string{ings}})
				alergenIndex[v] = len(alergenMapping) - 1
			}
		}
	}

	for {
		cont := false
		for _, v := range alergenMapping {
			if _, ok := alergenToIngredient[v.alergen]; ok {
				continue
			}
			ings := mapIngredient(v.alergen, v.foods)

			if len(ings) > 1 {
				cont = true
			} else {
				ingredientToAlergen[ings[0]] = v.alergen
				alergenToIngredient[v.alergen] = ings[0]
			}
		}

		if cont == false {
			break
		}
	}

	mapping := []struct{ ingredient, alergen string }{}
	for ing, alerg := range ingredientToAlergen {
		mapping = append(mapping, struct{ ingredient, alergen string }{ingredient: ing, alergen: alerg})
	}

	sort.Slice(mapping, func(i, j int) bool {
		return mapping[i].alergen < mapping[j].alergen
	})

	fmt.Println(mapping)

	i := 0
	for ; i < len(mapping)-1; i++ {
		fmt.Print(mapping[i].ingredient + ",")
	}
	fmt.Println(mapping[i].ingredient)

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

	//fmt.Printf("alergen: %s - ing: %v\n", alergen, res)
	return res
}
