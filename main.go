package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/daniel-dailey/gint/interpreter"
)

const debugProg = "PROGRAM Part10; VAR number : INTEGER; a, b, c, x : INTEGER; y : REAL; BEGIN {Part10} x := 11+4; x := 4+5; END."

func main() {
	reader := bufio.NewReader(os.Stdin)

	debug := flag.Bool("debug", false, "run debug prog")
	flag.Parse()

	if *debug {
		log.Println("run debug prog...", debugProg)
		lexer := interpreter.NewLexer(debugProg)
		parser := interpreter.NewParser(lexer)
		gint := interpreter.NewInterpreter(parser)
		gint.Interpret()
		return
	}

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
