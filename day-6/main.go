package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	maxTime int
	maxDist int
}

func parseFile(lines []string) []*game {
	space := regexp.MustCompile(`\s+`)
	times := strings.Split(space.ReplaceAllString(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " "), " ")
	dists := strings.Split(space.ReplaceAllString(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " "), " ")

	games := make([]*game, len(times))
	for i, t := range times {
		t, _ := strconv.Atoi(t)
		d, _ := strconv.Atoi(dists[i])

		games[i] = &game{
			maxTime: t,
			maxDist: d,
		}
	}
	return games
}

func part1(games []*game) {
	total := 1
	for _, g := range games {
		win := 0
		for i := 0; i <= g.maxTime; i++ {
			if (g.maxTime-i)*i > g.maxDist {
				win++
			}
		}
		total *= win
	}
	fmt.Println(total)
}

func main() {
	lines := readFile()
	games := parseFile(lines)

	part1(games)

	part1([]*game{
		{
			maxTime: 56977875,
			maxDist: 546192711311139,
		},
	})
}
