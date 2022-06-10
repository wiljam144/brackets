package main

import (
    "os"
    "io/ioutil"
    "log"
    "brackets/parser"
    "brackets/lexer"
    "brackets/evaluator"
)

func main() {
    file := os.Args[1]

    content, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }

    input := string(content)
    l := lexer.New(input)
    p := parser.New(l)
    ast := p.ConstructAst()
    e := evaluator.New(ast)
    e.EvaluateProgram()
}
