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

func part1(lines []string) {
	nb := 0

	for _, l := range lines {
		re := regexp.MustCompile("\\D")
		l := re.ReplaceAllString(l, "")
		nbs := strings.Split(l, "")

		n := 0
		if len(l) >= 2 {
			n, _ = strconv.Atoi(nbs[0] + nbs[len(nbs)-1])
		} else if len(l) < 2 {
			n, _ = strconv.Atoi(nbs[0] + nbs[0])
		}
		nb += n
	}
	fmt.Println(nb)
}

func part2(lines []string) {
	nb := 0

	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)

	for _, l := range lines {
		re := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")
		nbs := re.FindAllString(l, -1)
		for i, s := range nbs {
			nbs[i] = replacer.Replace(s)
		}

		n := 0
		if len(l) >= 2 {
			n, _ = strconv.Atoi(nbs[0] + nbs[len(nbs)-1])
		} else if len(l) < 2 {
			n, _ = strconv.Atoi(nbs[0] + nbs[0])
		}
		nb += n
	}
	fmt.Println(nb)
}

func main() {
	lines := readFile()
	part1(lines)
	part2(lines)
}
