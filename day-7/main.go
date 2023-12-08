package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
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

type hand struct {
	cards string
	t     int
	bid   int
}

var cardScores = map[rune]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var cardScores2 = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func countRunes(s string, part2 bool) []int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}

	if part2 && strings.Contains(s, "J") {
		var maxC rune
		maxV := 0
		for k, v := range counts {
			if k != 'J' && v > maxV {
				maxC = k
				maxV = v
			}
		}
		if maxV != 0 {
			counts[maxC] += counts['J']
			counts['J'] = 0
		}
	}

	values := make([]int, 0, len(counts))
	for _, v := range counts {
		values = append(values, v)
	}
	return values
}

func parse(lines []string, part2 bool) []*hand {
	hands := make([]*hand, 0, len(lines))
	for _, l := range lines {
		t := strings.Split(l, " ")
		c := t[0]
		bid, _ := strconv.Atoi(t[1])

		h := &hand{
			cards: c,
			bid:   bid,
		}

		rc := countRunes(c, part2)

		twoTwo := 0
		for _, r := range rc {
			if r == 2 {
				twoTwo++
			}
		}

		if slices.Contains(rc, 5) {
			h.t = 7
		} else if slices.Contains(rc, 4) {
			h.t = 6
		} else if slices.Contains(rc, 3) && slices.Contains(rc, 2) {
			h.t = 5
		} else if slices.Contains(rc, 3) {
			h.t = 4
		} else if twoTwo == 2 {
			h.t = 3
		} else if slices.Contains(rc, 2) {
			h.t = 2
		} else {
			h.t = 1
		}

		hands = append(hands, h)
	}
	return hands
}

func part1(hands []*hand) {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].t != hands[j].t {
			return hands[i].t < hands[j].t
		}
		for k := 0; k < 5; k++ {
			score1 := cardScores[rune(hands[i].cards[k])]
			score2 := cardScores[rune(hands[j].cards[k])]
			if score1 != score2 {
				return score1 < score2
			}
		}
		return false
	})

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	fmt.Println(total)
}

func part2(hands []*hand) {
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].t != hands[j].t {
			return hands[i].t < hands[j].t
		}
		for k := 0; k < 5; k++ {
			score1 := cardScores2[rune(hands[i].cards[k])]
			score2 := cardScores2[rune(hands[j].cards[k])]
			if score1 != score2 {
				return score1 < score2
			}
		}
		return false
	})

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	fmt.Println(total)
}

func main() {
	lines := readFile()
	h := parse(lines, false)
	part1(h)

	h = parse(lines, true)
	part2(h)
}
