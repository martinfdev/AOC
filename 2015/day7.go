package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Circuit struct {
	wires   map[string]uint16
	actions []string
}

func NewCircuit() *Circuit {
	return &Circuit{
		wires: make(map[string]uint16),
	}
}

func (c *Circuit) execute(instruction string) {
	parts := strings.Split(instruction, " -> ")
	command, output := parts[0], parts[1]

	switch {
	case strings.Contains(command, "AND"):
		c.andGate(command, output)
	case strings.Contains(command, "OR"):
		c.orGate(command, output)
	case strings.Contains(command, "LSHIFT"):
		c.leftShift(command, output)
	case strings.Contains(command, "RSHIFT"):
		c.rightShift(command, output)
	case strings.Contains(command, "NOT"):
		c.notGate(command, output)
	default:
		c.directWire(command, output)
	}
}

func (c *Circuit) andGate(command, output string) {
	parts := strings.Split(command, " AND ")
	x, y := parts[0], parts[1]
	c.wires[output] = c.getSignal(x) & c.getSignal(y)
}

func (c *Circuit) orGate(command, output string) {
	parts := strings.Split(command, " OR ")
	x, y := parts[0], parts[1]
	c.wires[output] = c.getSignal(x) | c.getSignal(y)
}

func (c *Circuit) leftShift(command, output string) {
	parts := strings.Split(command, " LSHIFT ")
	x, shift := parts[0], parts[1]
	xSignal := c.getSignal(x)
	shiftValue, _ := strconv.Atoi(shift)
	c.wires[output] = xSignal << uint(shiftValue)
}

func (c *Circuit) rightShift(command, output string) {
	parts := strings.Split(command, " RSHIFT ")
	x, shift := parts[0], parts[1]
	xSignal := c.getSignal(x)
	shiftValue, _ := strconv.Atoi(shift)
	c.wires[output] = xSignal >> uint(shiftValue)
}

func (c *Circuit) notGate(command, output string) {
	x := strings.Split(command, "NOT ")[1]
	c.wires[output] = ^c.getSignal(x)
}

func (c *Circuit) directWire(command, output string) {
	x, _ := strconv.Atoi(command)
	c.wires[output] = uint16(x)
}

func (c *Circuit) getSignal(wire string) uint16 {
	if signal, ok := c.wires[wire]; ok {
		return signal
	}
	return 0
}

func proccesInstructions(data string) *Circuit {
	circuit := NewCircuit()
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		circuit.actions = append(circuit.actions, line)
	}
	for _, action := range circuit.actions {
		circuit.execute(action)
	}
	//update values for part 2
	for _, action := range circuit.actions {
		circuit.execute(action)
	}

	return circuit
}

func Day7() {
	data := Read_file("files/day7.txt")
	circuit := proccesInstructions(data)
	final_signal := circuit.getSignal("a")
	fmt.Println(final_signal)
}
