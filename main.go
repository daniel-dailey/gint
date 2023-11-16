package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s := strings.Split(line, "\n")
		lexer := NewLexer(s[0])
		parser := NewParser(lexer)
		interpreter := NewInterpreter(parser)
		interpreter.Interpret()
	}
}
