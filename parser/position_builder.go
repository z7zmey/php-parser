package parser

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

// PositionBuilder provide functions to constuct positions
type PositionBuilder struct {
	Positions *Positions
}

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

	p := (*b.Positions)[n]
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

	p := (*b.Positions)[n]
	if p != nil {
		el = p.EndLine
		ep = p.EndPos
	}

	return endPos{el, ep}
}

// NewNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeListPosition(list []node.Node) *position.Position {
	return position.NewPosition(
		b.getListStartPos(list).startLine,
		b.getListEndPos(list).endLine,
		b.getListStartPos(list).startPos,
		b.getListEndPos(list).endPos,
	)
}

// NewNodePosition returns new Position
func (b *PositionBuilder) NewNodePosition(n node.Node) *position.Position {
	return position.NewPosition(
		b.getNodeStartPos(n).startLine,
		b.getNodeEndPos(n).endLine,
		b.getNodeStartPos(n).startPos,
		b.getNodeEndPos(n).endPos,
	)
}

// NewTokenPosition returns new Position
func (b *PositionBuilder) NewTokenPosition(t *scanner.Token) *position.Position {
	return t.Position
}

// NewTokensPosition returns new Position
func (b *PositionBuilder) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) *position.Position {
	return position.NewPosition(
		startToken.Position.StartLine,
		endToken.Position.EndLine,
		startToken.Position.StartPos,
		endToken.Position.EndPos,
	)
}

// NewTokenNodePosition returns new Position
func (b *PositionBuilder) NewTokenNodePosition(t *scanner.Token, n node.Node) *position.Position {
	return position.NewPosition(
		t.Position.StartLine,
		b.getNodeEndPos(n).endLine,
		t.Position.StartPos,
		b.getNodeEndPos(n).endPos,
	)
}

// NewNodeTokenPosition returns new Position
func (b *PositionBuilder) NewNodeTokenPosition(n node.Node, t *scanner.Token) *position.Position {
	return position.NewPosition(
		b.getNodeStartPos(n).startLine,
		t.Position.EndLine,
		b.getNodeStartPos(n).startPos,
		t.Position.EndPos,
	)
}

// NewNodesPosition returns new Position
func (b *PositionBuilder) NewNodesPosition(startNode node.Node, endNode node.Node) *position.Position {
	return position.NewPosition(
		b.getNodeStartPos(startNode).startLine,
		b.getNodeEndPos(endNode).endLine,
		b.getNodeStartPos(startNode).startPos,
		b.getNodeEndPos(endNode).endPos,
	)
}

// NewNodeListTokenPosition returns new Position
func (b *PositionBuilder) NewNodeListTokenPosition(list []node.Node, t *scanner.Token) *position.Position {
	return position.NewPosition(
		b.getListStartPos(list).startLine,
		t.Position.EndLine,
		b.getListStartPos(list).startPos,
		t.Position.EndPos,
	)
}

// NewTokenNodeListPosition returns new Position
func (b *PositionBuilder) NewTokenNodeListPosition(t *scanner.Token, list []node.Node) *position.Position {
	return position.NewPosition(
		t.Position.StartLine,
		b.getListEndPos(list).endLine,
		t.Position.StartPos,
		b.getListEndPos(list).endPos,
	)
}

// NewNodeNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeNodeListPosition(n node.Node, list []node.Node) *position.Position {
	return position.NewPosition(
		b.getNodeStartPos(n).startLine,
		b.getListEndPos(list).endLine,
		b.getNodeStartPos(n).startPos,
		b.getListEndPos(list).endPos,
	)
}

// NewNodeListNodePosition returns new Position
func (b *PositionBuilder) NewNodeListNodePosition(list []node.Node, n node.Node) *position.Position {
	return position.NewPosition(
		b.getListStartPos(list).startLine,
		b.getNodeEndPos(n).endLine,
		b.getListStartPos(list).startPos,
		b.getNodeEndPos(n).endPos,
	)
}

// NewOptionalListTokensPosition returns new Position
func (b *PositionBuilder) NewOptionalListTokensPosition(list []node.Node, t *scanner.Token, endToken *scanner.Token) *position.Position {
	if list == nil {
		return position.NewPosition(
			t.Position.StartLine,
			endToken.Position.EndLine,
			t.Position.StartPos,
			endToken.Position.EndPos,
		)
	}

	return position.NewPosition(
		b.getListStartPos(list).startLine,
		endToken.Position.EndLine,
		b.getListStartPos(list).startPos,
		endToken.Position.EndPos,
	)
}
