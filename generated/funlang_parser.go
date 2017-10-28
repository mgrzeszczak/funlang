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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 142,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 3, 7, 3, 39, 10, 3, 12, 3, 14, 3, 42, 11, 3, 3, 4, 7, 4, 45,
	10, 4, 12, 4, 14, 4, 48, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 5, 6, 59, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 67, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 76, 10, 9,
	12, 9, 14, 9, 79, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 7, 10, 90, 10, 10, 12, 10, 14, 10, 93, 11, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 100, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 7, 12, 108, 10, 12, 12, 12, 14, 12, 111, 11, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 121, 10, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 17, 6, 17, 138, 10, 17, 13, 17, 14, 17, 139, 3, 17,
	2, 4, 18, 22, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 2, 3, 3, 2, 13, 14, 2, 138, 2, 34, 3, 2, 2, 2, 4, 40, 3, 2, 2, 2, 6,
	46, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 58, 3, 2, 2, 2, 12, 66, 3, 2, 2,
	2, 14, 68, 3, 2, 2, 2, 16, 73, 3, 2, 2, 2, 18, 80, 3, 2, 2, 2, 20, 99,
	3, 2, 2, 2, 22, 101, 3, 2, 2, 2, 24, 120, 3, 2, 2, 2, 26, 122, 3, 2, 2,
	2, 28, 127, 3, 2, 2, 2, 30, 129, 3, 2, 2, 2, 32, 137, 3, 2, 2, 2, 34, 35,
	5, 4, 3, 2, 35, 36, 5, 6, 4, 2, 36, 3, 3, 2, 2, 2, 37, 39, 5, 14, 8, 2,
	38, 37, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3,
	2, 2, 2, 41, 5, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 45, 5, 8, 5, 2, 44,
	43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2,
	2, 47, 7, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 5, 16, 9, 2, 50, 51,
	5, 12, 7, 2, 51, 52, 7, 3, 2, 2, 52, 9, 3, 2, 2, 2, 53, 59, 5, 32, 17,
	2, 54, 55, 5, 32, 17, 2, 55, 56, 7, 4, 2, 2, 56, 57, 5, 10, 6, 2, 57, 59,
	3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 58, 54, 3, 2, 2, 2, 59, 11, 3, 2, 2, 2,
	60, 61, 7, 11, 2, 2, 61, 67, 7, 12, 2, 2, 62, 63, 7, 11, 2, 2, 63, 64,
	5, 10, 6, 2, 64, 65, 7, 12, 2, 2, 65, 67, 3, 2, 2, 2, 66, 60, 3, 2, 2,
	2, 66, 62, 3, 2, 2, 2, 67, 13, 3, 2, 2, 2, 68, 69, 5, 16, 9, 2, 69, 70,
	7, 5, 2, 2, 70, 71, 5, 18, 10, 2, 71, 72, 7, 3, 2, 2, 72, 15, 3, 2, 2,
	2, 73, 77, 7, 13, 2, 2, 74, 76, 9, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 79,
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 17, 3, 2, 2, 2,
	79, 77, 3, 2, 2, 2, 80, 81, 8, 10, 1, 2, 81, 82, 5, 22, 12, 2, 82, 91,
	3, 2, 2, 2, 83, 84, 12, 3, 2, 2, 84, 85, 7, 6, 2, 2, 85, 86, 7, 11, 2,
	2, 86, 87, 5, 20, 11, 2, 87, 88, 7, 12, 2, 2, 88, 90, 3, 2, 2, 2, 89, 83,
	3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2,
	92, 19, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 100, 5, 18, 10, 2, 95, 96,
	5, 18, 10, 2, 96, 97, 7, 4, 2, 2, 97, 98, 5, 20, 11, 2, 98, 100, 3, 2,
	2, 2, 99, 94, 3, 2, 2, 2, 99, 95, 3, 2, 2, 2, 100, 21, 3, 2, 2, 2, 101,
	102, 8, 12, 1, 2, 102, 103, 5, 24, 13, 2, 103, 109, 3, 2, 2, 2, 104, 105,
	12, 4, 2, 2, 105, 106, 7, 7, 2, 2, 106, 108, 5, 24, 13, 2, 107, 104, 3,
	2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2,
	2, 110, 23, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 113, 7, 11, 2, 2, 113,
	114, 5, 18, 10, 2, 114, 115, 7, 12, 2, 2, 115, 121, 3, 2, 2, 2, 116, 121,
	5, 26, 14, 2, 117, 121, 5, 28, 15, 2, 118, 121, 5, 30, 16, 2, 119, 121,
	5, 16, 9, 2, 120, 112, 3, 2, 2, 2, 120, 116, 3, 2, 2, 2, 120, 117, 3, 2,
	2, 2, 120, 118, 3, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 25, 3, 2, 2, 2,
	122, 123, 7, 8, 2, 2, 123, 124, 7, 11, 2, 2, 124, 125, 5, 32, 17, 2, 125,
	126, 7, 12, 2, 2, 126, 27, 3, 2, 2, 2, 127, 128, 7, 9, 2, 2, 128, 29, 3,
	2, 2, 2, 129, 130, 7, 10, 2, 2, 130, 131, 7, 11, 2, 2, 131, 132, 5, 32,
	17, 2, 132, 133, 7, 4, 2, 2, 133, 134, 5, 32, 17, 2, 134, 135, 7, 12, 2,
	2, 135, 31, 3, 2, 2, 2, 136, 138, 7, 14, 2, 2, 137, 136, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 33, 3,
	2, 2, 2, 12, 40, 46, 58, 66, 77, 91, 99, 109, 120, 139,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "','", "'='", "'|'", "'&'", "'$Z'", "'$S'", "'$P'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "OpenParam", "CloseParam", "Nondigit",
	"Digit", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "definitions", "tasks", "task", "args", "parameters", "definition",
	"name", "function", "functions", "composition", "primitiveRecursion", "zero",
	"successor", "projection", "number",
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
	FunlangParserEOF          = antlr.TokenEOF
	FunlangParserT__0         = 1
	FunlangParserT__1         = 2
	FunlangParserT__2         = 3
	FunlangParserT__3         = 4
	FunlangParserT__4         = 5
	FunlangParserT__5         = 6
	FunlangParserT__6         = 7
	FunlangParserT__7         = 8
	FunlangParserOpenParam    = 9
	FunlangParserCloseParam   = 10
	FunlangParserNondigit     = 11
	FunlangParserDigit        = 12
	FunlangParserWhitespace   = 13
	FunlangParserNewline      = 14
	FunlangParserBlockComment = 15
	FunlangParserLineComment  = 16
)

// FunlangParser rules.
const (
	FunlangParserRULE_program            = 0
	FunlangParserRULE_definitions        = 1
	FunlangParserRULE_tasks              = 2
	FunlangParserRULE_task               = 3
	FunlangParserRULE_args               = 4
	FunlangParserRULE_parameters         = 5
	FunlangParserRULE_definition         = 6
	FunlangParserRULE_name               = 7
	FunlangParserRULE_function           = 8
	FunlangParserRULE_functions          = 9
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

func (s *ProgramContext) Definitions() IDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionsContext)
}

func (s *ProgramContext) Tasks() ITasksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasksContext)
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
		p.SetState(32)
		p.Definitions()
	}
	{
		p.SetState(33)
		p.Tasks()
	}

	return localctx
}

// IDefinitionsContext is an interface to support dynamic dispatch.
type IDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionsContext differentiates from other interfaces.
	IsDefinitionsContext()
}

type DefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionsContext() *DefinitionsContext {
	var p = new(DefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_definitions
	return p
}

func (*DefinitionsContext) IsDefinitionsContext() {}

func NewDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionsContext {
	var p = new(DefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_definitions

	return p
}

func (s *DefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionsContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DefinitionsContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitDefinitions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Definitions() (localctx IDefinitionsContext) {
	localctx = NewDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FunlangParserRULE_definitions)

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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(35)
				p.Definition()
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// ITasksContext is an interface to support dynamic dispatch.
type ITasksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasksContext differentiates from other interfaces.
	IsTasksContext()
}

type TasksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasksContext() *TasksContext {
	var p = new(TasksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_tasks
	return p
}

func (*TasksContext) IsTasksContext() {}

func NewTasksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TasksContext {
	var p = new(TasksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_tasks

	return p
}

func (s *TasksContext) GetParser() antlr.Parser { return s.parser }

func (s *TasksContext) AllTask() []ITaskContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITaskContext)(nil)).Elem())
	var tst = make([]ITaskContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITaskContext)
		}
	}

	return tst
}

func (s *TasksContext) Task(i int) ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *TasksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TasksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TasksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitTasks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Tasks() (localctx ITasksContext) {
	localctx = NewTasksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FunlangParserRULE_tasks)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FunlangParserNondigit {
		{
			p.SetState(41)
			p.Task()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITaskContext is an interface to support dynamic dispatch.
type ITaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaskContext differentiates from other interfaces.
	IsTaskContext()
}

type TaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaskContext() *TaskContext {
	var p = new(TaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TaskContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitTask(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Task() (localctx ITaskContext) {
	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FunlangParserRULE_task)

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
		p.SetState(47)
		p.Name()
	}
	{
		p.SetState(48)
		p.Parameters()
	}
	{
		p.SetState(49)
		p.Match(FunlangParserT__0)
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArgsContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FunlangParserRULE_args)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Number()
		}
		{
			p.SetState(53)
			p.Match(FunlangParserT__1)
		}
		{
			p.SetState(54)
			p.Args()
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

func (s *ParametersContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
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
	p.EnterRule(localctx, 10, FunlangParserRULE_parameters)

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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(59)
			p.Match(FunlangParserCloseParam)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(61)
			p.Args()
		}
		{
			p.SetState(62)
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

func (s *DefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
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
	p.EnterRule(localctx, 12, FunlangParserRULE_definition)

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
		p.SetState(66)
		p.Name()
	}
	{
		p.SetState(67)
		p.Match(FunlangParserT__2)
	}
	{
		p.SetState(68)
		p.function(0)
	}
	{
		p.SetState(69)
		p.Match(FunlangParserT__0)
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
	p.EnterRule(localctx, 14, FunlangParserRULE_name)
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
		p.SetState(71)
		p.Match(FunlangParserNondigit)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(72)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FunlangParserNondigit || _la == FunlangParserDigit) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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

func (s *FunctionContext) OpenParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserOpenParam, 0)
}

func (s *FunctionContext) Functions() IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *FunctionContext) CloseParam() antlr.TerminalNode {
	return s.GetToken(FunlangParserCloseParam, 0)
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, FunlangParserRULE_function, _p)

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
		p.SetState(79)
		p.composition(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunctionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FunlangParserRULE_function)
			p.SetState(81)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(82)
				p.Match(FunlangParserT__3)
			}
			{
				p.SetState(83)
				p.Match(FunlangParserOpenParam)
			}
			{
				p.SetState(84)
				p.Functions()
			}
			{
				p.SetState(85)
				p.Match(FunlangParserCloseParam)
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FunlangParserRULE_functions
	return p
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FunlangParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionsContext) Functions() IFunctionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FunlangVisitor:
		return t.VisitFunctions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FunlangParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FunlangParserRULE_functions)

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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.function(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.function(0)
		}
		{
			p.SetState(94)
			p.Match(FunlangParserT__1)
		}
		{
			p.SetState(95)
			p.Functions()
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
		p.SetState(100)
		p.PrimitiveRecursion()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCompositionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, FunlangParserRULE_composition)
			p.SetState(102)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(103)
				p.Match(FunlangParserT__4)
			}
			{
				p.SetState(104)
				p.PrimitiveRecursion()
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FunlangParserOpenParam:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(FunlangParserOpenParam)
		}
		{
			p.SetState(111)
			p.function(0)
		}
		{
			p.SetState(112)
			p.Match(FunlangParserCloseParam)
		}

	case FunlangParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Zero()
		}

	case FunlangParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Successor()
		}

	case FunlangParserT__7:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Projection()
		}

	case FunlangParserNondigit:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(117)
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
		p.SetState(120)
		p.Match(FunlangParserT__5)
	}
	{
		p.SetState(121)
		p.Match(FunlangParserOpenParam)
	}
	{
		p.SetState(122)
		p.Number()
	}
	{
		p.SetState(123)
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
		p.SetState(125)
		p.Match(FunlangParserT__6)
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
		p.SetState(127)
		p.Match(FunlangParserT__7)
	}
	{
		p.SetState(128)
		p.Match(FunlangParserOpenParam)
	}
	{
		p.SetState(129)
		p.Number()
	}
	{
		p.SetState(130)
		p.Match(FunlangParserT__1)
	}
	{
		p.SetState(131)
		p.Number()
	}
	{
		p.SetState(132)
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
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FunlangParserDigit {
		{
			p.SetState(134)
			p.Match(FunlangParserDigit)
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *FunlangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
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
