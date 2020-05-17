package positionbuilder

import (
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/position"
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

func (b *PositionBuilder) getListStartPos(l []ast.Vertex) startPos {
	if l == nil {
		return startPos{-1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1}
	}

	return b.getNodeStartPos(l[0])
}

func (b *PositionBuilder) getNodeStartPos(n ast.Vertex) startPos {
	sl := -1
	sp := -1

	if n == nil {
		return startPos{-1, -1}
	}

	p := n.GetNode().Position
	if p != nil {
		sl = p.StartLine
		sp = p.StartPos
	}

	return startPos{sl, sp}
}

func (b *PositionBuilder) getListEndPos(l []ast.Vertex) endPos {
	if l == nil {
		return endPos{-1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1}
	}

	return b.getNodeEndPos(l[len(l)-1])
}

func (b *PositionBuilder) getNodeEndPos(n ast.Vertex) endPos {
	el := -1
	ep := -1

	if n == nil {
		return endPos{-1, -1}
	}

	p := n.GetNode().Position
	if p != nil {
		el = p.EndLine
		ep = p.EndPos
	}

	return endPos{el, ep}
}

// NewNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeListPosition(list []ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   b.getListEndPos(list).endLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    b.getListEndPos(list).endPos,
	}
}

// NewNodePosition returns new Position
func (b *PositionBuilder) NewNodePosition(n ast.Vertex) *position.Position {
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
		StartLine: t.Position.StartLine,
		EndLine:   t.Position.EndLine,
		StartPos:  t.Position.StartPos,
		EndPos:    t.Position.EndPos,
	}
}

// NewTokensPosition returns new Position
func (b *PositionBuilder) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: startToken.Position.StartLine,
		EndLine:   endToken.Position.EndLine,
		StartPos:  startToken.Position.StartPos,
		EndPos:    endToken.Position.EndPos,
	}
}

// NewTokenNodePosition returns new Position
func (b *PositionBuilder) NewTokenNodePosition(t *scanner.Token, n ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: t.Position.StartLine,
		EndLine:   b.getNodeEndPos(n).endLine,
		StartPos:  t.Position.StartPos,
		EndPos:    b.getNodeEndPos(n).endPos,
	}
}

// NewNodeTokenPosition returns new Position
func (b *PositionBuilder) NewNodeTokenPosition(n ast.Vertex, t *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(n).startLine,
		EndLine:   t.Position.EndLine,
		StartPos:  b.getNodeStartPos(n).startPos,
		EndPos:    t.Position.EndPos,
	}
}

// NewNodesPosition returns new Position
func (b *PositionBuilder) NewNodesPosition(startNode ast.Vertex, endNode ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(startNode).startLine,
		EndLine:   b.getNodeEndPos(endNode).endLine,
		StartPos:  b.getNodeStartPos(startNode).startPos,
		EndPos:    b.getNodeEndPos(endNode).endPos,
	}
}

// NewNodeListTokenPosition returns new Position
func (b *PositionBuilder) NewNodeListTokenPosition(list []ast.Vertex, t *scanner.Token) *position.Position {
	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   t.Position.EndLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    t.Position.EndPos,
	}
}

// NewTokenNodeListPosition returns new Position
func (b *PositionBuilder) NewTokenNodeListPosition(t *scanner.Token, list []ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: t.Position.StartLine,
		EndLine:   b.getListEndPos(list).endLine,
		StartPos:  t.Position.StartPos,
		EndPos:    b.getListEndPos(list).endPos,
	}
}

// NewNodeNodeListPosition returns new Position
func (b *PositionBuilder) NewNodeNodeListPosition(n ast.Vertex, list []ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: b.getNodeStartPos(n).startLine,
		EndLine:   b.getListEndPos(list).endLine,
		StartPos:  b.getNodeStartPos(n).startPos,
		EndPos:    b.getListEndPos(list).endPos,
	}
}

// NewNodeListNodePosition returns new Position
func (b *PositionBuilder) NewNodeListNodePosition(list []ast.Vertex, n ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   b.getNodeEndPos(n).endLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    b.getNodeEndPos(n).endPos,
	}
}

// NewOptionalListTokensPosition returns new Position
func (b *PositionBuilder) NewOptionalListTokensPosition(list []ast.Vertex, t *scanner.Token, endToken *scanner.Token) *position.Position {
	if list == nil {
		return &position.Position{
			StartLine: t.Position.StartLine,
			EndLine:   endToken.Position.EndLine,
			StartPos:  t.Position.StartPos,
			EndPos:    endToken.Position.EndPos,
		}
	}

	return &position.Position{
		StartLine: b.getListStartPos(list).startLine,
		EndLine:   endToken.Position.EndLine,
		StartPos:  b.getListStartPos(list).startPos,
		EndPos:    endToken.Position.EndPos,
	}
}
