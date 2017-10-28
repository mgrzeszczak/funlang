// Generated from Funlang.g4 by ANTLR 4.7.

package generated // Funlang
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFunlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFunlangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitDefinitions(ctx *DefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitTasks(ctx *TasksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitTask(ctx *TaskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitFunctions(ctx *FunctionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitComposition(ctx *CompositionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitPrimitiveRecursion(ctx *PrimitiveRecursionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitZero(ctx *ZeroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitSuccessor(ctx *SuccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitProjection(ctx *ProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFunlangVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}
