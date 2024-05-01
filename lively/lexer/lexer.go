package lexer

import (
	"github.com/lankaiyun/kaiyunchain/lively/token"
)

type Lexer struct {
	CurrIndex int
	NextIndex int
	CurrChar  rune
	Chars     []rune
	Line      int
	Column    int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{Chars: []rune(input), Line: 1}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.NextIndex >= len(l.Chars) {
		l.CurrChar = rune(0)
	} else {
		l.CurrChar = l.Chars[l.NextIndex]
	}
	l.CurrIndex = l.NextIndex
	l.NextIndex++
	if l.CurrChar == '\n' {
		l.Column = 0
		l.Line++
	} else {
		l.Column++
	}
}

func (l *Lexer) GetTokenS() token.TokenS {
	var tokS token.TokenS
	l.SkipWhitespace()
	if l.CurrChar == '/' && l.PeekChar() == '/' {
		l.SkipComment()
		return l.GetTokenS()
	}
	switch l.CurrChar {
	case '"':
		tokS.Line = l.Line
		tokS.Column = l.Column
		var temp string
		temp += string('"')
		l.ReadChar()
		for l.CurrChar != '"' {
			if l.CurrChar == rune(0) || l.CurrChar == '\n' {
				return token.TokenS{Token: token.ILLEGAL, Literal: "unterminated string", Line: tokS.Line, Column: tokS.Column}
			} else {
				temp += string(l.CurrChar)
				l.ReadChar()
			}
		}
		temp += string('"')
		tokS.Token = token.STRING
		tokS.Literal = temp
	case ';':
		tokS.Line = l.Line
		tokS.Column = l.Column
		tokS.Token = token.SEMICOLON
		tokS.Literal = string(l.CurrChar)
	case '(':
		tokS.Line = l.Line
		tokS.Column = l.Column
		tokS.Token = token.LPAREN
		tokS.Literal = string(l.CurrChar)
	case ')':
		tokS.Line = l.Line
		tokS.Column = l.Column
		tokS.Token = token.RPAREN
		tokS.Literal = string(l.CurrChar)
	case '{':
		tokS.Line = l.Line
		tokS.Column = l.Column
		tokS.Token = token.LBRACE
		tokS.Literal = string(l.CurrChar)
	case '}':
		tokS.Line = l.Line
		tokS.Column = l.Column
		tokS.Token = token.RBRACE
		tokS.Literal = string(l.CurrChar)
	case '=':
		tokS.Line = l.Line
		tokS.Column = l.Column
		if l.PeekChar() == '=' {
			l.ReadChar()
			tokS.Token = token.EQL
			tokS.Literal = string("==")
		} else {
			tokS.Token = token.ASSIGN
			tokS.Literal = string(l.CurrChar)
		}
	case rune(0):
		tokS.Line = l.Line
		tokS.Column = l.Column
		tokS.Token = token.EOF
		tokS.Literal = ""
	default:
		tokS.Line = l.Line
		tokS.Column = l.Column
		if IsDigit(l.CurrChar) {
			var temp string
			for IsDigit(l.CurrChar) {
				temp += string(l.CurrChar)
				l.ReadChar()
			}
			tokS.Token = token.INT
			tokS.Literal = temp
		} else if IsIdentLetter(l.CurrChar) {
			var temp string
			for IsIdentLetter(l.CurrChar) {
				temp += string(l.CurrChar)
				l.ReadChar()
			}
			tokS.Token = token.Lookup(temp)
			tokS.Literal = temp
		} else {
			tokS.Token = token.ILLEGAL
			tokS.Literal = string(l.CurrChar)
		}
		return tokS
	}
	l.ReadChar()
	return tokS
}

func (l *Lexer) PeekChar() rune {
	if l.NextIndex >= len(l.Chars) {
		return rune(0)
	}
	return l.Chars[l.NextIndex]
}

func (l *Lexer) SkipComment() {
	for l.CurrChar != '\n' && l.CurrChar != rune(0) {
		l.ReadChar()
	}
	l.SkipWhitespace()
}

func (l *Lexer) SkipWhitespace() {
	for IsWhitespace(l.CurrChar) {
		l.ReadChar()
	}
}

func IsWhitespace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func IsDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func IsLetter(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z'
}

func IsIdentLetter(char rune) bool {
	return char == '_' || IsLetter(char)
}
