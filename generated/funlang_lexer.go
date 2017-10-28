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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 106,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 6, 15, 67, 10, 15,
	13, 15, 14, 15, 68, 3, 15, 3, 15, 3, 16, 3, 16, 5, 16, 75, 10, 16, 3, 16,
	5, 16, 78, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 86,
	10, 17, 12, 17, 14, 17, 89, 11, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 100, 10, 18, 12, 18, 14, 18, 103, 11,
	18, 3, 18, 3, 18, 3, 87, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 3, 2, 6, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2,
	11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 110, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 3, 37, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2, 2, 2, 9, 43, 3, 2,
	2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 53,
	3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 59, 3, 2, 2, 2,
	25, 61, 3, 2, 2, 2, 27, 63, 3, 2, 2, 2, 29, 66, 3, 2, 2, 2, 31, 77, 3,
	2, 2, 2, 33, 81, 3, 2, 2, 2, 35, 95, 3, 2, 2, 2, 37, 38, 7, 92, 2, 2, 38,
	4, 3, 2, 2, 2, 39, 40, 7, 85, 2, 2, 40, 6, 3, 2, 2, 2, 41, 42, 7, 82, 2,
	2, 42, 8, 3, 2, 2, 2, 43, 44, 7, 102, 2, 2, 44, 45, 7, 103, 2, 2, 45, 46,
	7, 104, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 63, 2, 2, 48, 12, 3, 2, 2,
	2, 49, 50, 7, 126, 2, 2, 50, 14, 3, 2, 2, 2, 51, 52, 7, 40, 2, 2, 52, 16,
	3, 2, 2, 2, 53, 54, 7, 46, 2, 2, 54, 18, 3, 2, 2, 2, 55, 56, 7, 61, 2,
	2, 56, 20, 3, 2, 2, 2, 57, 58, 7, 42, 2, 2, 58, 22, 3, 2, 2, 2, 59, 60,
	7, 43, 2, 2, 60, 24, 3, 2, 2, 2, 61, 62, 9, 2, 2, 2, 62, 26, 3, 2, 2, 2,
	63, 64, 9, 3, 2, 2, 64, 28, 3, 2, 2, 2, 65, 67, 9, 4, 2, 2, 66, 65, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69,
	70, 3, 2, 2, 2, 70, 71, 8, 15, 2, 2, 71, 30, 3, 2, 2, 2, 72, 74, 7, 15,
	2, 2, 73, 75, 7, 12, 2, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75,
	78, 3, 2, 2, 2, 76, 78, 7, 12, 2, 2, 77, 72, 3, 2, 2, 2, 77, 76, 3, 2,
	2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 8, 16, 2, 2, 80, 32, 3, 2, 2, 2, 81,
	82, 7, 49, 2, 2, 82, 83, 7, 44, 2, 2, 83, 87, 3, 2, 2, 2, 84, 86, 11, 2,
	2, 2, 85, 84, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 87, 85,
	3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 91, 7, 44, 2, 2,
	91, 92, 7, 49, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 8, 17, 2, 2, 94, 34, 3,
	2, 2, 2, 95, 96, 7, 49, 2, 2, 96, 97, 7, 49, 2, 2, 97, 101, 3, 2, 2, 2,
	98, 100, 10, 5, 2, 2, 99, 98, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99,
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 101, 3, 2,
	2, 2, 104, 105, 8, 18, 2, 2, 105, 36, 3, 2, 2, 2, 8, 2, 68, 74, 77, 87,
	101, 3, 8, 2, 2,
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
	"", "'Z'", "'S'", "'P'", "'def'", "'='", "'|'", "'&'", "','", "';'", "'('",
	"')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "Define", "AssignmentOperator", "CompositionOperator",
	"PrimitiveRecursionOperator", "Comma", "Semicolon", "OpenParam", "CloseParam",
	"Nondigit", "Digit", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "Define", "AssignmentOperator", "CompositionOperator",
	"PrimitiveRecursionOperator", "Comma", "Semicolon", "OpenParam", "CloseParam",
	"Nondigit", "Digit", "Whitespace", "Newline", "BlockComment", "LineComment",
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
	FunlangLexerDefine                     = 4
	FunlangLexerAssignmentOperator         = 5
	FunlangLexerCompositionOperator        = 6
	FunlangLexerPrimitiveRecursionOperator = 7
	FunlangLexerComma                      = 8
	FunlangLexerSemicolon                  = 9
	FunlangLexerOpenParam                  = 10
	FunlangLexerCloseParam                 = 11
	FunlangLexerNondigit                   = 12
	FunlangLexerDigit                      = 13
	FunlangLexerWhitespace                 = 14
	FunlangLexerNewline                    = 15
	FunlangLexerBlockComment               = 16
	FunlangLexerLineComment                = 17
)
