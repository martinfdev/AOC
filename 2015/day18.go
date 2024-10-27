package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"strings"
)

type Row struct {
	Lo, Hi uint64
}

type Board struct {
	Rows []Row
	W, H int
}

func Day18() {
	content := Read_file("./files/day18.txt")
	board, err := newBoardFromString(content)
	if err != nil {
		panic(err)
	}
	steps := 100
	for i := 0; i < steps; i++ {
		board.Step()
	}
	onCount := board.CountOn()
	fmt.Printf("Day18: after %d steps, lights on: %d\n", steps, onCount)
	fmt.Printf("Board:\n%s\n", board.String())
}

func NewBoard() *Board { b := &Board{W: 100, H: 100}; b.Rows = make([]Row, b.H); return b }

// set/clear/togle bit (c: 0...99)
func (b *Board) Set(r, c int) {
	if c < 64 {
		b.Rows[r].Lo |= 1 << c
	} else {
		b.Rows[r].Hi |= 1 << (c - 64)
	}
}
func (b *Board) Clear(r, c int) {
	if c < 64 {
		b.Rows[r].Lo &^= 1 << c
	} else {
		b.Rows[r].Hi &^= 1 << (c - 64)
	}
}
func (b *Board) IsOn(r, c int) bool {
	if c < 64 {
		return (b.Rows[r].Lo>>c)&1 == 1
	}
	return (b.Rows[r].Hi>>(c-64))&1 == 1
}

func shiftLeft(row Row) Row {
	carry := (row.Lo >> 63) & 1
	lo := row.Lo << 1
	hi := (row.Hi << 1) | carry
	hi &= (uint64(1) << 36) - 1 // clear bits beyond 100
	return Row{Lo: lo, Hi: hi}
}

func shiftRight(row Row) Row {
	carry := (row.Hi & 1) << 63
	hi := row.Hi >> 1
	lo := (row.Lo >> 1) | carry
	return Row{Lo: lo, Hi: hi}
}

func newBoardFromString(content string) (*Board, error) {
	b := NewBoard()
	sc := bufio.NewScanner(strings.NewReader(content))

	row := 0
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		if err := loadRow(b, row, line); err != nil {
			return nil, err
		}
		row++
		if row >= b.H {
			break
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if row != b.H {
		return nil, fmt.Errorf("expected %d rows, got %d", b.H, row)
	}
	return b, nil
}

func loadRow(b *Board, row int, line string) error {
	if row < 0 || row >= b.H {
		return fmt.Errorf("row index out of bounds: %d", row)
	}
	if len(line) < b.W {
		return fmt.Errorf("the line has %d columns, expected at least %d", len(line), b.W)
	}
	for col := 0; col < b.W; col++ {
		switch line[col] {
		case '#':
			b.Set(row, col)
		case '.':
			b.Clear(row, col)
		default:
			return fmt.Errorf("invalid character at row %d, col %d: %c", row, col, line[col])
		}
	}
	return nil
}

func (b *Board) Step() {
	var next [100]Row
	var counts [100]uint8

	for r := 0; r < b.H; r++ {
		// vecinos verticales: arriba (ru), mismo (r), abajo (rd)
		ru := Row{}
		if r > 0 {
			ru = b.Rows[r-1]
		}
		rc := b.Rows[r]
		rd := Row{}
		if r+1 < b.H {
			rd = b.Rows[r+1]
		}

		// 8 máscaras de vecinos (excluye el centro rc)
		masks := [8]Row{
			shiftLeft(ru), ru, shiftRight(ru),
			shiftLeft(rc) /* center rc excluido */, shiftRight(rc),
			shiftLeft(rd), rd, shiftRight(rd),
		}

		// reset de contadores columna a columna
		for i := range counts {
			counts[i] = 0
		}

		// acumular vecinos en counts a partir de las 8 máscaras
		accumulateCounts := func(row Row) {
			// Lo: 64 bits
			for c := 0; c < 64; c++ {
				if (row.Lo>>c)&1 == 1 {
					counts[c]++
				}
			}
			// Hi: 36 bits → columnas 64..99
			for c := 0; c < 36; c++ {
				if (row.Hi>>c)&1 == 1 {
					counts[64+c]++
				}
			}
		}
		for _, m := range masks {
			accumulateCounts(m)
		}

		// aplicar regla B3/S23 y construir next[r]
		var nlo, nhi uint64
		for c := 0; c < 64; c++ {
			cnt := counts[c]
			on := (rc.Lo>>c)&1 == 1
			keep := on && (cnt == 2 || cnt == 3)
			born := !on && (cnt == 3)
			if keep || born {
				nlo |= 1 << c
			}
		}
		for c := 0; c < 36; c++ {
			cnt := counts[64+c]
			on := (rc.Hi>>c)&1 == 1
			keep := on && (cnt == 2 || cnt == 3)
			born := !on && (cnt == 3)
			if keep || born {
				nhi |= 1 << c
			}
		}
		next[r] = Row{Lo: nlo, Hi: nhi}
	}
	b.Rows = next[:]
}

func (b *Board) CountOn() int {
	sum := 0
	for r := 0; r < b.H; r++ {
		sum += bits.OnesCount64(b.Rows[r].Lo)
		// limitar Hi a 36 bits shifts, pero OnesCount64 ignora bits altos si están a 0
		sum += bits.OnesCount64(b.Rows[r].Hi)
	}
	return sum
}

func (b *Board) String() string {
	var sb strings.Builder
	for r := 0; r < b.H; r++ {
		for c := 0; c < b.W; c++ {
			if b.IsOn(r, c) {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
