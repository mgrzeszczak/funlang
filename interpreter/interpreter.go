package interpreter

import (
	"github.com/mgrzeszczak/funlang/generated"
	"fmt"
	"strconv"
)

type FunlangInterpreter struct {
	functions map[string]*Function
}

func NewFunlangInterpreter() *FunlangInterpreter {
	return &FunlangInterpreter{
		functions: make(map[string]*Function),
	}
}

func (i *FunlangInterpreter) Execute(ctx *generated.ProgramContext) {
	for _, child := range ctx.Definitions().GetChildren() {
		i.visitDefinition(child.(*generated.DefinitionContext))
	}

	for _, child := range ctx.Tasks().GetChildren() {
		task := child.(*generated.TaskContext)
		name := task.Name().GetText()
		input := []int{}
		args := task.Parameters().(*generated.ParametersContext).Args()
		for args != nil {
			argsctx := args.(*generated.ArgsContext)
			n, e := strconv.Atoi(argsctx.Number().GetText())
			if e != nil {
				panic(e)
			}
			input = append(input, n)
			args = argsctx.Args()
		}
		f, ok := i.functions[name];
		if !ok {
			panic(fmt.Sprintf("Function %v is not defined", name))
		}
		fmt.Printf("%v = %v\n", task.GetText(), f.Call(input))
	}
}

func (i *FunlangInterpreter) visitDefinition(d *generated.DefinitionContext) {
	fmt.Printf("Visiting definition: %v\n", d.GetText())
	name := d.Name().GetText()
	function := i.visitFunction(d.Function().(*generated.FunctionContext))
	i.functions[name] = function
}

func (i *FunlangInterpreter) visitFunction(f *generated.FunctionContext) *Function {
	fmt.Printf("Visiting function: %v\n", f.GetText())
	if f.Function() == nil {
		return i.visitComposition(f.Composition().(*generated.CompositionContext))
	} else {
		left := i.visitFunction(f.Function().(*generated.FunctionContext))
		funs := []*Function{}
		funsCtx := f.Functions()
		for funsCtx != nil {
			fsctx := funsCtx.(*generated.FunctionsContext)
			funs = append(funs, i.visitFunction(fsctx.Function().(*generated.FunctionContext)))
			funsCtx = fsctx.Functions()
		}
		return left.Compose(funs)
	}
}

func (i *FunlangInterpreter) visitComposition(f *generated.CompositionContext) *Function {
	fmt.Printf("Visiting composition: %v\n", f.GetText())
	right := i.visitPrimitiveRecursion(f.PrimitiveRecursion().(*generated.PrimitiveRecursionContext))
	if f.Composition() == nil {
		return right
	} else {
		left := i.visitComposition(f.Composition().(*generated.CompositionContext))
		fmt.Printf("g(%v), h(%v)\n", left.domain, right.domain)
		return left.Recurse(right)
	}
}

func (i *FunlangInterpreter) visitPrimitiveRecursion(f *generated.PrimitiveRecursionContext) *Function {
	fmt.Printf("Visiting primitiveRecursion: %v\n", f.GetText())
	if f.OpenParam() != nil {
		return i.visitFunction(f.Function().(*generated.FunctionContext))
	} else if f.Name() != nil {
		return i.functions[f.Name().GetText()]
	} else if f.Zero() != nil {
		return i.visitZero(f.Zero().(*generated.ZeroContext))
	} else if f.Projection() != nil {
		return i.visitProjection(f.Projection().(*generated.ProjectionContext))
	} else if f.Successor() != nil {
		return i.visitSuccessor(f.Successor().(*generated.SuccessorContext))
	}
	panic(fmt.Sprintf("invalid primitive recursion node: %v", f.GetText()))
}

func (i *FunlangInterpreter) visitZero(f *generated.ZeroContext) *Function {
	fmt.Printf("Visiting zero: %v\n", f.GetText())
	domain, e := strconv.Atoi(f.Number().GetText())
	if e != nil {
		panic(e)
	}
	return Zero(domain)
}

func (i *FunlangInterpreter) visitSuccessor(f *generated.SuccessorContext) *Function {
	fmt.Printf("Visiting successor: %v\n", f.GetText())
	return Successor()
}

func (i *FunlangInterpreter) visitProjection(f *generated.ProjectionContext) *Function {
	fmt.Printf("Visiting projection: %v\n", f.GetText())
	from, e := strconv.Atoi(f.Number(0).GetText())
	if e != nil {
		panic(e)
	}
	which, e := strconv.Atoi(f.Number(1).GetText())
	if e != nil {
		panic(e)
	}
	return Projection(from, which)
}