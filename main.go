package main

import (
	"github.com/mgrzeszczak/funlang/interpreter"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mgrzeszczak/funlang/generated"
	"flag"
)

func main() {
	filename := flag.String("file","","Script to execute")
	flag.Parse()
	input, e := antlr.NewFileStream(*filename)
	if e != nil {
		panic(e)
	}
	lexer := generated.NewFunlangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := generated.NewFunlangParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	program := parser.Program().(*generated.ProgramContext)
	interpreter.NewFunlangInterpreter().Execute(program)
}
