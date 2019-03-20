package ast

import (
	"github.com/z7zmey/php-parser/scanner"
)

// PositionBuilder provide functions to constuct positions
type PositionBuilder struct {
	ast *AST
}

func NewPositionBuilder(a *AST) *PositionBuilder {
	return &PositionBuilder{
		ast: a,
	}
}

func (b *PositionBuilder) SetAST(a *AST) {
	b.ast = a
}

func (b *PositionBuilder) getListStartPosID(l []NodeID) PositionID {
	if len(l) > 0 {
		if l[0] == 0 {
			return 0
		}
		return b.ast.Nodes.Get(l[0]).Pos
	}

	return 0
}

func (b *PositionBuilder) getListEndPosID(l []NodeID) PositionID {
	if len(l) > 0 {
		if l[len(l)-1] == 0 {
			return 0
		}
		return b.ast.Nodes.Get(l[len(l)-1]).Pos
	}

	return 0
}

func (b *PositionBuilder) NewNodeListPosition(list []NodeID) PositionID {
	sPosID := b.getListStartPosID(list)
	ePosID := b.getListEndPosID(list)

	if sPosID == 0 || ePosID == 0 {
		return 0
	}

	s := b.ast.Positions.Get(sPosID)
	e := b.ast.Positions.Get(ePosID)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (b *PositionBuilder) NewTokenPosition(t *scanner.Token) PositionID {
	if t == nil {
		return PositionID(0)
	}

	return b.ast.Positions.Create(Position{
		PS: t.StartPos,
		PE: t.EndPos,
		LS: t.StartLine,
		LE: t.EndLine,
	})
}

func (b *PositionBuilder) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) PositionID {
	if startToken == nil || endToken == nil {
		return PositionID(0)
	}

	return b.ast.Positions.Create(Position{
		PS: startToken.StartPos,
		PE: endToken.EndPos,
		LS: startToken.StartLine,
		LE: endToken.EndLine,
	})
}

func (b *PositionBuilder) NewTokenNodePosition(t *scanner.Token, n NodeID) PositionID {
	if t == nil || n == 0 {
		return PositionID(0)
	}

	nPos := b.ast.Nodes.Get(n).Pos
	if nPos == 0 {
		return 0
	}
	e := b.ast.Positions.Get(nPos)

	return b.ast.Positions.Create(Position{
		PS: t.StartPos,
		PE: e.PE,
		LS: t.StartLine,
		LE: e.LE,
	})
}

func (b *PositionBuilder) NewNodeTokenPosition(n NodeID, t *scanner.Token) PositionID {
	if n == 0 || t == nil {
		return PositionID(0)
	}

	nPos := b.ast.Nodes.Get(n).Pos
	if nPos == 0 {
		return 0
	}
	s := b.ast.Positions.Get(nPos)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: t.EndPos,
		LS: s.LS,
		LE: t.EndLine,
	})
}

func (b *PositionBuilder) NewNodesPosition(startNodeID NodeID, endNodeID NodeID) PositionID {
	if startNodeID == 0 || endNodeID == 0 {
		return PositionID(0)
	}

	sPos := b.ast.Nodes.Get(startNodeID).Pos
	ePos := b.ast.Nodes.Get(endNodeID).Pos

	if sPos == 0 || ePos == 0 {
		return 0
	}
	s := b.ast.Positions.Get(sPos)
	e := b.ast.Positions.Get(ePos)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

// NewNodePosition returns new Position
func (b *PositionBuilder) NewNodePosition(nodeID NodeID) PositionID {
	if nodeID == 0 {
		return PositionID(0)
	}

	posID := b.ast.Nodes.Get(nodeID).Pos
	pos := b.ast.Positions.Get(posID)

	return b.ast.Positions.Create(Position{
		PS: pos.PS,
		PE: pos.PE,
		LS: pos.LS,
		LE: pos.LE,
	})
}

func (b *PositionBuilder) NewNodeListTokenPosition(list []NodeID, t *scanner.Token) PositionID {
	if list == nil || t == nil {
		return PositionID(0)
	}

	sPosID := b.getListStartPosID(list)
	if sPosID == 0 {
		return 0
	}
	s := b.ast.Positions.Get(sPosID)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: t.EndPos,
		LS: s.LS,
		LE: t.EndLine,
	})
}

func (b *PositionBuilder) NewTokenNodeListPosition(t *scanner.Token, list []NodeID) PositionID {
	if t == nil || list == nil {
		return PositionID(0)
	}

	ePosID := b.getListEndPosID(list)
	if ePosID == 0 {
		return 0
	}
	e := b.ast.Positions.Get(ePosID)

	return b.ast.Positions.Create(Position{
		PS: t.StartPos,
		PE: e.PE,
		LS: t.StartLine,
		LE: e.LE,
	})
}

func (b *PositionBuilder) NewNodeNodeListPosition(n NodeID, list []NodeID) PositionID {
	if n == 0 || list == nil {
		return PositionID(0)
	}

	nPos := b.ast.Nodes.Get(n).Pos
	ePosID := b.getListEndPosID(list)
	if nPos == 0 || ePosID == 0 {
		return 0
	}
	s := b.ast.Positions.Get(nPos)
	e := b.ast.Positions.Get(ePosID)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (b *PositionBuilder) NewNodeListNodePosition(list []NodeID, n NodeID) PositionID {
	if list == nil || n == 0 {
		return PositionID(0)
	}

	sPosID := b.getListStartPosID(list)
	nPos := b.ast.Nodes.Get(n).Pos
	if sPosID == 0 || nPos == 0 {
		return 0
	}
	s := b.ast.Positions.Get(sPosID)
	e := b.ast.Positions.Get(nPos)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (b *PositionBuilder) NewOptionalListTokensPosition(list []NodeID, startToken *scanner.Token, endToken *scanner.Token) PositionID {
	if list == nil {
		if startToken == nil || endToken == nil {
			return PositionID(0)
		}

		return b.ast.Positions.Create(Position{
			PS: startToken.StartPos,
			PE: endToken.EndPos,
			LS: startToken.StartLine,
			LE: endToken.EndLine,
		})
	}

	if list == nil || endToken == nil {
		return PositionID(0)
	}

	sPosID := b.getListStartPosID(list)
	if sPosID == 0 {
		return 0
	}
	s := b.ast.Positions.Get(sPosID)

	return b.ast.Positions.Create(Position{
		PS: s.PS,
		PE: endToken.EndPos,
		LS: s.LS,
		LE: endToken.EndLine,
	})
}
