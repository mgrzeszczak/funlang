package interpreter

import (
	"fmt"
	"errors"
)

var (
	ERR_WRONG_FUNC_COUNT = errors.New("wrong amount of funs passed to compose")
	ERR_INVALID_DOMAINS  = errors.New("functions to compose have invalid domains")
	ERR_WRONG_ARG_COUNT  = errors.New("wrong amount of args passed to function")
	ERR_NEGATIVE_ARG     = errors.New("argument is negative")
	ERR_NEGATIVE_DOMAIN  = errors.New("domain must be >=0")
	ERR_RECURSE_DOMAINS  = errors.New("invalid recurse function domains")
)

type Function struct {
	domain int
	name   string
	body   func([]int) int
}

func (f *Function) Domain() int {
	return f.domain
}

func (f *Function) Name() string {
	return f.name
}

func (f *Function) Call(args []int) int {
	f.checkArgs(args)
	return f.body(args)
}

func (f *Function) checkArgs(args []int) {
	if len(args) != f.domain {
		panic(ERR_WRONG_ARG_COUNT)
	}
	for _, a := range args {
		f.checkArg(a)
	}
}

func (f *Function) checkArg(arg int) {
	if arg < 0 {
		panic(ERR_NEGATIVE_ARG)
	}
}

func checkDomain(domain int) {
	if domain < 0 {
		panic(ERR_NEGATIVE_DOMAIN)
	}
}

func NewFunction(name string, domain int, body func([]int) int) *Function {
	checkDomain(domain)
	return &Function{
		domain: domain,
		name:   name,
		body:   body,
	}
}

func (f *Function) Compose(funs []*Function) *Function {
	if len(funs) != f.domain {
		panic(ERR_WRONG_FUNC_COUNT)
	}
	var domain int
	if len(funs) == 0 {
		domain = 0
	} else {
		domain = funs[0].domain
		for _, fun := range funs {
			if domain != fun.domain {
				panic(ERR_INVALID_DOMAINS)
			}
		}
	}
	name := fmt.Sprintf("%v(", f.name)
	for i, f := range funs {
		if i != len(funs)-1 {
			name += fmt.Sprintf("%v,", f.name)
		} else {
			name += fmt.Sprintf("%v)", f.name)
		}
	}

	return NewFunction(
		name,
		domain,
		func(input []int) int {
			args := make([]int, f.domain)
			for i, fun := range funs {
				args[i] = fun.Call(input)
			}
			return f.Call(args)
		},
	)
}

func (g *Function) Recurse(h *Function) *Function {
	if h.domain != g.domain+2 {
		panic(ERR_RECURSE_DOMAINS)
	}
	domain := g.domain + 1
	return NewFunction(
		g.name+"_r_"+h.name,
		domain,
		func(args []int) int {
			hargs := make([]int, domain+1)
			n := args[domain-1]
			values := make([]int, n+1)
			i := 0
			for i <= n {
				var v int
				if i == 0 {
					v = g.Call(args[:g.domain])
				} else {
					copy(hargs[:g.domain], args[:g.domain])
					hargs[g.domain] = i-1
					hargs[g.domain+1] = values[i-1]
					v = h.Call(hargs)
				}
				values[i] = v
				i++
			}
			return values[n]
		},
	)
}

func Zero(domain int) *Function {
	return NewFunction(
		fmt.Sprintf("Zero%v", domain),
		domain,
		func(args []int) int {
			return 0;
		},
	)
}

func Successor() *Function {
	return NewFunction(
		"Successor",
		1,
		func(args []int) int {
			return args[0] + 1
		},
	)
}

func Projection(from, which int) *Function {
	return NewFunction(
		fmt.Sprintf("Projection(%v,%v)", from, which),
		from,
		func(args []int) int {
			return args[which]
		},
	)
}