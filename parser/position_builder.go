package parser

import (
	"sync"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/scanner"
)

// PositionBuilder provide functions to constuct positions
type PositionBuilder struct {
	Positions    *Positions
	PositionPool *sync.Pool
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
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getListStartPos(list).startLine
	pos.EndLine = b.getListEndPos(list).endLine
	pos.StartPos = b.getListStartPos(list).startPos
	pos.EndPos = b.getListEndPos(list).endPos

	return pos
}

// NewNodePosition returns new Position
func (b *PositionBuilder) NewNodePosition(n node.Node) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getNodeStartPos(n).startLine
	pos.EndLine = b.getNodeEndPos(n).endLine
	pos.StartPos = b.getNodeStartPos(n).startPos
	pos.EndPos = b.getNodeEndPos(n).endPos

	return pos
}

// NewTokenPosition returns new Position
func (b *PositionBuilder) NewTokenPosition(t *scanner.Token) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = t.Position.StartLine
	pos.EndLine = t.Position.EndLine
	pos.StartPos = t.Position.StartPos
	pos.EndPos = t.Position.EndPos

	return pos
}

// NewTokensPosition returns new Position
func (b *PositionBuilder) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = startToken.Position.StartLine
	pos.EndLine = endToken.Position.EndLine
	pos.StartPos = startToken.Position.StartPos
	pos.EndPos = endToken.Position.EndPos

	return pos
}

// NewTokenNodePosition returns new Position
func (b *PositionBuilder) NewTokenNodePosition(t *scanner.Token, n node.Node) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = t.Position.StartLine
	pos.EndLine = b.getNodeEndPos(n).endLine
	pos.StartPos = t.Position.StartPos
	pos.EndPos = b.getNodeEndPos(n).endPos

	return pos
}

// NewNodeTokenPosition returns new Position
func (b *PositionBuilder) NewNodeTokenPosition(n node.Node, t *scanner.Token) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getNodeStartPos(n).startLine
	pos.EndLine = t.Position.EndLine
	pos.StartPos = b.getNodeStartPos(n).startPos
	pos.EndPos = t.Position.EndPos

	return pos
}

// NewNodesPosition returns new Position
func (b *PositionBuilder) NewNodesPosition(startNode node.Node, endNode node.Node) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getNodeStartPos(startNode).startLine
	pos.EndLine = b.getNodeEndPos(endNode).endLine
	pos.StartPos = b.getNodeStartPos(startNode).startPos
	pos.EndPos = b.getNodeEndPos(endNode).endPos

	return pos
}

// NewNodeListTokenPosition returns new Position
func (b *PositionBuilder) NewNodeListTokenPosition(list []node.Node, t *scanner.Token) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getListStartPos(list).startLine
	pos.EndLine = t.Position.EndLine
	pos.StartPos = b.getListStartPos(list).startPos
	pos.EndPos = t.Position.EndPos

	return pos
}

// NewTokenNodeListPosition returns new Position
func (b *PositionBuilder) NewTokenNodeListPosition(t *scanner.Token, list []node.Node) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = t.Position.StartLine
	pos.EndLine = b.getListEndPos(list).endLine
	pos.StartPos = t.Position.StartPos
	pos.EndPos = b.getListEndPos(list).endPos

	return pos
}

// NewNodeNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeNodeListPosition(n node.Node, list []node.Node) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getNodeStartPos(n).startLine
	pos.EndLine = b.getListEndPos(list).endLine
	pos.StartPos = b.getNodeStartPos(n).startPos
	pos.EndPos = b.getListEndPos(list).endPos

	return pos
}

// NewNodeListNodePosition returns new Position
func (b *PositionBuilder) NewNodeListNodePosition(list []node.Node, n node.Node) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)
	pos.StartLine = b.getListStartPos(list).startLine
	pos.EndLine = b.getNodeEndPos(n).endLine
	pos.StartPos = b.getListStartPos(list).startPos
	pos.EndPos = b.getNodeEndPos(n).endPos

	return pos
}

// NewOptionalListTokensPosition returns new Position
func (b *PositionBuilder) NewOptionalListTokensPosition(list []node.Node, t *scanner.Token, endToken *scanner.Token) *position.Position {
	pos := b.PositionPool.Get().(*position.Position)

	if list == nil {
		pos.StartLine = t.Position.StartLine
		pos.EndLine = endToken.Position.EndLine
		pos.StartPos = t.Position.StartPos
		pos.EndPos = endToken.Position.EndPos
		return pos
	}

	pos.StartLine = b.getListStartPos(list).startLine
	pos.EndLine = endToken.Position.EndLine
	pos.StartPos = b.getListStartPos(list).startPos
	pos.EndPos = endToken.Position.EndPos
	return pos
}
