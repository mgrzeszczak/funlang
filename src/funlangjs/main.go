package main

import (
	"github.com/mgrzeszczak/funlang/src/interpreter"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/mgrzeszczak/funlang/src/generated"
	"github.com/gopherjs/gopherjs/js"
	"fmt"
	"strings"
)

type WebConsole struct {
	log func(string)
}

func (c *WebConsole) Log(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	text = strings.Replace(text, "[", "(", 1)
	text = strings.Replace(text, "]", ")", 1)
	c.log(text)
}

func run(input string, log func(string)) {
	in := antlr.NewInputStream(input)
	lexer := generated.NewFunlangLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := generated.NewFunlangParser(stream)
	parser.RemoveErrorListeners()
	parser.BuildParseTrees = true
	program := parser.Program().(*generated.ProgramContext)
	interpreter.NewFunlangInterpreter(&WebConsole{log: log}).Run(program)
}

func main() {
	js.Global.Set("funlang", map[string]interface{}{
		"runScript": run,
	})
}
