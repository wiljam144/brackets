package repl

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "brackets/parser"
    "brackets/lexer"
    "brackets/evaluator"
)

func Repl() {
    var env evaluator.Environment = make(evaluator.Environment)

    for true {
        fmt.Print("brackets $ ")
        in := bufio.NewReader(os.Stdin)
        input, _ := in.ReadString('\n')
        input = strings.TrimSpace(input)

        l := lexer.New(input)
        p := parser.New(l)
        ast := p.ConstructAst()
        e := evaluator.NewWithEnv(ast, env)
        for _, elem := range e.EvaluateProgram() {
            fmt.Printf("%s\n", elem.String())
        }
        env = e.GetEnv()
    }
}
