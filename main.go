package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/daniel-dailey/gint/interpreter"
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
		lexer := interpreter.NewLexer(s[0])
		parser := interpreter.NewParser(lexer)
		gint := interpreter.NewInterpreter(parser)
		gint.Interpret()
	}
}
