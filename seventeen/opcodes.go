package seventeen

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Instruction int

const (
	adv Instruction = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type Program struct {
	A, B, C  int
	InstrPtr int
	Program  []int
	StdOut   []int
}

func (p *Program) Combo(op int) int {
	switch op {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return op
	case 4:
		return p.A
	case 5:
		return p.B
	case 6:
		return p.C
	case 7:
		// invalid
		return 0
	}

	return 0
}

func (p *Program) Exec() {
	var opcode Instruction
	var operand int

	opcode, operand = Instruction(p.Program[p.InstrPtr]), p.Program[p.InstrPtr+1]

	fmt.Printf("[%d] %d\n", opcode, operand)

	switch opcode {
	case adv:
		p.A = int(float64(p.A) / math.Pow(2, float64(p.Combo(operand))))
	case bxl:
		p.B ^= operand
	case bst:
		p.B = p.Combo(operand) % 8
	case jnz:
		if p.A != 0 {
			p.InstrPtr = operand
			return
		}
	case bxc:
		p.B = p.B ^ p.C
	case out:
		p.StdOut = append(p.StdOut, p.Combo(operand)%8)
	case bdv:
		p.B = int(float64(p.A) / math.Pow(2, float64(p.Combo(operand))))
	case cdv:
		p.C = int(float64(p.A) / math.Pow(2, float64(p.Combo(operand))))
	}

	p.InstrPtr += 2
}

func (p Program) String() string {
	var strs []string
	for _, s := range p.StdOut {
		strs = append(strs, strconv.Itoa(s))
	}

	return strings.Join(strs, ",")
}

func (p *Program) Run() string {
	for p.InstrPtr < len(p.Program) {
		p.Exec()
	}

	fmt.Printf("%v\n", p.StdOut)

	return p.String()
}

func ParseProg(r io.Reader) *Program {
	var prg Program

	fmt.Fscanf(r, "Register A: %d\n", &prg.A)
	fmt.Fscanf(r, "Register B: %d\n", &prg.B)
	fmt.Fscanf(r, "Register C: %d\n\n", &prg.C)

	var progStr string
	fmt.Fscanf(r, "Program: %s", &progStr)

	fmt.Printf("Gave %s\n", progStr)
	progs := strings.Split(progStr, ",")
	for _, i := range progs {
		n, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		prg.Program = append(prg.Program, n)
	}

	return &prg
}
