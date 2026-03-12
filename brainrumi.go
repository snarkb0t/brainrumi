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
	raw   bool
	debug bool
)

func eval(code string, cells *[]int32, p *int) {
	var (
		scopes map[string]int = make(map[string]int, 0)

		open []int = make([]int, 0)

		tick time.Time = time.Now()
	)

	for s := 0; s < len(code); s++ {
		switch code[s] {
		case '[':
			open = append(open, s)
		case ']':
			scopes[strconv.Itoa(open[len(open)-1])] = s
			scopes[strconv.Itoa(s)] = open[len(open)-1]
			open = open[:len(open)-1]
		}
	}

	open = nil

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			if *p += 1; *p > len(*cells)-1 {
				*cells = append(*cells, 0)
			}
		case '<':
			if *p == 0 {
				continue
			}

			*p -= 1
		case '+':
			(*cells)[*p] += 1
		case '-':
			(*cells)[*p] -= 1
		case '.':
			if raw {
				fmt.Print((*cells)[*p])
			} else {
				fmt.Print(string((*cells)[*p]))
			}
		case ',':
			var putInto string
			fmt.Scanln(&putInto)
			(*cells)[*p] = int32([]byte(putInto)[0])
		case '[':
			if (*cells)[*p] == 0 {
				i = scopes[strconv.Itoa(i)]
			}
		case ']':
			if (*cells)[*p] != 0 {
				i = scopes[strconv.Itoa(i)]
			}
		}
	}

	if debug {
		fmt.Println("------")
		fmt.Println("executed in", time.Since(tick))
	}
}

func main() {
	var (
		cells []int32 = make([]int32, 1)
		p     int     = 0
	)

	flag.BoolVar(&raw, "r", false, "display raw cell data instead of encoded text")
	flag.BoolVar(&debug, "d", false, "display interpreter info")

	flag.Parse()

	if len(flag.Args()) > 0 {

		file, err := os.ReadFile(flag.Arg(0))

		if err != nil {
			panic("file not found")
		} else {
			eval(string(file), &cells, &p)
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

		eval(line, &cells, &p)
	}
}
