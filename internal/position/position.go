package position

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

type startPos struct {
	startLine int
	startPos  int
}

type endPos struct {
	endLine int
	endPos  int
}

func getListStartPos(l []ast.Vertex) startPos {
	if l == nil {
		return startPos{-1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1}
	}

	return getNodeStartPos(l[0])
}

func getNodeStartPos(n ast.Vertex) startPos {
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

func getListEndPos(l []ast.Vertex) endPos {
	if l == nil {
		return endPos{-1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1}
	}

	return getNodeEndPos(l[len(l)-1])
}

func getNodeEndPos(n ast.Vertex) endPos {
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
func NewNodeListPosition(list []ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: getListStartPos(list).startLine,
		EndLine:   getListEndPos(list).endLine,
		StartPos:  getListStartPos(list).startPos,
		EndPos:    getListEndPos(list).endPos,
	}
}

// NewNodePosition returns new Position
func NewNodePosition(n ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: getNodeStartPos(n).startLine,
		EndLine:   getNodeEndPos(n).endLine,
		StartPos:  getNodeStartPos(n).startPos,
		EndPos:    getNodeEndPos(n).endPos,
	}
}

// NewTokenPosition returns new Position
func NewTokenPosition(t *token.Token) *position.Position {
	return &position.Position{
		StartLine: t.Position.StartLine,
		EndLine:   t.Position.EndLine,
		StartPos:  t.Position.StartPos,
		EndPos:    t.Position.EndPos,
	}
}

// NewTokensPosition returns new Position
func NewTokensPosition(startToken *token.Token, endToken *token.Token) *position.Position {
	return &position.Position{
		StartLine: startToken.Position.StartLine,
		EndLine:   endToken.Position.EndLine,
		StartPos:  startToken.Position.StartPos,
		EndPos:    endToken.Position.EndPos,
	}
}

// NewTokenNodePosition returns new Position
func NewTokenNodePosition(t *token.Token, n ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: t.Position.StartLine,
		EndLine:   getNodeEndPos(n).endLine,
		StartPos:  t.Position.StartPos,
		EndPos:    getNodeEndPos(n).endPos,
	}
}

// NewNodeTokenPosition returns new Position
func NewNodeTokenPosition(n ast.Vertex, t *token.Token) *position.Position {
	return &position.Position{
		StartLine: getNodeStartPos(n).startLine,
		EndLine:   t.Position.EndLine,
		StartPos:  getNodeStartPos(n).startPos,
		EndPos:    t.Position.EndPos,
	}
}

// NewNodesPosition returns new Position
func NewNodesPosition(startNode ast.Vertex, endNode ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: getNodeStartPos(startNode).startLine,
		EndLine:   getNodeEndPos(endNode).endLine,
		StartPos:  getNodeStartPos(startNode).startPos,
		EndPos:    getNodeEndPos(endNode).endPos,
	}
}

// NewNodeListTokenPosition returns new Position
func NewNodeListTokenPosition(list []ast.Vertex, t *token.Token) *position.Position {
	return &position.Position{
		StartLine: getListStartPos(list).startLine,
		EndLine:   t.Position.EndLine,
		StartPos:  getListStartPos(list).startPos,
		EndPos:    t.Position.EndPos,
	}
}

// NewTokenNodeListPosition returns new Position
func NewTokenNodeListPosition(t *token.Token, list []ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: t.Position.StartLine,
		EndLine:   getListEndPos(list).endLine,
		StartPos:  t.Position.StartPos,
		EndPos:    getListEndPos(list).endPos,
	}
}

// NewNodeNodeListPosition returns new Position
func NewNodeNodeListPosition(n ast.Vertex, list []ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: getNodeStartPos(n).startLine,
		EndLine:   getListEndPos(list).endLine,
		StartPos:  getNodeStartPos(n).startPos,
		EndPos:    getListEndPos(list).endPos,
	}
}

// NewNodeListNodePosition returns new Position
func NewNodeListNodePosition(list []ast.Vertex, n ast.Vertex) *position.Position {
	return &position.Position{
		StartLine: getListStartPos(list).startLine,
		EndLine:   getNodeEndPos(n).endLine,
		StartPos:  getListStartPos(list).startPos,
		EndPos:    getNodeEndPos(n).endPos,
	}
}

// NewOptionalListTokensPosition returns new Position
func NewOptionalListTokensPosition(list []ast.Vertex, t *token.Token, endToken *token.Token) *position.Position {
	if list == nil {
		return &position.Position{
			StartLine: t.Position.StartLine,
			EndLine:   endToken.Position.EndLine,
			StartPos:  t.Position.StartPos,
			EndPos:    endToken.Position.EndPos,
		}
	}

	return &position.Position{
		StartLine: getListStartPos(list).startLine,
		EndLine:   endToken.Position.EndLine,
		StartPos:  getListStartPos(list).startPos,
		EndPos:    endToken.Position.EndPos,
	}
}
