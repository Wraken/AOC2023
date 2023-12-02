package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type game struct {
	nb       int
	maxRed   int
	maxBlue  int
	maxGreen int
	minRed   int
	minBlue  int
	minGreen int
	power    int
}

func parsGames(lines []string) []*game {
	games := make([]*game, 0, len(lines))

	for _, l := range lines {
		g := &game{}
		t := strings.Split(l, ":")
		g.nb, _ = strconv.Atoi(strings.ReplaceAll(t[0], "Game ", ""))

		t = strings.Split(strings.ReplaceAll(t[1], ";", ","), ",")
		for _, c := range t {
			cc := strings.Split(strings.TrimSpace(c), " ")
			fmt.Println(cc)
			if cc[1] == "red" {
				nb, _ := strconv.Atoi(cc[0])
				if g.maxRed < nb {
					g.maxRed = nb
				}
			}
			if cc[1] == "green" {
				nb, _ := strconv.Atoi(cc[0])
				if g.maxGreen < nb {
					g.maxGreen = nb
				}
			}
			if cc[1] == "blue" {
				nb, _ := strconv.Atoi(cc[0])
				if g.maxBlue < nb {
					g.maxBlue = nb
				}
			}
		}
		g.power = g.maxBlue * g.maxGreen * g.maxRed
		games = append(games, g)
	}
	return games
}

func part1(lines []string) {
	games := parsGames(lines)
	maxR := 12
	maxG := 13
	maxB := 14
	total := 0

	for _, g := range games {
		if g.maxRed <= maxR && g.maxGreen <= maxG && g.maxBlue <= maxB {
			total += g.nb
		}
	}
	print(total)
}

func part2(lines []string) {
	games := parsGames(lines)
	total := 0
	for _, g := range games {
		total += g.power
	}
	fmt.Println(total)
}

func main() {
	lines := readFile()
	part1(lines)
	part2(lines)
}
