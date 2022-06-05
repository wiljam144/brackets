package evaluator

import (
	"brackets/ast"
	"fmt"
)

type Evaluator struct {
    program ast.Program
    env map[string]ast.Node
}

func New(prog ast.Program) Evaluator {
    var environment = make(map[string]ast.Node)
    return Evaluator{program: prog, env: environment}
}

func (e *Evaluator) EvaluateProgram() []ast.Node {
    results := []ast.Node{}

    for _, expr := range e.program.Expressions {
        results = append(results, evaluateSexpr(expr, &e.env))
    }

    return results
}

func evaluateSexpr(expr ast.Sexpr, env *map[string]ast.Node) ast.Node {
    builtin := GetBuiltin(expr.Operator)

    if builtin == nil {
        fmt.Printf("not a builtin")
    }

    return builtin(expr.Arguments, env)
}
