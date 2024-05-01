package parser

import (
	"fmt"

	"github.com/lankaiyun/kaiyunchain/lively/ast"
	"github.com/lankaiyun/kaiyunchain/lively/token"
)

type (
	ParsePrefixFunc  func() ast.Expr
	ParseInfixFunc   func(ast.Expr) ast.Expr
	ParsePostfixFunc func() ast.Expr
)

type Parser struct {
	PrevIndex         int
	CurrIndex         int
	NextIndex         int
	CurrTokenS        token.TokenS
	TokenSs           []token.TokenS
	ParsePrefixFuncs  map[token.Token]ParsePrefixFunc
	ParseInfixFuncs   map[token.Token]ParseInfixFunc
	ParsePostfixFuncs map[token.Token]ParsePostfixFunc
	Errors            []string
	IsInFunc          bool
}

func NewParser(tokenSs []token.TokenS) *Parser {
	p := &Parser{TokenSs: tokenSs}
	p.ReadTokenS()
	p.ReadTokenS()

	p.ParsePrefixFuncs = make(map[token.Token]ParsePrefixFunc)
	p.ParsePrefixFuncs[token.IDENT] = p.ParseIdent
	p.ParseInfixFuncs = make(map[token.Token]ParseInfixFunc)
	p.ParsePostfixFuncs = make(map[token.Token]ParsePostfixFunc)
	return p
}

func (p *Parser) ReadTokenS() {
	p.CurrTokenS = p.TokenSs[p.NextIndex]
	p.PrevIndex = p.CurrIndex
	p.CurrIndex = p.NextIndex
	p.NextIndex++
}

func (p *Parser) Parse() (*ast.Contract, error) {
	contract := &ast.Contract{}
	contract.Stmts = []ast.Stmt{}
	if p.TokenSs[p.PrevIndex].Token != token.CONTRACT || p.TokenSs[p.NextIndex].Token != token.LBRACE {
		return nil, fmt.Errorf("illegal statement around %s", p.CurrTokenS.Position())
	} else {
		contract.Name = p.CurrTokenS.Literal
	}
	p.ReadTokenS()
	p.ReadTokenS()
	for p.CurrTokenS.Token != token.EOF {
		stmt := p.ParseStmt()
		if stmt == nil {
			p.Errors = append(p.Errors, fmt.Sprintf("unexpected err around %s", p.CurrTokenS.Position()))
		}
		contract.Stmts = append(contract.Stmts, stmt)
	}
	if len(p.Errors) == 0 {
		return contract, nil
	} else {
		return nil, fmt.Errorf("%s", p.Errors[0])
	}
}

func (p *Parser) ParseStmt() ast.Stmt {
	switch p.CurrTokenS.Token {
	case token.FUNC:
		p.ReadTokenS() // -> FuncName
		funcAst := &ast.Func{Name: p.CurrTokenS.Literal}
		// expect `(`
		if p.TokenSs[p.NextIndex].Token != token.LPAREN {
			return nil
		}
		p.ReadTokenS() // -> (
		// get parameters
		parameters := make([]*ast.ExprParameter, 0)
		if p.TokenSs[p.NextIndex].Token != token.RPAREN {
			p.ReadTokenS()
			for p.CurrTokenS.Token != token.RPAREN {
				parameter := &ast.ExprParameter{Type: p.CurrTokenS.Literal, Name: p.TokenSs[p.NextIndex].Literal}
				parameters = append(parameters, parameter)
				p.ReadTokenS()
				p.ReadTokenS()
				if p.CurrTokenS.Token == token.COMMA {
					p.ReadTokenS()
				}
			}
		} else {
			p.ReadTokenS() // -> )
		}
		funcAst.Parameters = parameters
		// expect `(`
		if p.TokenSs[p.NextIndex].Token != token.LBRACE {
			return nil
		}
		p.ReadTokenS() // -> {
		block := &ast.ExprBlock{}
		block.Exprs = []ast.Expr{}
		if p.TokenSs[p.NextIndex].Token != token.RBRACE {
			p.ReadTokenS()
			p.ParseBlock(block)
		} else {
			p.ReadTokenS() // -> }
		}
		funcAst.Block = block
		p.ReadTokenS()
		return funcAst
	case token.IDENT:
		if token.IsInternalType(p.CurrTokenS) {
			ident := &ast.StmtIdent{Type: p.CurrTokenS.Literal, Name: p.TokenSs[p.NextIndex].Literal}
			p.ReadTokenS()
			if p.TokenSs[p.NextIndex].Token == token.SEMICOLON {
				p.ReadTokenS()
				p.ReadTokenS()
				return ident
			} else {
				// todo
				return nil
			}
		} else {
			return nil
		}
	default:
		return nil
	}
}

func (p *Parser) ParseBlock(block *ast.ExprBlock) {
	for p.CurrTokenS.Token != token.RBRACE {
		expr := p.ParseExpr()
		block.Exprs = append(block.Exprs, expr)
		p.ReadTokenS()
	}
}

func (p *Parser) ParseExpr() ast.Expr {
	switch p.CurrTokenS.Token {
	case token.RETURN:
		returnAst := &ast.ExprS{}
		p.ReadTokenS()
		returnAst.Expr = p.ParseExpr2(token.LowestPrec)
		return returnAst
	default:
		expr := &ast.ExprS{}
		expr.Expr = p.ParseExpr2(token.LowestPrec)
		for p.TokenSs[p.NextIndex].Token == token.SEMICOLON {
			p.ReadTokenS()
		}
		return expr
	}
}

func (p *Parser) ParseExpr2(precedence int) ast.Expr {
	postfix := p.ParsePostfixFuncs[p.CurrTokenS.Token]
	if postfix != nil {
		return postfix()
	}
	prefix := p.ParsePrefixFuncs[p.CurrTokenS.Token]
	if prefix == nil {
		msg := fmt.Sprintf("unexpected nil expression around %s", p.CurrTokenS.Position())
		p.Errors = append(p.Errors, msg)
		return nil
	}
	leftExpr := prefix()
	for p.TokenSs[p.NextIndex].Token != token.SEMICOLON && precedence < p.CurrTokenS.Token.Precedence() {
		infix := p.ParseInfixFuncs[p.TokenSs[p.NextIndex].Token]
		if infix == nil {
			msg := fmt.Sprintf("unexpected nil expression around %s", p.CurrTokenS.Position())
			p.Errors = append(p.Errors, msg)
			return nil
		}
		p.ReadTokenS()
		leftExpr = infix(leftExpr)

		// Look for errors
		if leftExpr == nil {
			msg := fmt.Sprintf("unexpected nil expression around %s", p.CurrTokenS.Position())
			p.Errors = append(p.Errors, msg)
			return nil
		}
	}
	return leftExpr
}

func (p *Parser) ParseIdent() ast.Expr {
	return &ast.ExprIdent{Name: p.CurrTokenS.Literal}
}
