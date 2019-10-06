package parser

import (
	"github.com/z7zmey/php-parser/internal/parser/nodestack"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/internal/tree"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
)

type Parser interface {
	Parse([]byte, *tree.Tree) int
	GetErrors() []*errors.Error
	WithTokens() Parser
}

type AbstractParser struct {
	FileData     []byte
	Lexer        scanner.Scanner
	CurrentToken *scanner.Token
	List         nodestack.NodeStack
	Ast          *tree.Tree

	WithTokens bool

	tokenBuf []ast.Token
	nodeBuf  []ast.Node
}

func (p *AbstractParser) Error(msg string) {
	pos := &position.Position{
		StartLine: p.CurrentToken.StartLine,
		EndLine:   p.CurrentToken.EndLine,
		StartPos:  p.CurrentToken.StartPos,
		EndPos:    p.CurrentToken.EndPos,
	}

	p.Lexer.AddError(errors.NewError(msg, pos))
}

// GetErrors returns errors list
func (p *AbstractParser) GetErrors() []*errors.Error {
	return p.Lexer.GetErrors()
}

func (p *AbstractParser) Reset(src []byte, a *tree.Tree) {
	p.Lexer.Reset(src)
	p.List.Reset()

	a.Reset()
	p.Ast = a
	p.FileData = src
	p.Lexer.SetErrors(nil)
}

func (p *AbstractParser) getPos(n ast.Node) ast.Position {
	return n.Position
}

func (p *AbstractParser) getListStartPos(l []ast.Node) ast.Position {
	var pos ast.Position

	if len(l) > 0 {
		pos = l[0].Position
	}

	return pos
}

func (p *AbstractParser) getListEndPos(l []ast.Node) ast.Position {
	var pos ast.Position

	if len(l) > 0 {
		pos = l[len(l)-1].Position
	}

	return pos
}

func (p *AbstractParser) getStartToken(l []*scanner.Token) *scanner.Token {
	if len(l) > 0 {
		return l[0]
	}

	return nil
}

func (p *AbstractParser) getEndToken(l []*scanner.Token) *scanner.Token {
	if len(l) > 0 {
		return l[len(l)-1]
	}

	return nil
}

func (p *AbstractParser) NewPosition(startNodeIDlist []ast.Node, tokenList []*scanner.Token, endNodeIdList []ast.Node) ast.Position {
	var pos ast.Position

	// Get start pos

	if len(startNodeIDlist) > 0 {
		sPos := p.getListStartPos(startNodeIDlist)
		pos.PS = sPos.PS
		pos.LS = sPos.LS
	} else if len(tokenList) > 0 {
		sTok := p.getStartToken(tokenList)
		pos.PS = sTok.StartPos
		pos.LS = sTok.StartLine
	} else if len(endNodeIdList) > 0 {
		sPos := p.getListStartPos(endNodeIdList)
		pos.PS = sPos.PS
		pos.LS = sPos.LS
	}

	// get end pos

	if len(endNodeIdList) > 0 {
		ePos := p.getListEndPos(endNodeIdList)
		pos.PE = ePos.PE
		pos.LE = ePos.LE
	} else if len(tokenList) > 0 {
		eTok := p.getEndToken(tokenList)
		pos.PE = eTok.EndPos
		pos.LE = eTok.EndLine
	} else {
		ePos := p.getListEndPos(startNodeIDlist)
		pos.PE = ePos.PE
		pos.LE = ePos.LE
	}

	return pos
}

func (p *AbstractParser) PopTokens() []ast.Token {
	res := p.tokenBuf
	p.tokenBuf = p.tokenBuf[:0]
	return res
}

func (p *AbstractParser) Nodes(nn ...ast.Node) []ast.Node {
	for _, v := range nn {
		p.nodeBuf = append(p.nodeBuf, v)
	}
	res := p.nodeBuf
	p.nodeBuf = p.nodeBuf[:0]
	return res
}

func (p *AbstractParser) PushTokens(group ast.TokenGroup, tokens []scanner.Token) {
	for _, v := range tokens {
		pos := ast.Position{
			PS: v.StartPos,
			PE: v.EndPos,
			LS: v.StartLine,
			LE: v.EndLine,
		}
		token := ast.Token{
			Type:     v.Type,
			Group:    group,
			Position: pos,
		}

		p.tokenBuf = append(p.tokenBuf, token)
	}
}

func (p *AbstractParser) PrependToken(token *scanner.Token) {
	p.PushTokens(ast.TokenGroupStart, token.HiddenTokens)
	p.PushTokens(ast.TokenGroupStart, []scanner.Token{*token})
}
