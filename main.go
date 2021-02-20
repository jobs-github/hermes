package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"hermes/lexer"
	"hermes/token"
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%v\n", tok.String())
		}
	}
}

func main() {
	repl(os.Stdin, os.Stdout)
}
