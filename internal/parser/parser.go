package parser

import (
	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/position"
)

type Parser interface {
	Parse([]byte, *graph.Graph) int
	GetErrors() []*errors.Error
}

type AbstractParser struct {
	Lexer        scanner.Scanner
	CurrentToken *scanner.Token
	List         stackedNodeList
	Ast          *graph.Graph
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

func (p *AbstractParser) Reset(src []byte, a *graph.Graph) {
	p.Lexer.Reset(src)
	p.List.Reset()

	a.FileData = src

	p.Ast = a
	p.Lexer.SetErrors(nil)
}

func (p *AbstractParser) getPosID(n graph.NodeID) graph.PositionID {
	if n == 0 {
		return 0
	}

	node := p.Ast.Nodes.Get(n)
	var posID graph.PositionID
	p.Ast.EachEdge(node.Edges, func(e graph.Edge) bool {
		if e.Type != graph.EdgeTypePosition {
			return false
		}

		posID = graph.PositionID(e.Target)

		return true
	})

	return posID
}

func (p *AbstractParser) getListStartPosID(l []graph.NodeID) graph.PositionID {
	if len(l) > 0 {
		if l[0] == 0 {
			return 0
		}

		return p.getPosID(l[0])
	}

	return 0
}

func (p *AbstractParser) getListEndPosID(l []graph.NodeID) graph.PositionID {
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

func (p *AbstractParser) NewPosition(startNodeIDlist []graph.NodeID, tokenList []*scanner.Token, endNodeIdList []graph.NodeID) graph.PositionID {
	var pos ast.Position

	// Get start pos

	sPosID := p.getListStartPosID(startNodeIDlist)
	sTok := p.getStartToken(tokenList)

	if sPosID != 0 {
		sPos := p.Ast.Positions.Get(sPosID)
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
		ePos := p.Ast.Positions.Get(ePosID)
		pos.PE = ePos.PE
		pos.LE = ePos.LE
	}

	if ePosID == 0 && eTok != nil {
		pos.PE = sTok.EndPos
		pos.LE = sTok.EndLine
	}

	if ePosID == 0 && eTok != nil && esPosID != 0 {
		ePos := p.Ast.Positions.Get(esPosID)
		pos.PE = ePos.PE
		pos.LE = ePos.LE
	}

	return p.Ast.Positions.Create(pos)
}

func (p *AbstractParser) Children(prevNodeID graph.NodeID, parentNodeID graph.NodeID, nodeGroup ast.NodeGroup, children ...graph.NodeID) graph.NodeID {
	for _, childNodeID := range children {
		if childNodeID == 0 {
			continue
		}

		childNode := p.Ast.Nodes.Get(childNodeID)
		childNode.Group = nodeGroup
		p.Ast.Nodes.Save(childNodeID, childNode)

		p.Ast.Link(parentNodeID, graph.EdgeTypeNode, uint32(childNodeID))
	}

	return 0
}

func (p *AbstractParser) SavePosition(nodeID graph.NodeID, posID graph.PositionID) {
	p.Ast.Link(nodeID, graph.EdgeTypePosition, uint32(posID))
}

func (p *AbstractParser) AppendTokens(nodeID graph.NodeID, group ast.TokenGroup, tokens []scanner.Token) {
	for _, token := range tokens {
		tkn := p.convertToken(token)
		tkn.Group = group
		tokenID := p.Ast.Tokens.Create(tkn)

		p.Ast.Link(nodeID, graph.EdgeTypeToken, uint32(tokenID))

	}
}

func (p *AbstractParser) PrependTokens(nodeID graph.NodeID, group ast.TokenGroup, tokens []scanner.Token) {
	for i := len(tokens) - 1; i >= 0; i-- {
		tkn := p.convertToken(tokens[i])
		tkn.Group = group
		tokenID := p.Ast.Tokens.Create(tkn)

		edge := graph.Edge{
			Type:   graph.EdgeTypeToken,
			Target: uint32(tokenID),
		}
		edgeID := p.Ast.Edges.Put(edge)
		node := p.Ast.Nodes.Get(nodeID)

		node.Edges = p.Ast.AppendEdges(graph.EdgeList{edgeID, edgeID}, node.Edges)
		p.Ast.Nodes.Save(nodeID, node)
	}
}

func (p *AbstractParser) MoveStartTokens(src graph.NodeID, dst graph.NodeID) {
	list := p.Ast.RemoveEdges(src, func(e graph.Edge) bool {
		if e.Type != graph.EdgeTypeToken {
			return false
		}

		token := p.Ast.Tokens.Get(graph.TokenID(e.Target))

		if token.Group != ast.TokenGroupStart {
			return false
		}

		return true
	})

	dstNode := p.Ast.Nodes.Get(dst)
	dstNode.Edges = p.Ast.AppendEdges(list, dstNode.Edges)
	p.Ast.Nodes.Save(dst, dstNode)
}

func (p *AbstractParser) convertToken(token scanner.Token) graph.Token {
	pos := ast.Position{
		PS: token.StartPos,
		PE: token.EndPos,
		LS: token.StartLine,
	}
	posID := p.Ast.Positions.Create(pos)

	return graph.Token{
		Type: token.Type,
		Pos:  posID,
	}
}
