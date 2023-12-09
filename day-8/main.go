package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

type node struct {
	tag   string
	right string
	left  string
}

func parse(lines []string) (string, map[string]*node) {
	path := lines[0]

	maps := map[string]*node{}
	for _, l := range lines[2:] {
		re := regexp.MustCompile("[a-zA-Z\\d]+")
		t := re.FindAllString(l, -1)
		maps[t[0]] = &node{
			tag:   t[0],
			left:  t[1],
			right: t[2],
		}
	}
	return path, maps
}

func part1(p string, m map[string]*node) {
	curr := m["AAA"]
	i := 0
	score := 0
	for {
		d := string(p[i])
		if d == "R" {
			curr = m[curr.right]
			score++
		} else {
			curr = m[curr.left]
			score++
		}
		if curr.tag == "ZZZ" {
			break
		}
		i++
		if i >= len(p) {
			i = 0
		}
	}
	fmt.Println(score)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2(p string, m map[string]*node) {
	currs := []*node{}
	for k, v := range m {
		if string(k[2]) == "A" {
			currs = append(currs, v)
		}
	}
	scores := []int{}
	for _, curr := range currs {
		i := 0
		score := 0
		for {
			d := string(p[i])
			if d == "R" {
				curr = m[curr.right]
			} else {
				curr = m[curr.left]
			}
			score++

			if string(curr.tag[2]) == "Z" {
				break
			}
			i++
			if i >= len(p) {
				i = 0
			}
		}
		scores = append(scores, score)
	}
	fmt.Println(LCM(scores[0], scores[1], scores...))
}

func main() {
	lines := readFile()
	p, m := parse(lines)

	//part1(p, m)
	part2(p, m)
}
