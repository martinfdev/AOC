package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Gate interface {
	Signal() uint16
}

type Wire struct {
	signal uint16
}

type Circuit struct {
	signals map[string]uint16
	gates   map[string]Gate
}

type AndGate struct {
	input1, input2 Gate
}

type NotGate struct {
	input Gate
}

func (w *Wire) SetSignal(s uint16) {
	w.signal = s
}

func (w *Wire) Signal() uint16 {
	return w.signal
}

func (g *AndGate) Signal() uint16 {
	return g.input1.Signal() & g.input2.Signal()
}

func (g *NotGate) Signal() uint16 {
	return g.input.Signal()
}

func proccesInstructions(instructions string) Circuit {
	circuit := Circuit{
		signals: make(map[string]uint16),
		gates:   make(map[string]Gate),
	}
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
		// 123 -> x
		signal, err := strconv.ParseUint(parts[0], 10, 16)
		if err != nil {
			circuit.gates[parts[2]] = &Wire{}
		} else {
			circuit.signals[parts[2]] = uint16(signal)
		}
	case 4:
		// NOT e -> f
		gate := &NotGate{input: parts[1]}
		circuit.gates[parts[3]] = gate
	}
}

func Day7() {
	data := Read_file("files/day7.txt")

}
