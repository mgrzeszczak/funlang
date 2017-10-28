// Generated from Funlang.g4 by ANTLR 4.7.

package generated // Funlang
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FunlangParser.
type FunlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FunlangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by FunlangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by FunlangParser#execution.
	VisitExecution(ctx *ExecutionContext) interface{}

	// Visit a parse tree produced by FunlangParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by FunlangParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by FunlangParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by FunlangParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by FunlangParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by FunlangParser#innerFunctions.
	VisitInnerFunctions(ctx *InnerFunctionsContext) interface{}

	// Visit a parse tree produced by FunlangParser#functionList.
	VisitFunctionList(ctx *FunctionListContext) interface{}

	// Visit a parse tree produced by FunlangParser#composition.
	VisitComposition(ctx *CompositionContext) interface{}

	// Visit a parse tree produced by FunlangParser#primitiveRecursion.
	VisitPrimitiveRecursion(ctx *PrimitiveRecursionContext) interface{}

	// Visit a parse tree produced by FunlangParser#zero.
	VisitZero(ctx *ZeroContext) interface{}

	// Visit a parse tree produced by FunlangParser#successor.
	VisitSuccessor(ctx *SuccessorContext) interface{}

	// Visit a parse tree produced by FunlangParser#projection.
	VisitProjection(ctx *ProjectionContext) interface{}

	// Visit a parse tree produced by FunlangParser#number.
	VisitNumber(ctx *NumberContext) interface{}
}
