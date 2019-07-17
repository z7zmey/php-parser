package graph

import "github.com/z7zmey/php-parser/pkg/ast"

var EdgeTypePosition = NewEdgeType("position")

type PositionID uint32

// PositionStorage stores positions
type PositionStorage []ast.Position

// Create saves new Position in store
func (b *PositionStorage) Create(n ast.Position) PositionID {
	*b = append(*b, n)
	return PositionID(len(*b))
}

// Save modified Position
func (b PositionStorage) Save(id PositionID, n ast.Position) {
	b[id-1] = n
}

// Get returns position by PositionID
func (b PositionStorage) Get(id PositionID) ast.Position {
	return b[id-1]
}
