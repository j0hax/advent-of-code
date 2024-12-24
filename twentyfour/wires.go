package twentyfour

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
)

type GateType int

const (
	AND GateType = iota
	OR
	XOR
)

// A gate stores two inputs and its type
type Gate struct {
	inputs []string
	t      GateType
	dest   string
}

// InputsDefined checks if a gate's inputs are defined in the map
func (g *Gate) InputsDefined(circuit map[string]bool) bool {
	for _, i := range g.inputs {
		if _, ok := circuit[i]; !ok {
			return false
		}
	}

	return true
}

// Compute computes the gates' values iff they are available.
// If evaluation is successful, the result is then stored in the passed hashmap
// and a 1 is returned
func (g *Gate) Compute(circuit map[string]bool) bool {
	if !g.InputsDefined(circuit) {
		return false
	}

	// TODO: we could also make these gates use n inputs.
	a := circuit[g.inputs[0]]
	b := circuit[g.inputs[1]]
	var out bool

	switch g.t {
	case AND:
		out = a && b
	case OR:
		out = a || b
	case XOR:
		out = a != b
	default:
		panic(fmt.Sprintf("unexpected twentyfour.GateType: %#v", g.t))
	}

	circuit[g.dest] = out
	return true
}

type Circuit struct {
	// Inputs represents both the initial wires as well as gate outputs
	inputs map[string]bool

	// A gate is identified by its name and its required inputs
	gates []Gate
}

func (c *Circuit) Simulate() {
	// Gameplan:
	// Iteratively scan through gates, solving them if possible,
	// then adding them to inputs

	for len(c.gates) > 0 {
		for i, gate := range c.gates {
			if len(gate.inputs) == 0 {
				continue
			}
			if gate.Compute(c.inputs) {
				c.gates = slices.Delete(c.gates, i, i+1)
			}
		}
	}
}

func parseGate(s string) *Gate {
	var a, t, b, des string
	fmt.Sscanf(s, "%s %s %s -> %s", &a, &t, &b, &des)

	var gt GateType
	switch t {
	case "AND":
		gt = AND
	case "OR":
		gt = OR
	case "XOR":
		gt = XOR
	}

	return &Gate{
		inputs: []string{a, b},
		t:      gt,
		dest:   des,
	}
}

func parseWire(wires map[string]bool, s string) {
	parts := strings.Split(s, ": ")

	name := parts[0]
	res := parts[1] != "0"

	wires[name] = res
}

func ParseCircuit(r io.Reader) *Circuit {
	c := Circuit{
		inputs: make(map[string]bool),
		gates:  []Gate{},
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			parseWire(c.inputs, line)
		} else if strings.Contains(line, "->") {
			c.gates = append(c.gates, *parseGate(line))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return &c
}
