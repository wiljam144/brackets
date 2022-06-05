package token

type Token struct {
    TokenType string
    Literal string
}

const (
    EOF = "EOF"
    ILLEGAL = "ILLEGAL"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    IDENT = "IDENT"
    FLOAT = "FLOAT"
)
