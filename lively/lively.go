package lively

import (
	"fmt"
	"github.com/lankaiyun/kaiyunchain/lively/ast"
	"github.com/lankaiyun/kaiyunchain/lively/kvm"
	"github.com/lankaiyun/kaiyunchain/lively/lexer"
	"github.com/lankaiyun/kaiyunchain/lively/parser"
	"github.com/lankaiyun/kaiyunchain/lively/token"
)

type Lively struct {
	Machine *kvm.KVM
}

func NewLively() *Lively {
	return &Lively{}
}

func (liv *Lively) ContractToTokens(contract string) ([]token.TokenS, error) {
	l := lexer.NewLexer(contract)
	var arr []token.TokenS
	for {
		tok := l.GetTokenS()
		arr = append(arr, tok)
		if tok.Token == token.ILLEGAL {
			return nil, fmt.Errorf("illegal character %s at %s", tok.Literal, tok.Position())
		}
		if tok.Token == token.EOF {
			break
		}
	}
	return arr, nil
}

func (liv *Lively) TokensToAst(tokenSs []token.TokenS) (*ast.Contract, error) {
	p := parser.NewParser(tokenSs)
	return p.Parse()
}
