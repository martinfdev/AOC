package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Wire struct {
	name   string
	signal uint16
	status bool
}

type Circuit struct {
	wires   map[string]Wire
	actions []string
}

func NewCircuit() *Circuit {
	return &Circuit{
		wires: make(map[string]Wire),
	}
}

func (c *Circuit) execute(instruction string) bool {
	parts := strings.Split(instruction, " -> ")
	command, output := parts[0], parts[1]

	switch {
	case strings.Contains(command, "AND"):
		return c.andGate(command, output)
	case strings.Contains(command, "OR"):
		return c.orGate(command, output)
	case strings.Contains(command, "LSHIFT"):
		return c.leftShift(command, output)
	case strings.Contains(command, "RSHIFT"):
		return c.rightShift(command, output)
	case strings.Contains(command, "NOT"):
		return c.notGate(command, output)
	default:
		return c.directWire(command, output)
	}
}

func (c *Circuit) andGate(command, output string) bool {
	parts := strings.Split(command, " AND ")
	x, y := parts[0], parts[1]
	wire1 := c.getSignal(x)
	wire2 := c.getSignal(y)
	if wire1.status != false && wire2.status != false {
		result := wire1.signal & wire2.signal
		c.wires[output] = Wire{name: output, signal: result, status: true}
		return true
	}
	val_x, err := strconv.Atoi(x)
	if err == nil {
		wire2 := c.getSignal(y)
		if wire2.status != false {
			result := uint16(val_x) & wire2.signal
			c.wires[output] = Wire{name: output, signal: result, status: true}
			return true
		}
	}
	val_y, err := strconv.Atoi(y)
	if err == nil {
		wire1 := c.getSignal(x)
		if wire1.status != false {
			result := wire1.signal & uint16(val_y)
			c.wires[output] = Wire{name: output, signal: result, status: true}
			return true
		}
	}
	return false
}

func (c *Circuit) orGate(command, output string) bool {
	parts := strings.Split(command, " OR ")
	x, y := parts[0], parts[1]
	wire1 := c.getSignal(x)
	wire2 := c.getSignal(y)
	if wire1.status != false && wire2.status != false {
		result := wire1.signal | wire2.signal
		c.wires[output] = Wire{name: output, signal: result, status: true}
		return true
	}
	val_x, err := strconv.Atoi(x)
	if err == nil {
		wire2 := c.getSignal(y)
		if wire2.status != false {
			result := uint16(val_x) | wire2.signal
			c.wires[output] = Wire{name: output, signal: result, status: true}
			return true
		}
	}
	val_y, err := strconv.Atoi(y)
	if err == nil {
		wire1 := c.getSignal(x)
		if wire1.status != false {
			result := wire1.signal | uint16(val_y)
			c.wires[output] = Wire{name: output, signal: result, status: true}
			return true
		}
	}
	return false
}

func (c *Circuit) leftShift(command, output string) bool {
	parts := strings.Split(command, " LSHIFT ")
	x, shift := parts[0], parts[1]
	wire := c.getSignal(x)
	if wire.status != false {
		shiftValue, _ := strconv.Atoi(shift)
		result := wire.signal << uint(shiftValue)
		c.wires[output] = Wire{name: output, signal: result, status: true}
		return true
	}
	return false
}

func (c *Circuit) rightShift(command, output string) bool {
	parts := strings.Split(command, " RSHIFT ")
	x, shift := parts[0], parts[1]
	wire := c.getSignal(x)
	if wire.status != false {
		shiftValue, _ := strconv.Atoi(shift)
		result := wire.signal >> uint(shiftValue)
		c.wires[output] = Wire{name: output, signal: result, status: true}
		return true
	}
	return false
}

func (c *Circuit) notGate(command, output string) bool {
	x := strings.Split(command, "NOT ")[1]
	wire := c.getSignal(x)
	if wire.status != false {
		result := ^wire.signal
		c.wires[output] = Wire{name: output, signal: result, status: true}
		return true
	}
	return false
}

func (c *Circuit) directWire(command, output string) bool {
	x, err := strconv.Atoi(command)
	if err == nil {
		c.wires[output] = Wire{name: output, signal: uint16(x), status: true}
		return true
	} else {
		wire := c.getSignal(command)
		if wire.status != false {
			c.wires[output] = Wire{name: output, signal: wire.signal, status: true}
			return true
		}
		return false
	}
}

func (c *Circuit) getSignal(wire string) Wire {
	if wire, ok := c.wires[wire]; ok {
		return wire
	}
	return Wire{}
}

func proccesInstructions(data string) *Circuit {
	circuit := NewCircuit()
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		circuit.actions = append(circuit.actions, line)
	}
FOR:
	for i, action := range circuit.actions {
		result := circuit.execute(action)
		if result {
			//remove action from slice
			circuit.actions = append(circuit.actions[:i], circuit.actions[i+1:]...)
			goto FOR
		}
	}
	if len(circuit.actions) > 0 {
		goto FOR
	}
	return circuit
}

func Day7() {
	data := Read_file("files/day7.txt")
	circuit := proccesInstructions(data)
	final_signal := circuit.getSignal("a").signal
	fmt.Println(final_signal)
}
