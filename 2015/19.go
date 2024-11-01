package main

import (
	"bufio"
	"fmt"
	"strings"
)

func Day19() {
	content := Read_file("./files/day19.txt")
	grammar, medicineTokens, medicineStr, err := parseGrammar(content)
	if err != nil {
		panic(err)
	}
	count, _ := DistinctAfterOneReplacement(grammar, medicineTokens)
	fmt.Printf("Day19 Part 1: starting from %s, distinct molecules after one replacement = %d\n", medicineStr, count)
}

type Grammar struct {
	ByLHS     map[string][][]string
	LHSTokens map[string]string
	Alphabet  map[string]struct{}
}

type Rule struct {
	LHSKey string
	LHSTok []string
	RHSTok []string
}

func tokenizeMolecule(s string) []string {
	isUpper := func(b byte) bool { return b >= 'A' && b <= 'Z' }
	isLower := func(b byte) bool { return b >= 'a' && b <= 'z' }

	tokens := make([]string, 0, len(s))
	for i := 0; i < len(s); {
		switch {
		case s[i] == 'e': // aceptar 'e' como símbolo válido de 1 char
			tokens = append(tokens, "e")
			i++

		case isUpper(s[i]): // [A-Z][a-z]?
			if i+1 < len(s) && isLower(s[i+1]) {
				tokens = append(tokens, s[i:i+2])
				i += 2
			} else {
				tokens = append(tokens, s[i:i+1])
				i++
			}

		default:
			panic(fmt.Errorf("símbolo inválido en posición %d: %q (hex=%x)", i, s[i], s[i]))
		}
	}
	return tokens
}

func joinTokens(tokens []string) string {
	var b strings.Builder
	for _, t := range tokens {
		b.WriteString(t)
	}
	return b.String()
}

func parseGrammar(input string) (*Grammar, []string, string, error) {
	sc := bufio.NewScanner(strings.NewReader(input))

	g := &Grammar{
		ByLHS:     make(map[string][][]string),
		LHSTokens: make(map[string]string),
		Alphabet:  make(map[string]struct{}),
	}

	lines := make([]string, 0, 1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line != "" {
			lines = append(lines, strings.Trim(line, " "))
		}
	}
	if err := sc.Err(); err != nil {
		return nil, nil, "", err
	}
	if len(lines) == 0 {
		return nil, nil, "", fmt.Errorf("entrada vacía")
	}

	medicine := strings.TrimSpace(lines[len(lines)-1])
	ruleLines := lines[:len(lines)-1]

	for _, rl := range ruleLines {
		parts := strings.Split(rl, "=>")
		if len(parts) != 2 {
			return nil, nil, "", fmt.Errorf("regla no válida: %q", rl)
		}
		lhsRaw := strings.TrimSpace(parts[0])
		rhsRaw := strings.TrimSpace(parts[1])

		//remove spaces
		lhsRaw = strings.ReplaceAll(lhsRaw, " ", "")
		rhsRaw = strings.ReplaceAll(rhsRaw, " ", "")

		// Tokenizar
		lhsTok := tokenizeMolecule(lhsRaw)
		rhsTok := tokenizeMolecule(rhsRaw)

		// LHS como string clave (tal cual la entrada)
		g.ByLHS[lhsRaw] = append(g.ByLHS[lhsRaw], rhsTok)

		// Cache LHS tokenizado
		if _, ok := g.LHSTokens[lhsRaw]; !ok {
			g.LHSTokens[lhsRaw] = joinTokens(lhsTok)
		}

		// Alfabeto
		for _, t := range lhsTok {
			g.Alphabet[t] = struct{}{}
		}
		for _, t := range rhsTok {
			g.Alphabet[t] = struct{}{}
		}
	}

	medicineTokens := tokenizeMolecule(medicine)
	return g, medicineTokens, medicine, nil
}

// convert grammar to slice of rules and index for first token of LHS
func buildRulesIndex(g *Grammar) (rules []Rule, idx map[string][]int) {
	rules = make([]Rule, 0, 1024)
	for lhsKey, rhsList := range g.ByLHS {
		lhsTok := tokenizeMolecule(lhsKey)
		for _, rhsTok := range rhsList {
			rules = append(rules, Rule{
				LHSKey: lhsKey,
				LHSTok: lhsTok,
				RHSTok: rhsTok,
			})
		}
	}
	idx = make(map[string][]int)
	for i, r := range rules {
		if len(r.LHSTok) == 0 {
			continue
		}
		first := r.LHSTok[0]
		idx[first] = append(idx[first], i)
	}
	return rules, idx
}

func equalTokens(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// generate all possible molecules by applying one rule once
func DistinctAfterOneReplacement(g *Grammar, medTok []string) (int, map[string]struct{}) {
	rules, idx := buildRulesIndex(g)
	n := len(medTok)
	out := make(map[string]struct{}, 4096)

	for i := 0; i < n; i++ {
		candidates := idx[medTok[i]]
		if len(candidates) == 0 {
			continue
		}
		for _, rid := range candidates {
			r := rules[rid]
			L := len(r.LHSTok)
			if i+L > n {
				continue
			}
			if !equalTokens(medTok[i:i+L], r.LHSTok) {
				continue
			}
			// Construir la nueva molécula reemplazando SOLO esta ocurrencia
			newTok := make([]string, 0, n-L+len(r.RHSTok))
			newTok = append(newTok, medTok[:i]...)
			newTok = append(newTok, r.RHSTok...)
			newTok = append(newTok, medTok[i+L:]...)
			out[joinTokens(newTok)] = struct{}{}
		}
	}
	return len(out), out
}
