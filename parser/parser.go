package parser

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/graph"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

type lastEdgeCache struct {
	nodeID graph.NodeID
	edgeID graph.EdgeID
}

type Parser interface {
	Parse([]byte, *graph.AST) int
	GetErrors() []*errors.Error
}

type AbstractParser struct {
	Lexer         scanner.Scanner
	CurrentToken  *scanner.Token
	List          stackedNodeList
	Ast           *graph.AST
	lastEdgeCache lastEdgeCache
}

func (l *AbstractParser) Error(msg string) {
	pos := &position.Position{
		StartLine: l.CurrentToken.StartLine,
		EndLine:   l.CurrentToken.EndLine,
		StartPos:  l.CurrentToken.StartPos,
		EndPos:    l.CurrentToken.EndPos,
	}

	l.Lexer.AddError(errors.NewError(msg, pos))
}

// GetErrors returns errors list
func (p *AbstractParser) GetErrors() []*errors.Error {
	return p.Lexer.GetErrors()
}

func (p *AbstractParser) Reset(src []byte, a *graph.AST) {
	p.Lexer.Reset(src)
	p.List.Reset()
	p.lastEdgeCache = lastEdgeCache{}

	a.FileData = src

	p.Ast = a
	p.Lexer.SetErrors(nil)
}

func (p *AbstractParser) getPosID(n graph.NodeID) graph.PositionID {
	if n == 0 {
		return 0
	}

	e := p.Ast.Nodes.Get(n).Edge
	posEdges := p.Ast.Edges.Get(e, graph.EdgeTypePosition)
	if len(posEdges) == 0 {
		return 0
	}

	return graph.PositionID(posEdges[0].Target)
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
	lastEdgeID := p.lastEdge(parentNodeID)

	for _, childNodeID := range children {
		if childNodeID == 0 {
			continue
		}

		childNode := p.Ast.Nodes.Get(childNodeID)
		childNode.Group = nodeGroup
		p.Ast.Nodes.Save(childNodeID, childNode)

		edge := graph.Edge{
			Type:   graph.EdgeTypeNode,
			Target: uint32(childNodeID),
		}
		edgeID := p.Ast.Edges.Put(edge)

		if lastEdgeID == 0 {
			node := p.Ast.Nodes.Get(parentNodeID)
			node.Edge = edgeID
			p.Ast.Nodes.Save(parentNodeID, node)
		} else {
			lastEdge := p.Ast.Edges.GetOne(lastEdgeID)
			lastEdge.Next = edgeID
			p.Ast.Edges.Set(lastEdgeID, lastEdge)
		}

		lastEdgeID = edgeID
		prevNodeID = childNodeID
	}

	p.lastEdgeCache = lastEdgeCache{
		nodeID: parentNodeID,
		edgeID: lastEdgeID,
	}

	return prevNodeID
}

func (p *AbstractParser) SavePosition(nodeID graph.NodeID, posID graph.PositionID) {
	edge := graph.Edge{
		Type:   graph.EdgeTypePosition,
		Target: uint32(posID),
	}
	edgeID := p.Ast.Edges.Put(edge)

	lastEdgeID := p.lastEdge(nodeID)
	if lastEdgeID == 0 {
		node := p.Ast.Nodes.Get(nodeID)
		node.Edge = edgeID
		p.Ast.Nodes.Save(nodeID, node)
	} else {
		lastEdge := p.Ast.Edges.GetOne(lastEdgeID)
		lastEdge.Next = edgeID
		p.Ast.Edges.Set(lastEdgeID, lastEdge)
	}
}

func (p *AbstractParser) lastEdge(nodeID graph.NodeID) graph.EdgeID {
	if p.lastEdgeCache.nodeID == nodeID {
		return p.lastEdgeCache.edgeID
	}

	node := p.Ast.Nodes.Get(nodeID)
	edgeID := node.Edge

	return p.Ast.Edges.GetLastID(edgeID)
}

func (p *AbstractParser) AppendTokens(nodeID graph.NodeID, group ast.TokenGroup, ffStrs []scanner.Token) {
	lastEdgeID := p.lastEdge(nodeID)

	for _, str := range ffStrs {
		tkn := p.convertToken(str)
		tkn.Group = group
		tokenID := p.Ast.Tokens.Create(tkn)

		edge := graph.Edge{
			Type:   graph.EdgeTypeToken,
			Target: uint32(tokenID),
		}
		edgeID := p.Ast.Edges.Put(edge)

		if lastEdgeID == 0 {
			node := p.Ast.Nodes.Get(nodeID)
			node.Edge = edgeID
			p.Ast.Nodes.Save(nodeID, node)
		} else {
			lastEdge := p.Ast.Edges.GetOne(lastEdgeID)
			lastEdge.Next = edgeID
			p.Ast.Edges.Set(lastEdgeID, lastEdge)
		}

		lastEdgeID = edgeID
	}

	p.lastEdgeCache = lastEdgeCache{
		nodeID: nodeID,
		edgeID: lastEdgeID,
	}
}

func (p *AbstractParser) PrependTokens(nodeID graph.NodeID, group ast.TokenGroup, ffStrs []scanner.Token) {
	bufEdgeID := p.Ast.Nodes.Get(nodeID).Edge

	var lastEdgeID graph.EdgeID
	for _, str := range ffStrs {
		tkn := p.convertToken(str)
		tkn.Group = group
		tokenID := p.Ast.Tokens.Create(tkn)

		edge := graph.Edge{
			Type:   graph.EdgeTypeToken,
			Target: uint32(tokenID),
		}
		edgeID := p.Ast.Edges.Put(edge)

		if lastEdgeID == 0 {
			node := p.Ast.Nodes.Get(nodeID)
			node.Edge = edgeID
			p.Ast.Nodes.Save(nodeID, node)
		} else {
			lastEdge := p.Ast.Edges.GetOne(lastEdgeID)
			lastEdge.Next = edgeID
			p.Ast.Edges.Set(lastEdgeID, lastEdge)
		}

		lastEdgeID = edgeID
	}

	if lastEdgeID != 0 {
		lastEdge := p.Ast.Edges.GetOne(lastEdgeID)
		lastEdge.Next = bufEdgeID
		p.Ast.Edges.Set(lastEdgeID, lastEdge)
	}
}

// [] => (prev) => (start) => ()
func (p *AbstractParser) MoveStartTokens(src graph.NodeID, dst graph.NodeID) {
	srcNode := p.Ast.Nodes.Get(src)

	if srcNode.Edge == 0 {
		return
	}

	edgeID := srcNode.Edge

	var prevEdgeID graph.EdgeID
	for {
		edge := p.Ast.Edges.GetOne(edgeID)
		if edge.Type == graph.EdgeTypeToken {
			token := p.Ast.Tokens.Get(graph.TokenID(edge.Target))
			if token.Group == ast.TokenGroupStart {
				break
			}
		}

		if edge.Next == 0 {
			return
		}

		prevEdgeID = edgeID
		edgeID = edge.Next
	}

	startEdgeID := edgeID
	endEdgeID := startEdgeID

	for {
		edge := p.Ast.Edges.GetOne(edgeID)
		if edge.Type != graph.EdgeTypeToken {
			break
		}

		token := p.Ast.Tokens.Get(graph.TokenID(edge.Target))
		if token.Group != ast.TokenGroupStart {
			break
		}

		endEdgeID = edgeID

		if edge.Next == 0 {
			break
		}

		edgeID = edge.Next
	}

	edge := p.Ast.Edges.GetOne(edgeID)

	if prevEdgeID == 0 {
		node := p.Ast.Nodes.Get(src)
		node.Edge = edge.Next
		p.Ast.Nodes.Save(src, node)
	} else {
		prevEdge := p.Ast.Edges.GetOne(prevEdgeID)
		prevEdge.Next = edge.Next
		p.Ast.Edges.Set(prevEdgeID, prevEdge)
	}

	dstNode := p.Ast.Nodes.Get(dst)
	endEdge := p.Ast.Edges.GetOne(endEdgeID)

	endEdge.Next = dstNode.Edge
	dstNode.Edge = startEdgeID

	p.Ast.Nodes.Save(dst, dstNode)
	p.Ast.Edges.Set(endEdgeID, endEdge)
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
