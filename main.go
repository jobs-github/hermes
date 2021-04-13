package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"Q/lexer"
	"Q/object"
	"Q/parser"
)

func repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnv()
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
		val, err := program.Eval(env)
		if nil != err {
			io.WriteString(out, err.Error())
			io.WriteString(out, "\n")
		} else {
			if val != nil {
				io.WriteString(out, val.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func main() {
	repl(os.Stdin, os.Stdout)
}
