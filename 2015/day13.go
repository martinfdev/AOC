package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	ID    map[string]int
	Names []string
	W     [][]int
	Pair  [][]int
}

func Day13() {
	content := Read_file("files/day13.txt")
	reader := strings.NewReader(content)
	prob, err := parseInput(reader)
	if err != nil {
		panic(err)
	}

	best, order := MaxCycleDP(prob)
	fmt.Printf("Day 13 Part 1: Maximum happiness is %d\n", best)
	fmt.Printf("Order: ")
	for _, i := range order {
		fmt.Printf("%s ", prob.Names[i])
	}

}

// parseInput reads the problem input from r and returns a Problem instance.
func parseInput(r io.Reader) (*Problem, error) {
	re := regexp.MustCompile(`^([A-Za-z]+) would (gain|lose) (\d+) happiness units by sitting next to ([A-Za-z]+)\.$`)

	id := map[string]int{}
	var names []string
	get := func(s string) int {
		if i, ok := id[s]; ok {
			return i
		}
		id[s] = len(names)
		names = append(names, s)
		return id[s]
	}

	type edge struct{ a, b, v int }
	var es []edge

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		m := re.FindStringSubmatch(line)
		if m == nil {
			return nil, fmt.Errorf("invalid line: %q", line)
		}
		val, _ := strconv.Atoi(m[3])
		if m[2] == "lose" {
			val = -val
		}
		es = append(es, edge{get(m[1]), get(m[4]), val})
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	n := len(names)
	W := make([][]int, n)
	Pair := make([][]int, n)
	for i := range W {
		W[i] = make([]int, n)
		Pair[i] = make([]int, n)
	}
	for _, e := range es {
		W[e.a][e.b] = e.v
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s := W[i][j] + W[j][i]
			Pair[i][j] = s
			Pair[j][i] = s
		}
	}
	return &Problem{ID: id, Names: names, W: W, Pair: Pair}, nil
}

// Solve problem with algoritm Held–Karp for TSP
func MaxCycleDP(prob *Problem) (int, []int) {
	n := len(prob.Names)
	pair := prob.Pair
	if n == 0 {
		return 0, nil
	}

	s := 0
	full := (1 << n) - 1

	//dp[mask][list] -> int minimo muy negativo para "no alcanzable"
	const NEG = int(^uint(0)>>1) * -1 //aproximacion a -inf
	dp := make([][]int, 1<<n)
	par := make([][]int, 1<<n)
	for m := range dp {
		dp[m] = make([]int, n)
		par[m] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[m][j] = NEG
			par[m][j] = -1
		}
	}
	dp[1<<s][s] = 0

	// Iterate masks with included s
	for mask := 0; mask <= full; mask++ {
		if (mask & (1 << s)) == 0 {
			continue
		}
		for last := 0; last < n; last++ {
			if dp[mask][last] == NEG {
				continue
			}
			//Iterate adding "next" no used
			for next := 0; next < n; next++ {
				if (mask & (1 << next)) != 0 {
					continue
				}
				nmask := mask | (1 << next)
				cand := dp[mask][last] + pair[last][next]
				if cand > dp[nmask][next] {
					dp[nmask][next] = cand
					par[nmask][next] = last
				}
			}
		}
	}

	// close the cycle return to s
	best := NEG
	bestLast := -1
	for last := 0; last < n; last++ {
		if last == s || dp[full][last] == NEG {
			continue
		}
		cand := dp[full][last] + pair[last][s]
		if cand > best {
			best = cand
			bestLast = last
		}
	}

	// Reconstruct of the circular order starting at s
	// Lineas path s -> ... -> bestLast
	order := make([]int, 0, n)
	mask := full
	curr := bestLast
	for curr != -1 {
		order = append(order, curr)
		prev := par[mask][curr]
		if prev == -1 {
			break
		}
		mask ^= (1 << curr)
		curr = prev
	}
	// Now order is in reverse: [bestLast ... s], invert
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	//order start in s. Is a cycle;
	if len(order) > 0 && order[0] != s {
		// Rotate to start at s
		k := 0
		for k < len(order) && order[k] != s {
			k++
		}
		if k < len(order) {
			order = append(order[k:], order[:k]...)
		}
	}
	return best, order
}
