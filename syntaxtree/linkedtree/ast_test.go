package linkedtree

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/scanner"
)

func TestNewTokenPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 3,
		LS: 1,
		LE: 1,
	}

	tkn := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	posID := ast.NewTokenPosition(tkn)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewTokensPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 6,
		LS: 1,
		LE: 2,
	}

	token1 := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	token2 := &scanner.Token{
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    6,
	}

	posID := ast.NewTokensPosition(token1, token2)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewTokenNodePosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 12,
		LS: 1,
		LE: 2,
	}

	tkn := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	nodePosID := ast.Positions.Create(Position{
		PS: 4,
		PE: 12,
		LS: 2,
		LE: 2,
	})
	nID := ast.Nodes.Create(Node{Pos: nodePosID})

	posID := ast.NewTokenNodePosition(tkn, nID)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeTokenPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 12,
		LS: 1,
		LE: 2,
	}

	nodePosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nID := ast.Nodes.Create(Node{Pos: nodePosID})

	tkn := &scanner.Token{
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    12,
	}

	posID := ast.NewNodeTokenPosition(nID, tkn)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeListPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 19,
		LS: 1,
		LE: 2,
	}

	nsPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	posID := ast.NewNodeListPosition([]NodeID{nsID, neID})
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodesPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 19,
		LS: 1,
		LE: 2,
	}

	nsPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	posID := ast.NewNodesPosition(nsID, neID)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeListTokenPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 22,
		LS: 1,
		LE: 3,
	}

	nsPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	tkn := &scanner.Token{
		StartLine: 3,
		EndLine:   3,
		StartPos:  20,
		EndPos:    22,
	}

	posID := ast.NewNodeListTokenPosition([]NodeID{nsID, neID}, tkn)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewTokenNodeListPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 20,
		LS: 1,
		LE: 3,
	}

	tkn := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    2,
	}

	nsPosID := ast.Positions.Create(Position{
		PS: 3,
		PE: 10,
		LS: 2,
		LE: 2,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 11,
		PE: 20,
		LS: 3,
		LE: 3,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	posID := ast.NewTokenNodeListPosition(tkn, []NodeID{nsID, neID})
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeNodeListPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 26,
		LS: 1,
		LE: 3,
	}

	nPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nID := ast.Nodes.Create(Node{Pos: nPosID})

	nsPosID := ast.Positions.Create(Position{
		PS: 9,
		PE: 17,
		LS: 2,
		LE: 2,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 18,
		PE: 26,
		LS: 3,
		LE: 3,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	posID := ast.NewNodeNodeListPosition(nID, []NodeID{nsID, neID})
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeListNodePosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 26,
		LS: 1,
		LE: 3,
	}

	nsPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 9,
		PE: 17,
		LS: 2,
		LE: 2,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	nPosID := ast.Positions.Create(Position{
		PS: 18,
		PE: 26,
		LS: 3,
		LE: 3,
	})
	nID := ast.Nodes.Create(Node{Pos: nPosID})

	posID := ast.NewNodeListNodePosition([]NodeID{nsID, neID}, nID)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewOptionalListTokensPosition(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 6,
		LS: 1,
		LE: 2,
	}

	token1 := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	token2 := &scanner.Token{
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    6,
	}

	posID := ast.NewOptionalListTokensPosition(nil, token1, token2)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewOptionalListTokensPosition2(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := Position{
		PS: 0,
		PE: 25,
		LS: 1,
		LE: 4,
	}

	nsPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := ast.Nodes.Create(Node{Pos: nsPosID})

	nePosID := ast.Positions.Create(Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := ast.Nodes.Create(Node{Pos: nePosID})

	tkn1 := &scanner.Token{
		StartLine: 3,
		EndLine:   3,
		StartPos:  20,
		EndPos:    22,
	}
	tkn2 := &scanner.Token{
		StartLine: 4,
		EndLine:   4,
		StartPos:  23,
		EndPos:    25,
	}

	posID := ast.NewOptionalListTokensPosition([]NodeID{nsID, neID}, tkn1, tkn2)
	actual := ast.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNilNodePos(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := PositionID(0)

	actual := ast.NewNodesPosition(NodeID(0), NodeID(0))
	assert.Equal(t, expected, actual)
}

func TestNilNodeListPos(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := PositionID(0)

	nPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nID := ast.Nodes.Create(Node{Pos: nPosID})

	actual := ast.NewNodeNodeListPosition(nID, nil)
	assert.Equal(t, expected, actual)
}

func TestNilNodeListTokenPos(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := PositionID(0)

	token := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	actual := ast.NewNodeListTokenPosition(nil, token)
	assert.Equal(t, expected, actual)
}

func TestEmptyNodeListPos(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := PositionID(0)

	nPosID := ast.Positions.Create(Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nID := ast.Nodes.Create(Node{Pos: nPosID})

	actual := ast.NewNodeNodeListPosition(nID, []NodeID{})
	assert.Equal(t, expected, actual)
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	ast := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
		Edges:     NewEdgeStorage(nil),
	}

	expected := PositionID(0)

	token := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	actual := ast.NewNodeListTokenPosition([]NodeID{}, token)
	assert.Equal(t, expected, actual)
}
