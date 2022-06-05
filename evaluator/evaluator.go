package evaluator

import "brackets/ast"

type Environment map[string]ast.Node

type Evaluator struct {
    program ast.Program
    env Environment
}

func New(prog ast.Program) Evaluator {
    var environment = make(Environment)
    return Evaluator{program: prog, env: environment}
}

func (e *Evaluator) EvaluateProgram() []ast.Node {
    results := []ast.Node{}

    for _, expr := range e.program.Expressions {
        results = append(results, evaluateSexpr(expr, &e.env))
    }

    return results
}

func evaluateSexpr(expr ast.Sexpr, env *Environment) ast.Node {
    builtin := getBuiltin(expr.Operator)

    if builtin == nil {
        // get function from env
    }

    return builtin(expr.Arguments, env)
}
