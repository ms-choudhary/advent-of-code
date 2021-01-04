package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var langRules map[string][]string

func main() {
	langRules = readRules()

	//fmt.Println(generateStrings("0"))

	s42s := generateStrings("42")
	s31s := generateStrings("31")

	index := map[string]bool{}
	for _, s := range generateStrings("0") {
		index[s] = true
	}

	matches := map[string]bool{}
	for _, s := range readInput() {
		for _, s42 := range s42s {
			for _, s31 := range s31s {
				news := removeLoops(s42, s31, s)
				//if s != news {
				//fmt.Printf("%s %s\n", news, s)
				//}
				if _, ok := index[news]; ok {
					matches[s] = true
				}
			}
		}
	}

	fmt.Println(matches)

	fmt.Println(len(matches))
}

func readInput() []string {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func readRules() map[string][]string {
	f, err := os.Open("rules")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	rules := map[string][]string{}
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ":")
		if strings.Contains(fields[1], "|") {
			values := []string{}
			for _, v := range strings.Split(fields[1][1:], "|") {
				values = append(values, strings.Trim(v, " "))
			}
			rules[fields[0]] = values
		} else {
			rules[fields[0]] = []string{fields[1][1:]}
		}
	}
	return rules
}

func generateStrings(root string) []string {
	def := langRules[root]

	res := []string{}
	for _, d := range def {
		if strings.Contains(d, "\"") {
			res = append(res, strings.Trim(d, "\""))
		} else {
			val := [][]string{}
			for _, r := range strings.Split(d, " ") {
				val = append(val, generateStrings(r))
			}
			res = append(res, combine(val)...)
		}
	}
	return res
}

func combine(vals [][]string) []string {
	if len(vals) == 1 {
		return vals[0]
	}

	res := []string{}
	for i := 0; i < len(vals[0]); i++ {
		for j := 0; j < len(vals[1]); j++ {
			res = append(res, vals[0][i]+vals[1][j])
		}
	}
	if len(vals) > 2 {
		return combine([][]string{res, vals[2]})
	}
	return res
}

func removeLoops(s42, s31, inp string) string {
	if !strings.Contains(inp, s42) {
		return inp
	}

	i := strings.Index(inp, s42) + len(s42)
	res := inp[:i]
	for ; i < len(inp); i += len(s42) {
		if i+len(s42) <= len(inp) && inp[i:i+len(s42)] != s42 {
			break
		}
	}

	if i+len(s31) <= len(inp) && inp[i:i+len(s31)] == s31 {
		res += s31
		for ; i < len(inp); i += len(s31) {
			if i+len(s31) <= len(inp) && inp[i:i+len(s31)] != s31 {
				break
			}
		}
	}

	if i < len(inp) {
		res += inp[i:]
	}
	return res
}
