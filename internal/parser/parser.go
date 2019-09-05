package parser

import (
	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/internal/stxtree"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
)

type Parser interface {
	Parse([]byte, *stxtree.Graph) int
	GetErrors() []*errors.Error
	WithTokens() Parser
}

type AbstractParser struct {
	Lexer        scanner.Scanner
	CurrentToken *scanner.Token
	List         stackedNodeList
	Ast          *stxtree.Graph

	WithTokens bool
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

func (p *AbstractParser) Reset(src []byte, a *stxtree.Graph) {
	p.Lexer.Reset(src)
	p.List.Reset()

	a.FileData = src

	p.Ast = a
	p.Lexer.SetErrors(nil)
}

func (p *AbstractParser) getPosID(n graph.NodeID) graph.NodeID {
	if n == 0 {
		return 0
	}

	var posID graph.NodeID
	p.Ast.Foreach(n, func(edge graph.Edge, node graph.Node) bool {
		if node.Type == stxtree.NodeTypePosition {
			posID = edge.To
			return false
		}

		return true
	})

	return posID
}

func (p *AbstractParser) getListStartPosID(l []graph.NodeID) graph.NodeID {
	if len(l) > 0 {
		if l[0] == 0 {
			return 0
		}

		return p.getPosID(l[0])
	}

	return 0
}

func (p *AbstractParser) getListEndPosID(l []graph.NodeID) graph.NodeID {
	if len(l) > 0 {
		if l[len(l)-1] == 0 {
			return 0
		}

		return p.getPosID(l[len(l)-1])
	}

	return 0
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

func (p *AbstractParser) NewPosition(startNodeIDlist []graph.NodeID, tokenList []*scanner.Token, endNodeIdList []graph.NodeID) graph.NodeID {
	var pos ast.Position

	// Get start pos

	sPosID := p.getListStartPosID(startNodeIDlist)
	sTok := p.getStartToken(tokenList)

	if sPosID != 0 {
		sPos := p.Ast.GetPosition(sPosID)
		pos.PS = sPos.PS
		pos.LS = sPos.LS
	}

	if sPosID == 0 && sTok != nil {
		pos.PS = sTok.StartPos
		pos.LS = sTok.StartLine
	}

	// get end pos

	ePosID := p.getListEndPosID(endNodeIdList)
	eTok := p.getEndToken(tokenList)
	esPosID := p.getListEndPosID(endNodeIdList)

	if ePosID != 0 {
		ePos := p.Ast.GetPosition(ePosID)
		pos.PE = ePos.PE
		pos.LE = ePos.LE
	}

	if ePosID == 0 && eTok != nil {
		pos.PE = sTok.EndPos
		pos.LE = sTok.EndLine
	}

	if ePosID == 0 && eTok != nil && esPosID != 0 {
		ePos := p.Ast.GetPosition(esPosID)
		pos.PE = ePos.PE
		pos.LE = ePos.LE
	}

	return p.Ast.NewPosition(pos)
}

func (p *AbstractParser) Children(parentNodeID graph.NodeID, nodeGroup ast.NodeGroup, children ...graph.NodeID) {
	for _, childNodeID := range children {
		if childNodeID == 0 {
			continue
		}

		p.Ast.GetNode(childNodeID).Group = nodeGroup
		p.Ast.Link(parentNodeID, childNodeID, nil)
	}
}

func (p *AbstractParser) SavePosition(nodeID graph.NodeID, posID graph.NodeID) {
	p.Ast.Link(nodeID, posID, nil)
}

func (p *AbstractParser) AppendTokens(nodeID graph.NodeID, group ast.TokenGroup, tokens []scanner.Token) {
	if !p.WithTokens {
		return
	}

	for _, token := range tokens {
		tokenID := p.convertToken(token, group)
		p.Ast.Link(nodeID, tokenID, p.Ast.Append)
	}
}

func (p *AbstractParser) PrependTokens(nodeID graph.NodeID, group ast.TokenGroup, tokens []scanner.Token) {
	if !p.WithTokens {
		return
	}

	for i := len(tokens) - 1; i >= 0; i-- {
		tokenID := p.convertToken(tokens[i], group)
		p.Ast.Link(nodeID, tokenID, p.Ast.Prepend)
	}
}

func (p *AbstractParser) MoveStartTokens(src graph.NodeID, dst graph.NodeID) {
	if !p.WithTokens {
		return
	}

	// TODO: refactor or remove
	// list := p.Ast.RemoveEdges(src, func(e graph.Edge) bool {
	// 	if e.Type != graph.EdgeTypeToken {
	// 		return false
	// 	}

	// 	token := p.Ast.Tokens.Get(graph.TokenID(e.Target))

	// 	if token.Group != ast.TokenGroupStart {
	// 		return false
	// 	}

	// 	return true
	// })

	// node := &p.Ast.Nodes[dst-1]
	// node.Edges = p.Ast.AppendEdges(list, node.Edges)
}

func (p *AbstractParser) convertToken(token scanner.Token, group ast.TokenGroup) graph.NodeID {
	pos := ast.Position{
		PS: token.StartPos,
		PE: token.EndPos,
		LS: token.StartLine,
	}
	posID := p.Ast.NewPosition(pos)
	tokenID := p.Ast.NewToken(ast.Token{
		Type:  token.Type,
		Group: group,
	})

	p.Ast.Link(tokenID, posID, p.Ast.Append)

	return tokenID
}
