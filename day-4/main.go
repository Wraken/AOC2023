package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func readFile() []string {
	r, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	scan := bufio.NewScanner(r)

	scan.Split(bufio.ScanLines)

	lines := []string{}
	for scan.Scan() {
		l := scan.Text()
		lines = append(lines, l)
	}

	return lines
}

type card struct {
	id     int
	nbs    []int
	winNbs []int
}

func stringToInt(s string) []int {
	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, " ")
	ns := strings.Split(s, " ")
	nbs := make([]int, len(ns))
	for i, n := range ns {
		n, _ := strconv.Atoi(n)
		nbs[i] = n
	}
	return nbs
}

func parse(lines []string) []*card {
	cards := make([]*card, 0, len(lines))

	for i, l := range lines {
		c := &card{id: i}

		s := strings.Split(l, "|")
		t := strings.Split(s[0], ":")

		c.winNbs = stringToInt(strings.TrimSpace(t[1]))
		c.nbs = stringToInt(strings.TrimSpace(s[1]))

		cards = append(cards, c)
	}
	return cards
}

func part1(cards []*card) {
	total := 0
	for _, c := range cards {
		score := 0
		for _, nb := range c.nbs {
			if slices.Contains(c.winNbs, nb) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		total += score
	}
	fmt.Println(total)
}

func wins(cards []*card, c *card) int {
	score := 1
	match := 1
	for _, n := range c.nbs {
		if slices.Contains(c.winNbs, n) {
			score += wins(cards, cards[c.id+match])
			match++
		}
	}
	return score
}

func part2(cards []*card) {
	total := 0
	for _, c := range cards {
		total += wins(cards, c)
	}
	fmt.Println(total)
}

func main() {
	lines := readFile()
	cards := parse(lines)

	part1(cards)
	part2(cards)
}
