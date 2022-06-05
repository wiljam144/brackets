package evaluator

import "brackets/ast"

func getFromEnvironment(ident ast.Identifier, env *map[string]ast.Node) ast.Node {
    return (*env)[ident.Literal]
}

func getValueNumber(node ast.Node, env *map[string]ast.Node) float64 {
    if t, ok := node.(ast.Identifier); ok {
        return getValueNumber(getFromEnvironment(t, env), env)
    }
    if t, ok := node.(ast.Number); ok {
        return t.Value
    }
    if t, ok := node.(ast.Sexpr); ok {
        return getValueNumber(evaluateSexpr(t, env), env)
    }

    return 0
}
