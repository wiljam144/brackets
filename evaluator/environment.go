package evaluator

import "brackets/ast"

func getValueNumber(node ast.Node, env *Environment) float64 {
    if t, ok := node.(ast.Identifier); ok {
        return getValueNumber((*env)[t.Literal], env)
    }
    if t, ok := node.(ast.Number); ok {
        return t.Value
    }
    if t, ok := node.(ast.Sexpr); ok {
        return getValueNumber(evaluateSexpr(t, env), env)
    }

    return 0
}
