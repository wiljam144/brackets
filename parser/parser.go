package parser

import (
    "strconv"
    "brackets/ast"
    "brackets/token"
    "brackets/lexer"
)

type Parser struct {
    l *lexer.Lexer
}

func New(lex *lexer.Lexer) Parser {
    return Parser{l: lex}
}

func (p *Parser) ConstructAst() ast.Program {
    var program ast.Program;

    for true {
        tok := p.l.NextToken()
        if tok.TokenType == token.EOF {
            break
        }

        if tok.TokenType == token.LPAREN {
            program.Expressions = append(program.Expressions, p.parseSexpr())
        }
    }

    return program
}

func (p *Parser) parseSexpr() ast.Sexpr {
    var sexpr ast.Sexpr
    sexpr.Operator = p.l.NextToken().Literal

    for true {
        tok := p.l.NextToken()

        var nd ast.Node

        if tok.TokenType == token.RPAREN {
            break
        }

        if tok.TokenType == token.LPAREN {
            nd = p.parseSexpr()
        }

        if tok.TokenType == token.LBRACE {
            nd = p.parseQexpr()
        }

        if tok.TokenType == token.FLOAT {
            if s, err := strconv.ParseFloat(tok.Literal, 64); err == nil {
                nd = ast.Number{Value: s}
            }
        }

        if tok.TokenType == token.IDENT {
            nd = ast.Identifier{Literal: tok.Literal}
        }

        sexpr.Arguments = append(sexpr.Arguments, nd)
    }

    return sexpr
}

func (p *Parser) parseQexpr() ast.Qexpr {
    var qexpr ast.Qexpr

    for true {
        var nd ast.Node

        tok := p.l.NextToken()

        if tok.TokenType == token.RBRACE {
            break
        }

        if tok.TokenType == token.LPAREN {
            nd = p.parseSexpr()
        }

        if tok.TokenType == token.LBRACE {
            nd = p.parseQexpr()
        }

        if tok.TokenType == token.FLOAT {
            if s, err := strconv.ParseFloat(tok.Literal, 64); err == nil {
                nd = ast.Number{Value: s}
            }
        }

        if tok.TokenType == token.IDENT {
            nd = ast.Identifier{Literal: tok.Literal}
        }

        qexpr.Arguments = append(qexpr.Arguments, nd)
    }

    return qexpr
}
