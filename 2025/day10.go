package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day10() {
	input := read_file("./files/day10.txt")
	machines := parseInputMachine(input)

	totalPresses := 0
	for _, m := range machines {
		totalPresses += solveMachineBFS(m)
	}
	fmt.Println("Day 10:", totalPresses)
}

type State struct {
	Mask  uint32
	Steps int
}

type Machine struct {
	InitialState uint32
	TargetState  uint32
	Buttons      []uint32
}

func parseInputMachine(input string) []Machine {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var machines []Machine

	for _, line := range lines {

		startBracket := strings.Index(line, "[")
		endBracket := strings.Index(line, "]")
		diagram := line[startBracket+1 : endBracket]

		targetMask := parseDiagram(diagram)

		rest := line[endBracket+1:]
		braceIdx := strings.Index(rest, "{")
		if braceIdx != -1 {
			rest = rest[:braceIdx]
		}

		var buttons []uint32
		parts := strings.SplitSeq(rest, ")")
		for p := range parts {
			if !strings.Contains(p, "(") {
				continue
			}
			_, after, _ := strings.Cut(p, "(")
			numsStr := after

			btnMask := uint32(0)
			if strings.TrimSpace(numsStr) != "" {
				indices := strings.SplitSeq(numsStr, ",")
				for idxStr := range indices {
					idx, _ := strconv.Atoi(strings.TrimSpace(idxStr))
					btnMask |= (1 << idx)
				}
			}
			buttons = append(buttons, btnMask)
		}

		machines = append(machines, Machine{
			InitialState: 0,
			TargetState:  targetMask,
			Buttons:      buttons,
		})
	}
	return machines
}

func parseDiagram(diagram string) uint32 {
	mask := uint32(0)
	for i := 0; i < len(diagram); i++ {
		if diagram[i] == '#' {
			mask |= (1 << i)
		}
	}
	return mask
}

func solveMachineBFS(m Machine) int {
	startMask := uint32(0)
	target := m.TargetState

	if startMask == target {
		return 0
	}

	// stack for BFS
	queue := []State{{Mask: startMask, Steps: 0}}

	visited := make(map[uint32]bool)
	visited[startMask] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		//Test all button presses
		for _, btnMask := range m.Buttons {
			nextMask := curr.Mask ^ btnMask

			if nextMask == target {
				return curr.Steps + 1
			}

			if !visited[nextMask] {
				visited[nextMask] = true
				queue = append(queue, State{Mask: nextMask, Steps: curr.Steps + 1})
			}
		}
	}

	return -1
}
