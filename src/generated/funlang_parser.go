// Generated from Funlang.g4 by ANTLR 4.7.

package generated // Funlang
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 144,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 6, 2, 38, 10, 2, 13, 2, 14, 2, 39, 3, 2, 5, 2, 43, 10, 2, 3, 3,
	3, 3, 5, 3, 47, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 57, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 65, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 74, 10, 8, 12, 8, 14, 8, 77,
	11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 85, 10, 9, 12, 9, 14,
	9, 88, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 95, 10, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 102, 10, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 110, 10, 12, 12, 12, 14, 12, 113, 11, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 123, 10, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 6, 17, 140, 10, 17, 13, 17, 14, 17, 141,
	3, 17, 2, 4, 16, 22, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 2, 3, 3, 2, 14, 15, 2, 142, 2, 42, 3, 2, 2, 2, 4, 46, 3, 2,
	2, 2, 6, 48, 3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 64, 3, 2, 2, 2, 12, 66,
	3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 94, 3, 2, 2, 2,
	20, 101, 3, 2, 2, 2, 22, 103, 3, 2, 2, 2, 24, 122, 3, 2, 2, 2, 26, 124,
	3, 2, 2, 2, 28, 129, 3, 2, 2, 2, 30, 131, 3, 2, 2, 2, 32, 139, 3, 2, 2,
	2, 34, 35, 5, 4, 3, 2, 35, 36, 7, 11, 2, 2, 36, 38, 3, 2, 2, 2, 37, 34,
	3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2,
	40, 43, 3, 2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 37, 3, 2, 2, 2, 42, 41, 3,
	2, 2, 2, 43, 3, 3, 2, 2, 2, 44, 47, 5, 12, 7, 2, 45, 47, 5, 6, 4, 2, 46,
	44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47, 5, 3, 2, 2, 2, 48, 49, 5, 14, 8,
	2, 49, 50, 5, 10, 6, 2, 50, 7, 3, 2, 2, 2, 51, 57, 5, 32, 17, 2, 52, 53,
	5, 32, 17, 2, 53, 54, 7, 10, 2, 2, 54, 55, 5, 8, 5, 2, 55, 57, 3, 2, 2,
	2, 56, 51, 3, 2, 2, 2, 56, 52, 3, 2, 2, 2, 57, 9, 3, 2, 2, 2, 58, 59, 7,
	12, 2, 2, 59, 65, 7, 13, 2, 2, 60, 61, 7, 12, 2, 2, 61, 62, 5, 8, 5, 2,
	62, 63, 7, 13, 2, 2, 63, 65, 3, 2, 2, 2, 64, 58, 3, 2, 2, 2, 64, 60, 3,
	2, 2, 2, 65, 11, 3, 2, 2, 2, 66, 67, 7, 6, 2, 2, 67, 68, 5, 14, 8, 2, 68,
	69, 7, 7, 2, 2, 69, 70, 5, 16, 9, 2, 70, 13, 3, 2, 2, 2, 71, 75, 7, 14,
	2, 2, 72, 74, 9, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 15, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2,
	78, 79, 8, 9, 1, 2, 79, 80, 5, 22, 12, 2, 80, 86, 3, 2, 2, 2, 81, 82, 12,
	3, 2, 2, 82, 83, 7, 8, 2, 2, 83, 85, 5, 18, 10, 2, 84, 81, 3, 2, 2, 2,
	85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 17, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 90, 7, 12, 2, 2, 90, 91, 5, 20, 11, 2,
	91, 92, 7, 13, 2, 2, 92, 95, 3, 2, 2, 2, 93, 95, 5, 16, 9, 2, 94, 89, 3,
	2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 19, 3, 2, 2, 2, 96, 102, 5, 16, 9, 2,
	97, 98, 5, 16, 9, 2, 98, 99, 7, 10, 2, 2, 99, 100, 5, 20, 11, 2, 100, 102,
	3, 2, 2, 2, 101, 96, 3, 2, 2, 2, 101, 97, 3, 2, 2, 2, 102, 21, 3, 2, 2,
	2, 103, 104, 8, 12, 1, 2, 104, 105, 5, 24, 13, 2, 105, 111, 3, 2, 2, 2,
	106, 107, 12, 4, 2, 2, 107, 108, 7, 9, 2, 2, 108, 110, 5, 24, 13, 2, 109,
	106, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 23, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 115, 7, 12,
	2, 2, 115, 116, 5, 16, 9, 2, 116, 117, 7, 13, 2, 2, 117, 123, 3, 2, 2,
	2, 118, 123, 5, 26, 14, 2, 119, 123, 5, 28, 15, 2, 120, 123, 5, 30, 16,
	2, 121, 123, 5, 14, 8, 2, 122, 114, 3, 2, 2, 2, 122, 118, 3, 2, 2, 2, 122,
	119, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 121, 3, 2, 2, 2, 123, 25, 3,
	2, 2, 2, 124, 125, 7, 3, 2, 2, 125, 126, 7, 12, 2, 2, 126, 127, 5, 32,
	17, 2, 127, 128, 7, 13, 2, 2, 128, 27, 3, 2, 2, 2, 129, 130, 7, 4, 2, 2,
	130, 29, 3, 2, 2, 2, 131, 132, 7, 5, 2, 2, 132, 133, 7, 12, 2, 2, 133,
	134, 5, 32, 17, 2, 134, 135, 7, 10, 2, 2, 135, 136, 5, 32, 17, 2, 136,
	137, 7, 13, 2, 2, 137, 31, 3, 2, 2, 2, 138, 140, 7, 15, 2, 2, 139, 138,
	3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2,
	2, 2, 142, 33, 3, 2, 2, 2, 14, 39, 42, 46, 56, 64, 75, 86, 94, 101, 111,
	122, 141,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'Z'", "'S'", "'P'", "'def'", "'='", "'|'", "'&'", "','", "';'", "'('",
	"')'",
}
var symbolicNames = []string{
	"", "", "", "", "Define", "AssignmentOperator", "CompositionOperator",
	"PrimitiveRecursionOperator", "Comma", "Semicolon", "OpenParam", "CloseParam",
	"Nondigit", "Digit", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "statement", "execution", "parameterList", "parameters", "definition",
	"name", "function", "innerFunctions", "functionList", "composition", "primitiveRecursion",
	"zero", "successor", "projection", "number",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FunlangParser struct {
	*antlr.BaseParser
}

func NewFunlangParser(input antlr.TokenStream) *FunlangParser {
	this := new(FunlangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Funlang.g4"

	return this
}

// FunlangParser tokens.
const (
	FunlangParserEOF                        = antlr.TokenEOF
	FunlangParserT__0                       = 1
	FunlangParserT__1                       = 2
	FunlangParserT__2                       = 3
	FunlangParserDefine                     = 4
	FunlangParserAssignmentOperator         = 5
	FunlangParserCompositionOperator        = 6
	FunlangParserPrimitiveRecursionOperator = 7
	FunlangParserComma                      = 8
	FunlangParserSemicolon                  = 9
	FunlangParserOpenParam                  = 10
	FunlangParserCloseParam                 = 11
	FunlangParserNondigit                   = 12
	FunlangParserDigit                      = 13
	FunlangParserWhitespace                 = 14
	FunlangParserNewline                    = 15
	FunlangParserBlockComment               = 16
	FunlangParserLineComment                = 17
)

// FunlangParser rules.
const (
	FunlangParserRULE_program            = 0
	FunlangParserRULE_statement          = 1
	FunlangParserRULE_execution          = 2
	FunlangParserRULE_parameterList      = 3
	FunlangParserRULE_parameters         = 4
	FunlangParserRULE_definition         = 5
	FunlangParserRULE_name               = 6
	FunlangParserRULE_function           = 7
	FunlangParserRULE_innerFunctions     = 8
	FunlangParserRULE_functionList       = 9
	FunlangParserRULE_composition        = 10
	FunlangParserRULE_primitiveRecursion = 11
	FunlangParserRULE_zero               = 12
	FunlangParserRULE_successor          = 13
	FunlangParserRULE_projection         = 14
	FunlangParserRULE_number             = 15
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) AllSemicolon() []antlr.TerminalNode {
	return s.GetTokens(FunlangParserSemicolon)
}

func (s *ProgramContext) Semicolon(i int) antlr.TerminalNode {
	return s.GetToken(FunlangParserSemicolon, i)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FunlangParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FunlangParserDefine, FunlangParserNondigit:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FunlangParserDefine || _la == FunlangParserNondigit {
			{
				p.SetState(32)
				p.Statement()
			}
			{
				p.SetState(33)
				p.Match(FunlangParserSemicolon)
			}

			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case FunlangParserEOF:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Definition() IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *StatementContext) Execution() IExecutionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FunlangParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FunlangParserDefine:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Definition()
		}

	case FunlangParserNondigit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Execution()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecutionContext is an interface to support dynamic dispatch.
type IExecutionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutionContext differentiates from other interfaces.
	IsExecutionContext()
}

type ExecutionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutionContext() *ExecutionContext {
	var p = new(ExecutionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_execution
	return p
}

func (*ExecutionContext) IsExecutionContext() {}

func NewExecutionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutionContext {
	var p = new(ExecutionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_execution

	return p
}

func (s *ExecutionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExecutionContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ExecutionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitExecution(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Execution() (localctx IExecutionContext) {
	localctx = NewExecutionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FunlangParserRULE_execution)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Name()
	}
	{
		p.SetState(47)
		p.Parameters()
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ParameterListContext) Comma() antlr.TerminalNode {
	return s.GetToken(FunlangParserComma, 0)
}

func (s *ParameterListContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FunlangParserRULE_parameterList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Number()
		}
		{
			p.SetState(51)
			p.Match(FunlangParserComma)
		}
		{
			p.SetState(52)
			p.ParameterList()
		}

	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) OpenParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserOpenParam, 0)
}

func (s *ParametersContext) CloseParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserCloseParam, 0)
}

func (s *ParametersContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FunlangParserRULE_parameters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(57)
			p.Match(FunlangParserCloseParam)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(59)
			p.ParameterList()
		}
		{
			p.SetState(60)
			p.Match(FunlangParserCloseParam)
		}

	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Define() antlr.TerminalNode {
	return s.GetToken(FunlangParserDefine, 0)
}

func (s *DefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DefinitionContext) AssignmentOperator() antlr.TerminalNode {
	return s.GetToken(FunlangParserAssignmentOperator, 0)
}

func (s *DefinitionContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FunlangParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(FunlangParserDefine)
	}
	{
		p.SetState(65)
		p.Name()
	}
	{
		p.SetState(66)
		p.Match(FunlangParserAssignmentOperator)
	}
	{
		p.SetState(67)
		p.function(0)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) AllNondigit() []antlr.TerminalNode {
	return s.GetTokens(FunlangParserNondigit)
}

func (s *NameContext) Nondigit(i int) antlr.TerminalNode {
	return s.GetToken(FunlangParserNondigit, i)
}

func (s *NameContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(FunlangParserDigit)
}

func (s *NameContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(FunlangParserDigit, i)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FunlangParserRULE_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(FunlangParserNondigit)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(70)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FunlangParserNondigit || _la == FunlangParserDigit) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Composition() ICompositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositionContext)
}

func (s *FunctionContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionContext) CompositionOperator() antlr.TerminalNode {
	return s.GetToken(FunlangParserCompositionOperator, 0)
}

func (s *FunctionContext) InnerFunctions() IInnerFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInnerFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInnerFunctionsContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Function() (localctx IFunctionContext) {
	return p.function(0)
}

func (p *FunlangParser) function(_p int) (localctx IFunctionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFunctionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, FunlangParserRULE_function, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.composition(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunctionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FunlangParserRULE_function)
			p.SetState(79)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(80)
				p.Match(FunlangParserCompositionOperator)
			}
			{
				p.SetState(81)
				p.InnerFunctions()
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IInnerFunctionsContext is an interface to support dynamic dispatch.
type IInnerFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInnerFunctionsContext differentiates from other interfaces.
	IsInnerFunctionsContext()
}

type InnerFunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerFunctionsContext() *InnerFunctionsContext {
	var p = new(InnerFunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_innerFunctions
	return p
}

func (*InnerFunctionsContext) IsInnerFunctionsContext() {}

func NewInnerFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerFunctionsContext {
	var p = new(InnerFunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_innerFunctions

	return p
}

func (s *InnerFunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerFunctionsContext) OpenParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserOpenParam, 0)
}

func (s *InnerFunctionsContext) FunctionList() IFunctionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionListContext)
}

func (s *InnerFunctionsContext) CloseParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserCloseParam, 0)
}

func (s *InnerFunctionsContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *InnerFunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerFunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerFunctionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitInnerFunctions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) InnerFunctions() (localctx IInnerFunctionsContext) {
	localctx = NewInnerFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FunlangParserRULE_innerFunctions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(88)
			p.FunctionList()
		}
		{
			p.SetState(89)
			p.Match(FunlangParserCloseParam)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.function(0)
		}

	}

	return localctx
}

// IFunctionListContext is an interface to support dynamic dispatch.
type IFunctionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionListContext differentiates from other interfaces.
	IsFunctionListContext()
}

type FunctionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionListContext() *FunctionListContext {
	var p = new(FunctionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_functionList
	return p
}

func (*FunctionListContext) IsFunctionListContext() {}

func NewFunctionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionListContext {
	var p = new(FunctionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_functionList

	return p
}

func (s *FunctionListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionListContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionListContext) Comma() antlr.TerminalNode {
	return s.GetToken(FunlangParserComma, 0)
}

func (s *FunctionListContext) FunctionList() IFunctionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionListContext)
}

func (s *FunctionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitFunctionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) FunctionList() (localctx IFunctionListContext) {
	localctx = NewFunctionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FunlangParserRULE_functionList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.function(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.function(0)
		}
		{
			p.SetState(96)
			p.Match(FunlangParserComma)
		}
		{
			p.SetState(97)
			p.FunctionList()
		}

	}

	return localctx
}

// ICompositionContext is an interface to support dynamic dispatch.
type ICompositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompositionContext differentiates from other interfaces.
	IsCompositionContext()
}

type CompositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompositionContext() *CompositionContext {
	var p = new(CompositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_composition
	return p
}

func (*CompositionContext) IsCompositionContext() {}

func NewCompositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompositionContext {
	var p = new(CompositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_composition

	return p
}

func (s *CompositionContext) GetParser() antlr.Parser { return s.parser }

func (s *CompositionContext) PrimitiveRecursion() IPrimitiveRecursionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveRecursionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveRecursionContext)
}

func (s *CompositionContext) Composition() ICompositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompositionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompositionContext)
}

func (s *CompositionContext) PrimitiveRecursionOperator() antlr.TerminalNode {
	return s.GetToken(FunlangParserPrimitiveRecursionOperator, 0)
}

func (s *CompositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompositionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitComposition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Composition() (localctx ICompositionContext) {
	return p.composition(0)
}

func (p *FunlangParser) composition(_p int) (localctx ICompositionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCompositionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICompositionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, FunlangParserRULE_composition, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.PrimitiveRecursion()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCompositionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FunlangParserRULE_composition)
			p.SetState(104)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(105)
				p.Match(FunlangParserPrimitiveRecursionOperator)
			}
			{
				p.SetState(106)
				p.PrimitiveRecursion()
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitiveRecursionContext is an interface to support dynamic dispatch.
type IPrimitiveRecursionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveRecursionContext differentiates from other interfaces.
	IsPrimitiveRecursionContext()
}

type PrimitiveRecursionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveRecursionContext() *PrimitiveRecursionContext {
	var p = new(PrimitiveRecursionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_primitiveRecursion
	return p
}

func (*PrimitiveRecursionContext) IsPrimitiveRecursionContext() {}

func NewPrimitiveRecursionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveRecursionContext {
	var p = new(PrimitiveRecursionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_primitiveRecursion

	return p
}

func (s *PrimitiveRecursionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveRecursionContext) OpenParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserOpenParam, 0)
}

func (s *PrimitiveRecursionContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *PrimitiveRecursionContext) CloseParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserCloseParam, 0)
}

func (s *PrimitiveRecursionContext) Zero() IZeroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZeroContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZeroContext)
}

func (s *PrimitiveRecursionContext) Successor() ISuccessorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISuccessorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISuccessorContext)
}

func (s *PrimitiveRecursionContext) Projection() IProjectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *PrimitiveRecursionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *PrimitiveRecursionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveRecursionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveRecursionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitPrimitiveRecursion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) PrimitiveRecursion() (localctx IPrimitiveRecursionContext) {
	localctx = NewPrimitiveRecursionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FunlangParserRULE_primitiveRecursion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FunlangParserOpenParam:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(113)
			p.function(0)
		}
		{
			p.SetState(114)
			p.Match(FunlangParserCloseParam)
		}

	case FunlangParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Zero()
		}

	case FunlangParserT__1:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.Successor()
		}

	case FunlangParserT__2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)
			p.Projection()
		}

	case FunlangParserNondigit:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(119)
			p.Name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IZeroContext is an interface to support dynamic dispatch.
type IZeroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZeroContext differentiates from other interfaces.
	IsZeroContext()
}

type ZeroContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZeroContext() *ZeroContext {
	var p = new(ZeroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_zero
	return p
}

func (*ZeroContext) IsZeroContext() {}

func NewZeroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZeroContext {
	var p = new(ZeroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_zero

	return p
}

func (s *ZeroContext) GetParser() antlr.Parser { return s.parser }

func (s *ZeroContext) OpenParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserOpenParam, 0)
}

func (s *ZeroContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ZeroContext) CloseParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserCloseParam, 0)
}

func (s *ZeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZeroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZeroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitZero(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Zero() (localctx IZeroContext) {
	localctx = NewZeroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FunlangParserRULE_zero)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(FunlangParserT__0)
	}
	{
		p.SetState(123)
		p.Match(FunlangParserOpenParam)
	}
	{
		p.SetState(124)
		p.Number()
	}
	{
		p.SetState(125)
		p.Match(FunlangParserCloseParam)
	}

	return localctx
}

// ISuccessorContext is an interface to support dynamic dispatch.
type ISuccessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuccessorContext differentiates from other interfaces.
	IsSuccessorContext()
}

type SuccessorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuccessorContext() *SuccessorContext {
	var p = new(SuccessorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_successor
	return p
}

func (*SuccessorContext) IsSuccessorContext() {}

func NewSuccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuccessorContext {
	var p = new(SuccessorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_successor

	return p
}

func (s *SuccessorContext) GetParser() antlr.Parser { return s.parser }
func (s *SuccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuccessorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitSuccessor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Successor() (localctx ISuccessorContext) {
	localctx = NewSuccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FunlangParserRULE_successor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(FunlangParserT__1)
	}

	return localctx
}

// IProjectionContext is an interface to support dynamic dispatch.
type IProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionContext differentiates from other interfaces.
	IsProjectionContext()
}

type ProjectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionContext() *ProjectionContext {
	var p = new(ProjectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_projection
	return p
}

func (*ProjectionContext) IsProjectionContext() {}

func NewProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionContext {
	var p = new(ProjectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_projection

	return p
}

func (s *ProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionContext) OpenParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserOpenParam, 0)
}

func (s *ProjectionContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *ProjectionContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ProjectionContext) CloseParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserCloseParam, 0)
}

func (s *ProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitProjection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Projection() (localctx IProjectionContext) {
	localctx = NewProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FunlangParserRULE_projection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(FunlangParserT__2)
	}
	{
		p.SetState(130)
		p.Match(FunlangParserOpenParam)
	}
	{
		p.SetState(131)
		p.Number()
	}
	{
		p.SetState(132)
		p.Match(FunlangParserComma)
	}
	{
		p.SetState(133)
		p.Number()
	}
	{
		p.SetState(134)
		p.Match(FunlangParserCloseParam)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(FunlangParserDigit)
}

func (s *NumberContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(FunlangParserDigit, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FunlangParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FunlangParserDigit {
		{
			p.SetState(136)
			p.Match(FunlangParserDigit)
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *FunlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *FunctionContext = nil
		if localctx != nil {
			t = localctx.(*FunctionContext)
		}
		return p.Function_Sempred(t, predIndex)

	case 10:
		var t *CompositionContext = nil
		if localctx != nil {
			t = localctx.(*CompositionContext)
		}
		return p.Composition_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FunlangParser) Function_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FunlangParser) Composition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
