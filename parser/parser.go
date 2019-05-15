package parser

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/ast/linear"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

type lastTknCache struct {
	nodeID  linear.NodeID
	tokenID linear.TokenID
}

type Parser interface {
	Parse([]byte, *linear.AST) int
	GetErrors() []*errors.Error
}

type AbstractParser struct {
	Lexer        scanner.Scanner
	CurrentToken *scanner.Token
	List         stackedNodeList
	Ast          *linear.AST
	lastTknCache lastTknCache
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

func (p *AbstractParser) Reset(src []byte, a *linear.AST) {
	p.Lexer.Reset(src)
	p.List.Reset()
	p.lastTknCache = lastTknCache{}

	a.FileData = src

	p.Ast = a
	p.Lexer.SetErrors(nil)
}

func (p *AbstractParser) getListStartPosID(l []linear.NodeID) linear.PositionID {
	if len(l) > 0 {
		if l[0] == 0 {
			return 0
		}
		return p.Ast.Nodes.Get(l[0]).Pos
	}

	return 0
}

func (p *AbstractParser) getListEndPosID(l []linear.NodeID) linear.PositionID {
	if len(l) > 0 {
		if l[len(l)-1] == 0 {
			return 0
		}
		return p.Ast.Nodes.Get(l[len(l)-1]).Pos
	}

	return 0
}

func (p *AbstractParser) NewNodeListPosition(list []linear.NodeID) linear.PositionID {
	sPosID := p.getListStartPosID(list)
	ePosID := p.getListEndPosID(list)

	if sPosID == 0 || ePosID == 0 {
		return 0
	}

	s := p.Ast.Positions.Get(sPosID)
	e := p.Ast.Positions.Get(ePosID)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (p *AbstractParser) NewTokenPosition(t *scanner.Token) linear.PositionID {
	if t == nil {
		return linear.PositionID(0)
	}

	return p.Ast.Positions.Create(ast.Position{
		PS: t.StartPos,
		PE: t.EndPos,
		LS: t.StartLine,
		LE: t.EndLine,
	})
}

func (p *AbstractParser) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) linear.PositionID {
	if startToken == nil || endToken == nil {
		return linear.PositionID(0)
	}

	return p.Ast.Positions.Create(ast.Position{
		PS: startToken.StartPos,
		PE: endToken.EndPos,
		LS: startToken.StartLine,
		LE: endToken.EndLine,
	})
}

func (p *AbstractParser) NewTokenNodePosition(t *scanner.Token, n linear.NodeID) linear.PositionID {
	if t == nil || n == 0 {
		return linear.PositionID(0)
	}

	nPos := p.Ast.Nodes.Get(n).Pos
	if nPos == 0 {
		return 0
	}
	e := p.Ast.Positions.Get(nPos)

	return p.Ast.Positions.Create(ast.Position{
		PS: t.StartPos,
		PE: e.PE,
		LS: t.StartLine,
		LE: e.LE,
	})
}

func (p *AbstractParser) NewNodeTokenPosition(n linear.NodeID, t *scanner.Token) linear.PositionID {
	if n == 0 || t == nil {
		return linear.PositionID(0)
	}

	nPos := p.Ast.Nodes.Get(n).Pos
	if nPos == 0 {
		return 0
	}
	s := p.Ast.Positions.Get(nPos)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: t.EndPos,
		LS: s.LS,
		LE: t.EndLine,
	})
}

func (p *AbstractParser) NewNodesPosition(startNodeID linear.NodeID, endNodeID linear.NodeID) linear.PositionID {
	if startNodeID == 0 || endNodeID == 0 {
		return linear.PositionID(0)
	}

	sPos := p.Ast.Nodes.Get(startNodeID).Pos
	ePos := p.Ast.Nodes.Get(endNodeID).Pos

	if sPos == 0 || ePos == 0 {
		return 0
	}
	s := p.Ast.Positions.Get(sPos)
	e := p.Ast.Positions.Get(ePos)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (p *AbstractParser) NewNodePosition(nodeID linear.NodeID) linear.PositionID {
	if nodeID == 0 {
		return linear.PositionID(0)
	}

	posID := p.Ast.Nodes.Get(nodeID).Pos
	pos := p.Ast.Positions.Get(posID)

	return p.Ast.Positions.Create(ast.Position{
		PS: pos.PS,
		PE: pos.PE,
		LS: pos.LS,
		LE: pos.LE,
	})
}

func (p *AbstractParser) NewNodeListTokenPosition(list []linear.NodeID, t *scanner.Token) linear.PositionID {
	if list == nil || t == nil {
		return linear.PositionID(0)
	}

	sPosID := p.getListStartPosID(list)
	if sPosID == 0 {
		return 0
	}
	s := p.Ast.Positions.Get(sPosID)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: t.EndPos,
		LS: s.LS,
		LE: t.EndLine,
	})
}

func (p *AbstractParser) NewTokenNodeListPosition(t *scanner.Token, list []linear.NodeID) linear.PositionID {
	if t == nil || list == nil {
		return linear.PositionID(0)
	}

	ePosID := p.getListEndPosID(list)
	if ePosID == 0 {
		return 0
	}
	e := p.Ast.Positions.Get(ePosID)

	return p.Ast.Positions.Create(ast.Position{
		PS: t.StartPos,
		PE: e.PE,
		LS: t.StartLine,
		LE: e.LE,
	})
}

func (p *AbstractParser) NewNodeNodeListPosition(n linear.NodeID, list []linear.NodeID) linear.PositionID {
	if n == 0 || list == nil {
		return linear.PositionID(0)
	}

	nPos := p.Ast.Nodes.Get(n).Pos
	ePosID := p.getListEndPosID(list)
	if nPos == 0 || ePosID == 0 {
		return 0
	}
	s := p.Ast.Positions.Get(nPos)
	e := p.Ast.Positions.Get(ePosID)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (p *AbstractParser) NewNodeListNodePosition(list []linear.NodeID, n linear.NodeID) linear.PositionID {
	if list == nil || n == 0 {
		return linear.PositionID(0)
	}

	sPosID := p.getListStartPosID(list)
	nPos := p.Ast.Nodes.Get(n).Pos
	if sPosID == 0 || nPos == 0 {
		return 0
	}
	s := p.Ast.Positions.Get(sPosID)
	e := p.Ast.Positions.Get(nPos)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (p *AbstractParser) NewOptionalListTokensPosition(list []linear.NodeID, startToken *scanner.Token, endToken *scanner.Token) linear.PositionID {
	if list == nil {
		if startToken == nil || endToken == nil {
			return linear.PositionID(0)
		}

		return p.Ast.Positions.Create(ast.Position{
			PS: startToken.StartPos,
			PE: endToken.EndPos,
			LS: startToken.StartLine,
			LE: endToken.EndLine,
		})
	}

	if list == nil || endToken == nil {
		return linear.PositionID(0)
	}

	sPosID := p.getListStartPosID(list)
	if sPosID == 0 {
		return 0
	}
	s := p.Ast.Positions.Get(sPosID)

	return p.Ast.Positions.Create(ast.Position{
		PS: s.PS,
		PE: endToken.EndPos,
		LS: s.LS,
		LE: endToken.EndLine,
	})
}

func (p *AbstractParser) Children(prevNodeID linear.NodeID, parentNodeID linear.NodeID, edgeType ast.EdgeType, children ...linear.NodeID) linear.NodeID {
	for _, childNodeID := range children {
		if childNodeID == 0 {
			continue
		}

		if prevNodeID == 0 {
			p.linkChild(parentNodeID, childNodeID)
		} else {
			p.linkNext(prevNodeID, childNodeID)
		}

		p.linkParent(childNodeID, parentNodeID, edgeType)

		prevNodeID = childNodeID
	}

	return prevNodeID
}

func (p *AbstractParser) linkParent(childNodeID, parentNodeID linear.NodeID, key ast.EdgeType) {
	childNode := p.Ast.Nodes.Get(childNodeID)
	childNode.Parent = parentNodeID
	childNode.Key = key
	p.Ast.Nodes.Save(childNodeID, childNode)
}

func (p *AbstractParser) linkChild(parentNodeID, childNodeID linear.NodeID) {
	parentNode := p.Ast.Nodes.Get(parentNodeID)
	parentNode.Child = childNodeID
	p.Ast.Nodes.Save(parentNodeID, parentNode)
}

func (p *AbstractParser) linkNext(prevNodeID, nextNodeID linear.NodeID) {
	prevNode := p.Ast.Nodes.Get(prevNodeID)
	prevNode.Next = nextNodeID
	p.Ast.Nodes.Save(prevNodeID, prevNode)
}

func (p *AbstractParser) lastToken(nodeID linear.NodeID) linear.TokenID {
	if p.lastTknCache.nodeID == nodeID {
		return p.lastTknCache.tokenID
	}

	node := p.Ast.Nodes.Get(nodeID)
	tknID := node.Tkn

	if tknID == 0 {
		return tknID
	}

	for {
		token := p.Ast.Tokens.Get(tknID)

		if token.Next == 0 {
			break
		}

		tknID = token.Next
	}

	p.lastTknCache = lastTknCache{
		nodeID:  nodeID,
		tokenID: tknID,
	}

	return tknID
}

func (p *AbstractParser) AppendTokens(nodeID linear.NodeID, group ast.TokenGroup, ffStrs []scanner.Token) {
	lastTokenID := p.lastToken(nodeID)

	for _, str := range ffStrs {
		tkn := p.convertToken(str)
		tkn.Group = group
		tokenID := p.Ast.Tokens.Create(tkn)

		if lastTokenID == 0 {
			node := p.Ast.Nodes.Get(nodeID)
			node.Tkn = tokenID
			p.Ast.Nodes.Save(nodeID, node)
		} else {
			prevString := p.Ast.Tokens.Get(lastTokenID)
			prevString.Next = tokenID
			p.Ast.Tokens.Save(lastTokenID, prevString)
		}

		lastTokenID = tokenID
	}

	p.lastTknCache = lastTknCache{
		nodeID:  nodeID,
		tokenID: lastTokenID,
	}
}

func (p *AbstractParser) PrependTokens(nodeID linear.NodeID, group ast.TokenGroup, ffStrs []scanner.Token) {
	node := p.Ast.Nodes.Get(nodeID)
	firstTokenID := node.Tkn

	var prevTokenID linear.TokenID
	for _, str := range ffStrs {
		tkn := p.convertToken(str)
		tkn.Group = group
		tkn.Next = firstTokenID
		tokenID := p.Ast.Tokens.Create(tkn)

		if prevTokenID == 0 {
			node := p.Ast.Nodes.Get(nodeID)
			node.Tkn = tokenID
			p.Ast.Nodes.Save(nodeID, node)
		} else {
			prevToken := p.Ast.Tokens.Get(prevTokenID)
			prevToken.Next = tokenID
			p.Ast.Tokens.Save(prevTokenID, prevToken)
		}

		prevTokenID = tokenID
	}
}

func (p *AbstractParser) MoveStartTokens(src linear.NodeID, dst linear.NodeID) {
	srcNode := p.Ast.Nodes.Get(src)

	if srcNode.Tkn == 0 {
		return
	}

	srcStartFirstTkn := p.Ast.Tokens.Get(srcNode.Tkn)
	if srcStartFirstTkn.Group != ast.TokenGroupStart {
		return
	}

	srcStartLastTknID := srcNode.Tkn
	srcStartLastTkn := srcStartFirstTkn

	for {
		if srcStartLastTkn.Next == 0 {
			break
		}

		tkn := p.Ast.Tokens.Get(srcStartLastTkn.Next)
		if tkn.Group == ast.TokenGroupStart {
			srcStartLastTknID = srcStartLastTkn.Next
			srcStartLastTkn = tkn
		} else {
			break
		}
	}

	dstNode := p.Ast.Nodes.Get(dst)

	// move

	dstNode.Tkn, srcNode.Tkn, srcStartLastTkn.Next = srcNode.Tkn, srcStartLastTkn.Next, dstNode.Tkn

	// save

	p.Ast.Nodes.Save(src, srcNode)
	p.Ast.Nodes.Save(dst, dstNode)
	p.Ast.Tokens.Save(srcStartLastTknID, srcStartLastTkn)
}

func (p *AbstractParser) convertToken(token scanner.Token) linear.Token {
	pos := ast.Position{
		PS: token.StartPos,
		PE: token.EndPos,
		LS: token.StartLine,
	}
	posID := p.Ast.Positions.Create(pos)

	return linear.Token{
		Type: token.Type,
		Pos:  posID,
	}
}
