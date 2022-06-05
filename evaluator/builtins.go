package evaluator

import (
	"brackets/ast"
)

func GetBuiltin(operator string) func([]ast.Node, *map[string]ast.Node) ast.Node {
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
    if operator == "return" || operator == "ret" {
        return ret
    }
    if operator == "if" {
        return _if
    }
    if operator == "eval" {
        return eval
    }
    if operator == "head" {
        return head
    }
    if operator == "tail" {
        return tail
    }

    return nil
}

func def(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    var name string

    if t, ok := exprs[0].(ast.Identifier); ok {
        name = t.Literal
    }

    (*env)[name] = exprs[1]

    return ast.Number{Value: 0}
}

func _if(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    if getValueNumber(exprs[0], env) > 0 {
        return ast.Number{Value: getValueNumber(exprs[1], env)}
    } else {
        return ast.Number{Value: getValueNumber(exprs[2], env)}
    }
}

func ret(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    return ast.Number{Value: getValueNumber(exprs[0], env)}
}

func eval(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    if t, ok := exprs[0].(ast.Qexpr); ok {
        sexpr := ast.Sexpr{Operator: t.Arguments[0].String(), Arguments: t.Arguments[1:]}
        return ast.Number{Value: getValueNumber(evaluateSexpr(sexpr, env), env)}
    }

    return ast.Number{Value: 0}
}

func head(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    if t, ok := exprs[0].(ast.Qexpr); ok {
        return t.Arguments[1]
    }

    return ast.Number{Value: 0}
}

func tail(exprs []ast.Node, env *map[string]ast.Node) ast.Node {
    if t, ok := exprs[0].(ast.Qexpr); ok {
        return ast.Qexpr{Arguments: t.Arguments[2:]}
    }

    return ast.Number{Value: 0}
}
