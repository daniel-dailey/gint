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

func main() {
	debug := flag.Bool("debug", false, "run debug prog")
	fileName := flag.String("f", "", "name of pascal file to interpret")
	flag.Parse()

	if *debug {
		//Debug flag is true. When I have a better logger in place,
		// use this to set to debug level...
		return
	}

	if len(*fileName) > 0 {
		buf, err := os.ReadFile(*fileName)
		if err != nil {
			log.Panicln("error!", err.Error())
		}
		lexer := interpreter.NewLexer(string(buf))
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
