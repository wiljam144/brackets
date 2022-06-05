package main

import (
    "fmt"
    "brackets/parser"
    "brackets/lexer"
    "brackets/evaluator"
)

func main() {
    input := "(add 2 2)"
    l := lexer.New(input)
    p := parser.New(l)
    ast := p.ConstructAst()
    e := evaluator.New(ast)

    for _, elem := range e.EvaluateProgram() {
        fmt.Printf("%s\n", elem.String())
    }
}
