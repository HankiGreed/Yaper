package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	nameArray, err := ioutil.ReadFile("./p022_names.txt")
	if err != nil {
		log.Fatalln(err)
	}

	nameString := strings.ReplaceAll(string(nameArray), "\"", "")
	names := strings.Split(nameString, ",")
	sort.Strings(names)

	totalScore := 0
	for i, name := range names {
		totalScore += (calcNameScore(name) * (i + 1))
	}
	fmt.Println(totalScore)
}

func calcNameScore(name string) int {
	name = strings.ToUpper(name)

	score := 0
	for _, c := range name {
		score += (int(c) - int('A') + 1)
	}
	return score
}
