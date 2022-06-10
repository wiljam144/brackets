package ast

import (
	"fmt"
)

type Program struct {
    Expressions []Sexpr
}

type Node interface {
    String() string
}

type Number struct {
    Value float64
}

func (n Number) String() string {
    return fmt.Sprintf("%f", n.Value)
}

type Identifier struct {
    Literal string
}

func (i Identifier) String() string {
    return fmt.Sprintf("%s", i.Literal)
}

type Sexpr struct {
    Operator string
    Arguments []Node
}

func (s Sexpr) String() string {
    output := "("
    output += s.Operator
    output += " "
    for _, elem := range s.Arguments {
        output += elem.String()
        output += " "
    }
    output += ")"
    return output
}

type Qexpr struct {
    Arguments []Node
}

func (q Qexpr) String() string {
    output := "{"
    for _, elem := range q.Arguments {
        output += elem.String()
        output += " "
    }
    output += "}"
    return output
}

type Lambda struct {
    Body Sexpr
    Arguments []string
    Env map[string]Node
}

func (l Lambda) String() string {
    output := "(fn ("

    for _, elem := range l.Arguments {
        output += elem
        output += " "
    }
    output += ") " + l.Body.String() + ")"

    return output
}
