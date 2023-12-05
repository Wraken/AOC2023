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

type rg struct {
	srcFrom  int64
	srcTo    int64
	destFrom int64
	destTo   int64
	len      int64
}

type cor struct {
	rgs  []*rg
	next *cor
}

func stringToInt(s []string) []int64 {
	ints := make([]int64, len(s))
	for i, n := range s {
		nb, _ := strconv.ParseInt(n, 10, 64)
		ints[i] = nb
	}
	return ints
}

func parse(lines []string) ([]int64, *cor) {
	regex := regexp.MustCompile("\\d+")
	nbs := regex.FindAllString(strings.ReplaceAll(lines[0], "seeds: ", ""), -1)
	seeds := stringToInt(nbs)

	origin := &cor{}
	curr := origin
	currRgs := []*rg{}
	for _, l := range lines[3:] {
		if len(l) == 0 {
			continue
		}
		if strings.Contains(l, "map") {
			curr.next = &cor{}
			curr.rgs = currRgs
			curr = curr.next
			currRgs = []*rg{}
			continue
		}
		nbs := stringToInt(regex.FindAllString(l, -1))
		dest := nbs[0]
		src := nbs[1]
		lgt := nbs[2]

		currRgs = append(currRgs, &rg{
			srcFrom:  src,
			srcTo:    src + (lgt - 1),
			destFrom: dest,
			destTo:   dest + (lgt - 1),
			len:      lgt,
		})

	}
	curr.next = nil
	curr.rgs = currRgs
	return seeds, origin
}

func findPath(s int64, curr *cor) int64 {
	if curr != nil {
		for _, r := range curr.rgs {
			if s >= r.srcFrom && s <= r.srcTo {
				offset := s - r.srcFrom
				dest := r.destFrom + offset
				s = findPath(dest, curr.next)
				return s
			}
		}
		s = findPath(s, curr.next)
	}

	return s
}

func part1(seeds []int64, cor *cor) {
	var min int64 = -1
	for _, s := range seeds {
		d := findPath(s, cor)
		if min == -1 || min > d {
			min = d
		}
	}
	fmt.Println(min)
}

func part2(seeds []int64, cor *cor) {
	var min int64 = -1

	for i := 0; i < len(seeds); {
		var y int64 = 0
		for ; y < seeds[i+1]; y++ {
			d := findPath(seeds[i]+y, cor)
			if min == -1 || min > d {
				min = d
			}
		}
		i += 2
	}
	fmt.Println(min)
}

func main() {
	lines := readFile()
	seeds, cor := parse(lines)

	part1(seeds, cor)
	part2(seeds, cor)
}
