// Generated from Funlang.g4 by ANTLR 4.7.

package generated

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 100,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 6, 14, 61, 10, 14, 13, 14, 14, 14, 62, 3, 14, 3, 14, 3, 15,
	3, 15, 5, 15, 69, 10, 15, 3, 15, 5, 15, 72, 10, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 7, 16, 80, 10, 16, 12, 16, 14, 16, 83, 11, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 94,
	10, 17, 12, 17, 14, 17, 97, 11, 17, 3, 17, 3, 17, 3, 81, 2, 18, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 3, 2, 6, 5, 2, 67, 92, 97, 97, 99,
	124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 104,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 37, 3, 2, 2, 2, 7, 39, 3, 2, 2, 2, 9, 41,
	3, 2, 2, 2, 11, 43, 3, 2, 2, 2, 13, 45, 3, 2, 2, 2, 15, 47, 3, 2, 2, 2,
	17, 49, 3, 2, 2, 2, 19, 51, 3, 2, 2, 2, 21, 53, 3, 2, 2, 2, 23, 55, 3,
	2, 2, 2, 25, 57, 3, 2, 2, 2, 27, 60, 3, 2, 2, 2, 29, 71, 3, 2, 2, 2, 31,
	75, 3, 2, 2, 2, 33, 89, 3, 2, 2, 2, 35, 36, 7, 92, 2, 2, 36, 4, 3, 2, 2,
	2, 37, 38, 7, 85, 2, 2, 38, 6, 3, 2, 2, 2, 39, 40, 7, 82, 2, 2, 40, 8,
	3, 2, 2, 2, 41, 42, 7, 63, 2, 2, 42, 10, 3, 2, 2, 2, 43, 44, 7, 126, 2,
	2, 44, 12, 3, 2, 2, 2, 45, 46, 7, 40, 2, 2, 46, 14, 3, 2, 2, 2, 47, 48,
	7, 46, 2, 2, 48, 16, 3, 2, 2, 2, 49, 50, 7, 61, 2, 2, 50, 18, 3, 2, 2,
	2, 51, 52, 7, 42, 2, 2, 52, 20, 3, 2, 2, 2, 53, 54, 7, 43, 2, 2, 54, 22,
	3, 2, 2, 2, 55, 56, 9, 2, 2, 2, 56, 24, 3, 2, 2, 2, 57, 58, 9, 3, 2, 2,
	58, 26, 3, 2, 2, 2, 59, 61, 9, 4, 2, 2, 60, 59, 3, 2, 2, 2, 61, 62, 3,
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64,
	65, 8, 14, 2, 2, 65, 28, 3, 2, 2, 2, 66, 68, 7, 15, 2, 2, 67, 69, 7, 12,
	2, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 72,
	7, 12, 2, 2, 71, 66, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2,
	73, 74, 8, 15, 2, 2, 74, 30, 3, 2, 2, 2, 75, 76, 7, 49, 2, 2, 76, 77, 7,
	44, 2, 2, 77, 81, 3, 2, 2, 2, 78, 80, 11, 2, 2, 2, 79, 78, 3, 2, 2, 2,
	80, 83, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 84, 3,
	2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 44, 2, 2, 85, 86, 7, 49, 2, 2,
	86, 87, 3, 2, 2, 2, 87, 88, 8, 16, 2, 2, 88, 32, 3, 2, 2, 2, 89, 90, 7,
	49, 2, 2, 90, 91, 7, 49, 2, 2, 91, 95, 3, 2, 2, 2, 92, 94, 10, 5, 2, 2,
	93, 92, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3,
	2, 2, 2, 96, 98, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 99, 8, 17, 2, 2, 99,
	34, 3, 2, 2, 2, 8, 2, 62, 68, 71, 81, 95, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'Z'", "'S'", "'P'", "'='", "'|'", "'&'", "','", "';'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "AssignmentOperator", "CompositionOperator", "PrimitiveRecursionOperator",
	"Comma", "Semicolon", "OpenParam", "CloseParam", "Nondigit", "Digit", "Whitespace",
	"Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "AssignmentOperator", "CompositionOperator", "PrimitiveRecursionOperator",
	"Comma", "Semicolon", "OpenParam", "CloseParam", "Nondigit", "Digit", "Whitespace",
	"Newline", "BlockComment", "LineComment",
}

type FunlangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewFunlangLexer(input antlr.CharStream) *FunlangLexer {

	l := new(FunlangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Funlang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FunlangLexer tokens.
const (
	FunlangLexerT__0                       = 1
	FunlangLexerT__1                       = 2
	FunlangLexerT__2                       = 3
	FunlangLexerAssignmentOperator         = 4
	FunlangLexerCompositionOperator        = 5
	FunlangLexerPrimitiveRecursionOperator = 6
	FunlangLexerComma                      = 7
	FunlangLexerSemicolon                  = 8
	FunlangLexerOpenParam                  = 9
	FunlangLexerCloseParam                 = 10
	FunlangLexerNondigit                   = 11
	FunlangLexerDigit                      = 12
	FunlangLexerWhitespace                 = 13
	FunlangLexerNewline                    = 14
	FunlangLexerBlockComment               = 15
	FunlangLexerLineComment                = 16
)
