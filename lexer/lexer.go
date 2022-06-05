package lexer

import (
    "brackets/token"
)

type Lexer struct {
    input string
    ch byte
    position int
    readPosition int
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
        case '(': tok = newToken(token.LPAREN, l.ch)
        case ')': tok = newToken(token.RPAREN, l.ch)
        case '{': tok = newToken(token.LBRACE, l.ch)
        case '}': tok = newToken(token.RBRACE, l.ch)
        case 0: tok = newToken(token.EOF, ' ')
        default:
            if isDigit(l.ch) {
                tok.TokenType = token.FLOAT
                tok.Literal = l.readNumber()

                l.goBackOneChar()
            } else if isLetter(l.ch) {
                tok.TokenType = token.IDENT
                tok.Literal = l.readIdentifier()

                l.goBackOneChar()
            } else {
                tok.TokenType = token.ILLEGAL
                tok.Literal = ""
            }
    }

    l.readChar()

    return tok
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
    position := l.position;
    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
        l.readChar()
    }
}

func (l *Lexer) goBackOneChar() {
    l.position -= 1
    l.readPosition -= 1
    l.ch = l.input[l.readPosition]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9' || ch == '.'
}

func newToken(tokenType string, literal byte) token.Token {
    return token.Token{TokenType: tokenType, Literal: string(literal)}
}
