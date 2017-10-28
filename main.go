package main

import (
	"github.com/mgrzeszczak/funlang/interpreter"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mgrzeszczak/funlang/generated"
	"os"
)

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
	interpreter.NewFunlangInterpreter().Run(program)
}
