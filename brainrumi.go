/*
brainrumi - a brainfuck interpreter written in go by snarkb0t on github
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	raw       bool
	debug     bool
	optimized bool

	cells []int32 = make([]int32, 1)
	p     int     = 0
)

func evalCommand(cmd byte, increment int) {
	switch cmd {
	case '>':
		if p += increment; p > len(cells)-1 {
			for i := increment; i > 0; i-- {
				cells = append(cells, 0)
			}
		}
	case '<':
		if p == 0 {
			return
		}

		p -= increment
	case '+':
		cells[p] += int32(increment)
	case '-':
		cells[p] -= int32(increment)
	case '.':
		if raw {
			fmt.Print(cells[p])
		} else {
			fmt.Print(string(cells[p]))
		}

		if increment > 1 {
			evalCommand(cmd, increment-1)
		}
	case ',':
		var putInto string
		fmt.Scanln(&putInto)
		cells[p] = int32([]byte(putInto)[0])
	}
}

func eval(code string) {
	var (
		scopes map[string]int = make(map[string]int, 0)
		open   []int          = make([]int, 0)

		instructions [][]int = make([][]int, 0)

		tick time.Time = time.Now()
	)

	for s := 0; s < len(code); s++ {
		if optimized {

			if len(instructions) == 0 || code[s] == '[' || code[s] == ']' {
				goto addInstruction
			} else {
				latestInt := &instructions[len(instructions)-1]

				if (*latestInt)[0] == int(code[s]) {
					(*latestInt)[1]++
					continue
				} else {
					goto addInstruction
				}
			}

		addInstruction:
			secondPointer := 1

			switch code[s] {
			case '[':
				open = append(open, len(instructions))
			case ']':
				lastOpen := &(open[len(open)-1])
				instructions[*lastOpen][1] = len(instructions)
				secondPointer = *lastOpen

				open = open[:len(open)-1]
			}

			instructions = append(instructions, []int{int(code[s]), secondPointer})
		} else {
			switch code[s] {
			case '[':
				open = append(open, s)
			case ']':
				scopes[strconv.Itoa(open[len(open)-1])] = s
				scopes[strconv.Itoa(s)] = open[len(open)-1]
				open = open[:len(open)-1]
			}
		}
	}

	open = nil

	if optimized {
		for i := 0; i < len(instructions); i++ {
			instruction := instructions[i]
			switch instruction[0] {
			case '[':
				if cells[p] == 0 {
					i = instruction[1]
					continue
				}
			case ']':
				if cells[p] != 0 {
					i = instruction[1]
					continue
				}
			}

			evalCommand(byte(instruction[0]), instruction[1])
		}
	} else {
		for i := 0; i < len(code); i++ {
			switch code[i] {
			case '[':
				if cells[p] == 0 {
					i = scopes[strconv.Itoa(i)]
					continue
				}
			case ']':
				if cells[p] != 0 {
					i = scopes[strconv.Itoa(i)]
					continue
				}
			}

			evalCommand(code[i], 1)
		}
	}

	if debug {
		fmt.Println("------")
		fmt.Println("executed in", time.Since(tick), "|| optimizer", optimized)
	}
}

func main() {
	flag.BoolVar(&raw, "r", false, "display raw cell data instead of encoded text")
	flag.BoolVar(&debug, "d", false, "display interpreter info")
	flag.BoolVar(&optimized, "o", false, "toggle interpreter optimization")

	flag.Parse()

	if len(flag.Args()) > 0 {
		file, err := os.ReadFile(flag.Arg(0))

		if err != nil {
			panic("file not found")
		} else {
			eval(string(file))
			return
		}
	}

	fmt.Println("brainrumi || made by snarkb0t on github.")
	fmt.Println("enter 'q', 'quit' or 'exit' to exit. enter 'r' to change display mode (encoded text / raw cell data)")
	fmt.Println()

	for {
		var line string

		fmt.Print(">>> ")
		fmt.Scan(&line)

		if line == "q" || line == "quit" || line == "exit" {
			return
		}

		if line == "r" {
			raw = !raw
		}

		eval(line)
	}
}
