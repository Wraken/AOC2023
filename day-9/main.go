package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

type dataset struct {
	values []int
	pred   [][]int
}

func parse(lines []string) []*dataset {
	datasets := []*dataset{}
	for _, l := range lines {
		t := strings.Split(l, " ")
		data := []int{}
		for _, d := range t {
			n, _ := strconv.Atoi(d)
			data = append(data, n)
		}
		datasets = append(datasets, &dataset{values: data})
	}
	return datasets
}

func calcPred(d *dataset) [][]int {
	preds := [][]int{}
	preds = append(preds, d.values)
	curr := d.values
	for {
		p := []int{}
		for i := 0; i < len(curr)-1; i++ {
			p = append(p, curr[i+1]-curr[i])
		}
		preds = append(preds, p)
		curr = p
		if slices.Min(curr) == 0 && slices.Max(curr) == 0 {
			break
		}
	}
	return preds
}

func part1(datasets []*dataset) {
	score := 0

	for _, d := range datasets {
		preds := calcPred(d)
		v := 0
		for i := len(preds) - 1; i >= 0; i-- {
			v += preds[i][len(preds[i])-1]
		}
		score += v
	}

	fmt.Println(score)
}

func part2(datasets []*dataset) {
	score := 0

	for _, d := range datasets {
		preds := calcPred(d)
		v := 0
		for i := len(preds) - 1; i >= 0; i-- {
			// a - x = b -> -x = b - a
			v = (v - preds[i][0]) * -1
			fmt.Println(v)
		}
		score += v
	}

	fmt.Println(score)
}

func main() {
	lines := readFile()
	datas := parse(lines)
	part1(datas)
	part2(datas)
}
