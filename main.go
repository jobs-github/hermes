package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"hermes/lexer"
	"hermes/parser"
)

func repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p, err := parser.New(l)
		if nil != err {
			io.WriteString(out, fmt.Sprintf("\t%v\n", err))
			continue
		}

		program := p.ParseProgram()
		errs := p.Errors()
		if len(errs) != 0 {
			for _, msg := range errs {
				io.WriteString(out, fmt.Sprintf("\t%v\n", msg))
			}
			continue
		}
		val := program.Eval()
		if val != nil {
			io.WriteString(out, val.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func main() {
	repl(os.Stdin, os.Stdout)
}
