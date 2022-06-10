package evaluator

import (
	"brackets/ast"
	"fmt"
)

func getBuiltin(operator string) func([]ast.Node, *Environment) ast.Node {
    if operator == "println" {
        return println
    }
    if operator == "fn" {
        return fn
    }
    if operator == "add" || operator == "+" {
        return add
    }
    if operator == "sub" || operator == "-" {
        return sub
    }
    if operator == "mul" || operator == "*" {
        return mul
    }
    if operator == "div" || operator == "/" {
        return div
    }
    if operator == "def" {
        return def
    }
    if operator == "if" {
        return _if
    }
    if operator == "ret" || operator == "return" {
        return ret
    }
    if operator == "eval" {
        return eval
    }
    if operator == "list" {
        return list
    }
    if operator == "head" {
        return head
    }
    if operator == "tail" {
        return tail
    }

    return nil
}

func println(exprs []ast.Node, env *Environment) ast.Node {
    for _, elem := range exprs {
        fmt.Printf("%s\n", elem)
    }

    return ast.Number{Value: 0}
}

func add(exprs []ast.Node, env *Environment) ast.Node {
    var result float64 = 0

    for _, elem := range exprs {
        result += getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}

func sub(exprs []ast.Node, env *Environment) ast.Node {
    var exp ast.Node
    var result float64
    exp, exprs = exprs[0], exprs[1:]

    result = getValueNumber(exp, env)

    for _, elem := range exprs {
        result -= getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}

func mul(exprs []ast.Node, env *Environment) ast.Node {
    var result float64 = 1

    for _, elem := range exprs {
        result *= getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}

func div(exprs []ast.Node, env *Environment) ast.Node {
    var exp ast.Node
    var result float64
    exp, exprs = exprs[0], exprs[1:]

    result = getValueNumber(exp, env)

    for _, elem := range exprs {
        result /= getValueNumber(elem, env)
    }

    return ast.Number{Value: result}
}

func def(exprs []ast.Node, env *Environment) ast.Node {
    var name string

    if t, ok := exprs[0].(ast.Identifier); ok {
        name = t.Literal
    }

    (*env)[name] = exprs[1]

    return ast.Number{Value: 0}
}

func _if(exprs []ast.Node, env *Environment) ast.Node {
    if getValueNumber(exprs[0], env) == 1 {
        if t, ok := exprs[1].(ast.Sexpr); ok {
            return evaluateSexpr(t, env)
        }
    } else {
        if t, ok := exprs[2].(ast.Sexpr); ok {
            return evaluateSexpr(t, env)
        }
    }

    return ast.Number{Value: 0}
}

func ret(exprs []ast.Node, env *Environment) ast.Node {
    return exprs[0]
}

func eval(exprs []ast.Node, env *Environment) ast.Node {
    if t, ok := exprs[0].(ast.Qexpr); ok {
        sexpr := ast.Sexpr{Operator: t.Arguments[0].String(), Arguments: t.Arguments[1:]}
        return evaluateSexpr(sexpr, env)
    }
    if t, ok := exprs[0].(ast.Sexpr); ok {
        node := evaluateSexpr(t, env)
        if qexpr, ok := node.(ast.Qexpr); ok {
            sexpr := ast.Sexpr{Operator: qexpr.Arguments[0].String(), Arguments: t.Arguments[1:]}
            return evaluateSexpr(sexpr, env)
        }
    }

    return ast.Number{Value: 0}
}

func list(exprs []ast.Node, env *Environment) ast.Node {
    var qexpr ast.Qexpr

    for _, elem := range exprs {
        qexpr.Arguments = append(qexpr.Arguments, elem)
    }

    return qexpr
}

func head(exprs []ast.Node, env *Environment) ast.Node {
    if t, ok := exprs[0].(ast.Qexpr); ok {
        return t.Arguments[0]
    }
    if t, ok := exprs[0].(ast.Sexpr); ok {
        node := evaluateSexpr(t, env)
        if qexpr, ok := node.(ast.Qexpr); ok {
            return qexpr.Arguments[0]
        }
    }

    return ast.Number{Value: 0}
}

func tail(exprs []ast.Node, env *Environment) ast.Node {
    if t, ok := exprs[0].(ast.Qexpr); ok {
        return ast.Qexpr{Arguments: t.Arguments[1:]}
    }
    if t, ok := exprs[0].(ast.Sexpr); ok {
        node := evaluateSexpr(t, env)
        if qexpr, ok := node.(ast.Qexpr); ok {
            return ast.Qexpr{Arguments: qexpr.Arguments[1:]}
        }
    }

    return ast.Number{Value: 0}
}

func fn(exprs []ast.Node, env *Environment) ast.Node {
    var function ast.Lambda

    function.Env = *env
    if t, ok := exprs[0].(ast.Qexpr); ok {
        for _, elem := range t.Arguments {
            if ident, ok := elem.(ast.Identifier); ok {
                function.Env[ident.Literal] = ast.Number{Value: 0}
                function.Arguments = append(function.Arguments, ident.Literal)
            }
        }
    }

    if t, ok := exprs[1].(ast.Sexpr); ok {
        function.Body = t
    }

    return function
}
