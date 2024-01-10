package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Gate interface {
	Signal() uint16
}

type Circuit struct {
	signals map[string]uint16
	gates   map[string]Gate
}

type Wire struct {
	signal uint16
}

type Not struct {
	signal uint16
}

func (w *Wire) Signal() uint16 {
	return w.signal
}

func (n *Not) Signal() uint16 {
	return ^n.signal
}

func proccesInstructions(instructions string) Circuit {
	circuit := Circuit{make(map[string]uint16), make(map[string]Gate)}
	scanner := bufio.NewScanner(strings.NewReader(instructions))
	for scanner.Scan() {
		instruction := scanner.Text()
		proccesInstruction(&circuit, instruction)
	}
	return circuit
}

func proccesInstruction(circuit *Circuit, instruction string) {
	parts := strings.Fields(instruction)
	switch len(parts) {
	case 3:
		// assignment
		signal, err := strconv.ParseUint(parts[0], 10, 16)
		if err != nil {
			circuit.gates[parts[2]] = &Wire{circuit.signals[parts[0]]}
		} else {
			circuit.signals[parts[2]] = uint16(signal)
		}
	case 4:
		// NOT
		circuit.gates[parts[3]] = &Not{circuit.signals[parts[1]]}
		circuit.signals[parts[3]] = circuit.gates[parts[3]].Signal()
	case 5:
		// AND, OR, LSHIFT, RSHIFT

	}
}

func Day7() {
	data := Read_file("files/day7.txt")
	circuit := proccesInstructions(data)
	final_signal := circuit.signals["i"]
	fmt.Println(final_signal)

}
