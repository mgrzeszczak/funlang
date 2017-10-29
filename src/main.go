package main

import (
	"github.com/mgrzeszczak/funlang/src/interpreter"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mgrzeszczak/funlang/src/generated"
	"os"
	"fmt"
)

type StdOutConsole struct {
}

func (c *StdOutConsole) Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func main() {
	input, e := antlr.NewFileStream(os.Args[1])

	if e != nil {
		panic(e)
	}
	lexer := generated.NewFunlangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := generated.NewFunlangParser(stream)
	parser.RemoveErrorListeners()
	parser.BuildParseTrees = true
	program := parser.Program().(*generated.ProgramContext)
	interpreter.NewFunlangInterpreter(new(StdOutConsole)).Run(program)
}
