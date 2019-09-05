package ast

import "github.com/z7zmey/php-parser/pkg/ast"

var EdgeTypePosition = NewEdgeType("position")

type PositionID uint32

// PositionStorage stores positions
type PositionStorage []ast.Position

// Put saves new Position in store
func (b *PositionStorage) Create(n ast.Position) PositionID {
	*b = append(*b, n)
	return PositionID(len(*b))
}

// Get returns position by PositionID
func (b PositionStorage) Get(id PositionID) ast.Position {
	return b[id-1]
}
