package evaluator

import "brackets/ast"

func add(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    var sum float64 = 0

    for _, elem := range exprs {
        sum += getValueNumber(elem, env)
    }

    return ast.Number{Value: sum}
}

func sub(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    var exp ast.Node
    var result float64
    exp, exprs = exprs[0], exprs[1:]

    result = getValueNumber(exp, env)

    for _, elem := range(exprs) {
        result -= getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}

func mul(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    var result float64 = 1

    for _, elem := range(exprs) {
        result *= getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}

func div(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    var exp ast.Node
    var result float64
    exp, exprs = exprs[0], exprs[1:]

    result = getValueNumber(exp, env)

    for _, elem := range(exprs) {
        result /= getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}
