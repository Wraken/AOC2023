package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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

type sym struct {
	x      int
	y      int
	l      int
	symbol rune
}
type nb struct {
	n    int
	calc bool
}

func parseMap(lines []string) (map[int]*nb, []*sym) {
	m := map[int]*nb{}
	symbols := []*sym{}

	for y, l := range lines {
		r := []rune(l)
		for x := 0; x < len(r); {
			if unicode.IsDigit(r[x]) {
				ret := x
				for x < len(r) && unicode.IsDigit(r[x]) {
					x++
				}
				n, _ := strconv.Atoi(l[ret:x])
				nb := &nb{
					n:    n,
					calc: false,
				}
				for ; ret < x; ret++ {
					m[y*len(r)+ret] = nb
				}
				continue
			} else if r[x] != '.' {
				symbols = append(symbols, &sym{
					x:      x,
					y:      y,
					l:      len(r),
					symbol: r[x],
				})
			}
			x++
		}
	}
	return m, symbols
}

func part1(lines []string) {
	m, sym := parseMap(lines)

	total := 0
	for _, s := range sym {
		if n, ok := m[(s.y-1)*s.l+s.x]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[(s.y+1)*s.l+s.x]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[s.y*s.l+s.x+1]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[(s.y*s.l)+(s.x-1)]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[(s.y-1)*s.l+s.x+1]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[(s.y-1)*s.l+s.x-1]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[(s.y+1)*s.l+s.x+1]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
		if n, ok := m[(s.y+1)*s.l+s.x-1]; ok {
			if !n.calc {
				total += n.n
				n.calc = true
			}
		}
	}
	fmt.Println(total)
}

func findAdjacent(m map[int]*nb, s *sym) []*nb {
	nbs := []*nb{}
	if n, ok := m[(s.y-1)*s.l+s.x]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[(s.y+1)*s.l+s.x]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[s.y*s.l+s.x+1]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[(s.y*s.l)+(s.x-1)]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[(s.y-1)*s.l+s.x+1]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[(s.y-1)*s.l+s.x-1]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[(s.y+1)*s.l+s.x+1]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	if n, ok := m[(s.y+1)*s.l+s.x-1]; ok {
		if !n.calc {
			n.calc = true
			nbs = append(nbs, n)
		}
	}
	return nbs
}

func part2(lines []string) {
	m, sym := parseMap(lines)

	total := 0
	for _, s := range sym {
		if s.symbol == '*' {
			nbs := findAdjacent(m, s)
			if len(nbs) == 2 {
				total += nbs[0].n * nbs[1].n
			}
			for _, a := range nbs {
				a.calc = false
			}
		}
	}
	fmt.Println(total)
}

func main() {
	lines := readFile()
	part1(lines)

	part2(lines)
}
