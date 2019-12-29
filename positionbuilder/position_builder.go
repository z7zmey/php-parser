package positionbuilder

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

// PositionBuilder provide functions to constuct positions
type PositionBuilder struct{}

type startPos struct {
	startLine int
	startPos  int
}

type endPos struct {
	endLine int
	endPos  int
}

func (b *PositionBuilder) getListStartPos(l []node.Node) startPos {
	if l == nil {
		return startPos{-1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1}
	}

	return b.getNodeStartPos(l[0])
}

func (b *PositionBuilder) getNodeStartPos(n node.Node) startPos {
	sl := -1
	sp := -1

	if n == nil {
		return startPos{-1, -1}
	}

	p := n.GetPosition()
	if p != nil {
		sl = p.StartLine
		sp = p.StartPos
	}

	return startPos{sl, sp}
}

func (b *PositionBuilder) getListEndPos(l []node.Node) endPos {
	if l == nil {
		return endPos{-1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1}
	}

	return b.getNodeEndPos(l[len(l)-1])
}

func (b *PositionBuilder) getNodeEndPos(n node.Node) endPos {
	el := -1
	ep := -1

	if n == nil {
		return endPos{-1, -1}
	}

	p := n.GetPosition()
	if p != nil {
		el = p.EndLine
		ep = p.EndPos
	}

	return endPos{el, ep}
}

// NewNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeListPosition(list []node.Node) *position.Position {
	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   b.getListEndPos(list).endLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    b.getListEndPos(list).endPos,
	}
}

// NewNodePosition returns new Position
func (b *PositionBuilder) NewNodePosition(n node.Node) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(n).startLine,
		EndLine:   b.getNodeEndPos(n).endLine,
		StartPos:  b.getNodeStartPos(n).startPos,
		EndPos:    b.getNodeEndPos(n).endPos,
	}
}

// NewTokenPosition returns new Position
func (b *PositionBuilder) NewTokenPosition(t *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: t.StartLine,
		EndLine:   t.EndLine,
		StartPos:  t.StartPos,
		EndPos:    t.EndPos,
	}
}

// NewTokensPosition returns new Position
func (b *PositionBuilder) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: startToken.StartLine,
		EndLine:   endToken.EndLine,
		StartPos:  startToken.StartPos,
		EndPos:    endToken.EndPos,
	}
}

// NewTokenNodePosition returns new Position
func (b *PositionBuilder) NewTokenNodePosition(t *scanner.Token, n node.Node) *position.Position {
	return &position.Position{
		StartLine: t.StartLine,
		EndLine:   b.getNodeEndPos(n).endLine,
		StartPos:  t.StartPos,
		EndPos:    b.getNodeEndPos(n).endPos,
	}
}

// NewNodeTokenPosition returns new Position
func (b *PositionBuilder) NewNodeTokenPosition(n node.Node, t *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(n).startLine,
		EndLine:   t.EndLine,
		StartPos:  b.getNodeStartPos(n).startPos,
		EndPos:    t.EndPos,
	}
}

// NewNodesPosition returns new Position
func (b *PositionBuilder) NewNodesPosition(startNode node.Node, endNode node.Node) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(startNode).startLine,
		EndLine:   b.getNodeEndPos(endNode).endLine,
		StartPos:  b.getNodeStartPos(startNode).startPos,
		EndPos:    b.getNodeEndPos(endNode).endPos,
	}
}

// NewNodeListTokenPosition returns new Position
func (b *PositionBuilder) NewNodeListTokenPosition(list []node.Node, t *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   t.EndLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    t.EndPos,
	}
}

// NewTokenNodeListPosition returns new Position
func (b *PositionBuilder) NewTokenNodeListPosition(t *scanner.Token, list []node.Node) *position.Position {
	return &position.Position{
		StartLine: t.StartLine,
		EndLine:   b.getListEndPos(list).endLine,
		StartPos:  t.StartPos,
		EndPos:    b.getListEndPos(list).endPos,
	}
}

// NewNodeNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeNodeListPosition(n node.Node, list []node.Node) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(n).startLine,
		EndLine:   b.getListEndPos(list).endLine,
		StartPos:  b.getNodeStartPos(n).startPos,
		EndPos:    b.getListEndPos(list).endPos,
	}
}

// NewNodeListNodePosition returns new Position
func (b *PositionBuilder) NewNodeListNodePosition(list []node.Node, n node.Node) *position.Position {
	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   b.getNodeEndPos(n).endLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    b.getNodeEndPos(n).endPos,
	}
}

// NewOptionalListTokensPosition returns new Position
func (b *PositionBuilder) NewOptionalListTokensPosition(list []node.Node, t *scanner.Token, endToken *scanner.Token) *position.Position {
	if list == nil {
		return &position.Position{
			StartLine: t.StartLine,
			EndLine:   endToken.EndLine,
			StartPos:  t.StartPos,
			EndPos:    endToken.EndPos,
		}
	}

	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   endToken.EndLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    endToken.EndPos,
	}
}
