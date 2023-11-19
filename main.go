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

const debugProg = "PROGRAM Part10; VAR number : INTEGER; a, b, c, x : INTEGER; y : REAL; BEGIN {Part10} a:= 2*2*6*777+555555; b := 10 * a + 10 * number DIV 4; x := 1; y := 20.0 / 7.9; END."

func main() {
	debug := flag.Bool("debug", false, "run debug prog")
	flag.Parse()

	if *debug {
		log.Println("run debug prog...", debugProg)
		lexer := interpreter.NewLexer(debugProg)
		parser := interpreter.NewParser(lexer)
		interpreter.NewInterpreter(parser).Interpret(parser.Parse())
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s := strings.Split(line, "\n")
		cmd := s[0]
		switch cmd {
		case "help":
			fmt.Printf("work in progress go-based pascal interpreter! \n input program as:\n" +
				"\t-PROGRAM [name]; [VARS:TYPES]; [BLOCKS]; END.\n")
			continue
		}
		lexer := interpreter.NewLexer(s[0])
		parser := interpreter.NewParser(lexer)
		rootNode := parser.Parse()
		stBuilder := interpreter.InitSymbolTableBuilder()
		stBuilder.Visit(rootNode)
		interpreter.NewInterpreter(parser).Interpret(rootNode)

	}
}
