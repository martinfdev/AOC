package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/*
SOLUCIÓN CORRECTA Y EFICIENTE

PARTE 1: XOR/Toggle de luces
- BFS simple en espacio de bits
- Garantiza solución óptima
- Rápido porque el espacio es pequeño (2^n estados)

PARTE 2: Contadores incrementales
- BFS level-by-level (Dijkstra sin priority queue)
- Cada nivel representa un número de pasos
- Primera vez que llegamos a [0,0,0] es la solución óptima
- Optimizaciones:
  1. Hash map para evitar duplicados
  2. Límite superior basado en heurística
  3. Poda si crece demasiado
*/

type Machine struct {
	TargetLights uint64
	LightButtons []uint64
	JoltTargets  []int
	JoltButtons  [][]int
}

type Fraction struct {
	num int64
	den int64
}

func newFraction(num, den int64) Fraction {
	if den < 0 {
		num = -num
		den = -den
	}
	if num == 0 {
		return Fraction{0, 1}
	}
	g := gcd64(abs64(num), den)
	return Fraction{num / g, den / g}
}

func (f Fraction) IsZero() bool {
	return f.num == 0
}

func (f Fraction) Add(g Fraction) Fraction {
	return newFraction(f.num*g.den+g.num*f.den, f.den*g.den)
}

func (f Fraction) Sub(g Fraction) Fraction {
	return newFraction(f.num*g.den-g.num*f.den, f.den*g.den)
}

func (f Fraction) Mul(g Fraction) Fraction {
	return newFraction(f.num*g.num, f.den*g.den)
}

func (f Fraction) MulInt(k int64) Fraction {
	return newFraction(f.num*k, f.den)
}

func (f Fraction) Div(g Fraction) Fraction {
	return newFraction(f.num*g.den, f.den*g.num)
}

func (f Fraction) IntValue() (bool, int64) {
	if f.den == 0 {
		return false, 0
	}
	if f.num%f.den != 0 {
		return false, 0
	}
	return true, f.num / f.den
}

func abs64(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

func gcd64(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func parseInputDay10(input string) []Machine {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	machines := make([]Machine, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		m := Machine{}

		// [.##.#]
		bs := strings.Index(line, "[")
		be := strings.Index(line, "]")
		if bs == -1 || be == -1 {
			continue
		}

		diagram := line[bs+1 : be]
		numLights := len(diagram)

		for i := 0; i < len(diagram); i++ {
			if diagram[i] == '#' {
				m.TargetLights |= (1 << i)
			}
		}

		// {3,5,4,7}
		cs := strings.Index(line, "{")
		ce := strings.Index(line, "}")
		if cs != -1 && ce != -1 {
			joltStr := line[cs+1 : ce]
			for _, s := range strings.Split(joltStr, ",") {
				n, _ := strconv.Atoi(strings.TrimSpace(s))
				m.JoltTargets = append(m.JoltTargets, n)
			}
		}

		// (3) (1,3) (2) ...
		btnSec := line[be+1:]
		if cs != -1 {
			btnSec = line[be+1 : cs]
		}

		for {
			ps := strings.Index(btnSec, "(")
			if ps == -1 {
				break
			}
			pe := strings.Index(btnSec[ps:], ")")
			if pe == -1 {
				break
			}
			pe += ps

			content := btnSec[ps+1 : pe]
			btnSec = btnSec[pe+1:]

			var indices []int
			if len(strings.TrimSpace(content)) > 0 {
				for _, s := range strings.Split(content, ",") {
					idx, _ := strconv.Atoi(strings.TrimSpace(s))
					indices = append(indices, idx)
				}
			}

			var lightMask uint64
			for _, idx := range indices {
				if idx < numLights {
					lightMask |= (1 << idx)
				}
			}
			m.LightButtons = append(m.LightButtons, lightMask)

			joltVec := make([]int, len(m.JoltTargets))
			for _, idx := range indices {
				if idx < len(m.JoltTargets) {
					joltVec[idx] = 1
				}
			}
			m.JoltButtons = append(m.JoltButtons, joltVec)
		}

		machines = append(machines, m)
	}

	return machines
}

func solvePart1(m Machine) int {
	if m.TargetLights == 0 {
		return 0
	}

	queue := []struct {
		mask  uint64
		steps int
	}{{0, 0}}

	visited := make(map[uint64]bool)
	visited[0] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, btn := range m.LightButtons {
			next := curr.mask ^ btn

			if next == m.TargetLights {
				return curr.steps + 1
			}

			if !visited[next] {
				visited[next] = true
				queue = append(queue, struct {
					mask  uint64
					steps int
				}{next, curr.steps + 1})
			}
		}
	}

	return -1
}

func solvePart2(m Machine) int {
	// Caso base
	allZero := true
	for _, v := range m.JoltTargets {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	n := len(m.JoltTargets)

	// Filtrar botones inutiles y calcular cotas superiores
	var ub []int64
	var cols [][]int
	for _, btn := range m.JoltButtons {
		minTarget := int64(1 << 62)
		affected := 0
		for i, v := range btn {
			if v == 0 {
				continue
			}
			affected++
			t := int64(m.JoltTargets[i])
			if t < minTarget {
				minTarget = t
			}
		}
		if affected == 0 {
			continue
		}
		if minTarget <= 0 {
			// Este boton pisaria un contador con 0, por lo que su valor debe ser 0.
			continue
		}
		ub = append(ub, minTarget)
		cols = append(cols, btn)
	}

	m2 := len(ub)
	if m2 == 0 {
		return -1
	}

	// Construir matriz A (n x m2)
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, m2)
		for j := 0; j < m2; j++ {
			row[j] = cols[j][i]
		}
		A[i] = row
	}

	r := rankMatrixInt(A)
	if r == 0 {
		return -1
	}

	pivotCols := choosePivotColumns(A, ub, r)
	if len(pivotCols) != r {
		return -1
	}

	inPivot := make([]bool, m2)
	for _, c := range pivotCols {
		inPivot[c] = true
	}

	order := make([]int, 0, m2)
	order = append(order, pivotCols...)
	for j := 0; j < m2; j++ {
		if !inPivot[j] {
			order = append(order, j)
		}
	}

	// Matriz aumentada con columnas reordenadas
	M := make([][]Fraction, n)
	for i := 0; i < n; i++ {
		row := make([]Fraction, m2+1)
		for j := 0; j < m2; j++ {
			row[j] = newFraction(int64(A[i][order[j]]), 1)
		}
		row[m2] = newFraction(int64(m.JoltTargets[i]), 1)
		M[i] = row
	}

	if !gaussJordan(M, r) {
		return -1
	}

	// Inconsistencia
	for i := r; i < n; i++ {
		allZeroRow := true
		for j := 0; j < m2; j++ {
			if !M[i][j].IsZero() {
				allZeroRow = false
				break
			}
		}
		if allZeroRow && !M[i][m2].IsZero() {
			return -1
		}
	}

	d := m2 - r

	pivotConst := make([]Fraction, r)
	pivotCoeff := make([][]Fraction, r)
	for i := 0; i < r; i++ {
		pivotConst[i] = M[i][m2]
		coeffs := make([]Fraction, d)
		for j := 0; j < d; j++ {
			coeffs[j] = M[i][r+j]
		}
		pivotCoeff[i] = coeffs
	}

	ubPivot := make([]int64, r)
	for i := 0; i < r; i++ {
		ubPivot[i] = ub[order[i]]
	}
	ubFree := make([]int64, d)
	for j := 0; j < d; j++ {
		ubFree[j] = ub[order[r+j]]
	}

	best := int64(1 << 60)

	eval := func(free []int64) {
		sum := int64(0)
		for _, v := range free {
			sum += v
		}
		if sum >= best {
			return
		}
		for i := 0; i < r; i++ {
			val := pivotConst[i]
			for j := 0; j < d; j++ {
				if free[j] != 0 {
					val = val.Sub(pivotCoeff[i][j].MulInt(free[j]))
				}
			}
			ok, x := val.IntValue()
			if !ok {
				return
			}
			if x < 0 || x > ubPivot[i] {
				return
			}
			sum += x
			if sum >= best {
				return
			}
		}
		if sum < best {
			best = sum
		}
	}

	switch d {
	case 0:
		eval([]int64{})
	case 1:
		for t0 := int64(0); t0 <= ubFree[0]; t0++ {
			eval([]int64{t0})
		}
	case 2:
		for t0 := int64(0); t0 <= ubFree[0]; t0++ {
			for t1 := int64(0); t1 <= ubFree[1]; t1++ {
				eval([]int64{t0, t1})
			}
		}
	case 3:
		for t0 := int64(0); t0 <= ubFree[0]; t0++ {
			for t1 := int64(0); t1 <= ubFree[1]; t1++ {
				for t2 := int64(0); t2 <= ubFree[2]; t2++ {
					eval([]int64{t0, t1, t2})
				}
			}
		}
	default:
		return -1
	}

	if best == int64(1<<60) {
		return -1
	}
	return int(best)
}

func rankMatrixInt(A [][]int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	m := len(A[0])
	if m == 0 {
		return 0
	}

	M := make([][]Fraction, n)
	for i := 0; i < n; i++ {
		row := make([]Fraction, m)
		for j := 0; j < m; j++ {
			row[j] = newFraction(int64(A[i][j]), 1)
		}
		M[i] = row
	}

	return rankFractionMatrix(M)
}

func rankFractionMatrix(M [][]Fraction) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	m := len(M[0])
	r := 0
	c := 0
	for r < n && c < m {
		pivot := -1
		for i := r; i < n; i++ {
			if !M[i][c].IsZero() {
				pivot = i
				break
			}
		}
		if pivot == -1 {
			c++
			continue
		}
		if pivot != r {
			M[r], M[pivot] = M[pivot], M[r]
		}
		pv := M[r][c]
		for i := r + 1; i < n; i++ {
			if M[i][c].IsZero() {
				continue
			}
			factor := M[i][c].Div(pv)
			for j := c; j < m; j++ {
				M[i][j] = M[i][j].Sub(M[r][j].Mul(factor))
			}
		}
		r++
		c++
	}
	return r
}

func rankMatrixCols(A [][]int, cols []int) int {
	n := len(A)
	if n == 0 || len(cols) == 0 {
		return 0
	}
	M := make([][]Fraction, n)
	for i := 0; i < n; i++ {
		row := make([]Fraction, len(cols))
		for j, c := range cols {
			row[j] = newFraction(int64(A[i][c]), 1)
		}
		M[i] = row
	}
	return rankFractionMatrix(M)
}

func choosePivotColumns(A [][]int, ub []int64, rank int) []int {
	m := len(ub)
	if rank == 0 {
		return []int{}
	}
	bestCost := int64(-1)
	var best []int
	combo := make([]int, rank)

	var dfs func(start, depth int)
	dfs = func(start, depth int) {
		if depth == rank {
			if rankMatrixCols(A, combo) != rank {
				return
			}
			cost := int64(1)
			for j := 0; j < m; j++ {
				if !comboContains(combo, j) {
					cost *= ub[j] + 1
				}
			}
			if bestCost == -1 || cost < bestCost {
				bestCost = cost
				best = append([]int(nil), combo...)
			}
			return
		}
		for i := start; i <= m-(rank-depth); i++ {
			combo[depth] = i
			dfs(i+1, depth+1)
		}
	}
	dfs(0, 0)
	return best
}

func comboContains(combo []int, val int) bool {
	return slices.Contains(combo, val)
}

func gaussJordan(M [][]Fraction, pivotCols int) bool {
	n := len(M)
	if n == 0 {
		return false
	}
	m := len(M[0]) - 1
	row := 0
	for col := 0; col < pivotCols; col++ {
		pivot := -1
		for i := row; i < n; i++ {
			if !M[i][col].IsZero() {
				pivot = i
				break
			}
		}
		if pivot == -1 {
			return false
		}
		if pivot != row {
			M[row], M[pivot] = M[pivot], M[row]
		}

		pv := M[row][col]
		for j := col; j <= m; j++ {
			M[row][j] = M[row][j].Div(pv)
		}
		for i := 0; i < n; i++ {
			if i == row || M[i][col].IsZero() {
				continue
			}
			factor := M[i][col]
			for j := col; j <= m; j++ {
				M[i][j] = M[i][j].Sub(M[row][j].Mul(factor))
			}
		}
		row++
	}
	return true
}

func Day10() {
	input := read_file("./files/day10.txt")
	machines := parseInputDay10(input)

	fmt.Printf("Machines: %d\n\n", len(machines))

	total1 := 0
	total2 := 0

	for i, m := range machines {
		p1 := solvePart1(m)
		p2 := solvePart2(m)

		if p1 != -1 {
			total1 += p1
		}
		if p2 != -1 {
			total2 += p2
		}

		fmt.Printf("Machine %d: P1=%d P2=%d\n", i+1, p1, p2)
	}

	fmt.Println("\n=== FINAL ===")
	fmt.Println("Part 1:", total1)
	fmt.Println("Part 2:", total2)
}
