package linear

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/scanner"
)

func TestNewTokenPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
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

	posID := syntaxTree.NewTokenPosition(tkn)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewTokensPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
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

	posID := syntaxTree.NewTokensPosition(token1, token2)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewTokenNodePosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
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
	nodePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 4,
		PE: 12,
		LS: 2,
		LE: 2,
	})
	nID := syntaxTree.Nodes.Create(Node{Pos: nodePosID})

	posID := syntaxTree.NewTokenNodePosition(tkn, nID)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeTokenPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 12,
		LS: 1,
		LE: 2,
	}

	nodePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nID := syntaxTree.Nodes.Create(Node{Pos: nodePosID})

	tkn := &scanner.Token{
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    12,
	}

	posID := syntaxTree.NewNodeTokenPosition(nID, tkn)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeListPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 19,
		LS: 1,
		LE: 2,
	}

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

	posID := syntaxTree.NewNodeListPosition([]NodeID{nsID, neID})
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodesPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 19,
		LS: 1,
		LE: 2,
	}

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

	posID := syntaxTree.NewNodesPosition(nsID, neID)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeListTokenPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 22,
		LS: 1,
		LE: 3,
	}

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

	tkn := &scanner.Token{
		StartLine: 3,
		EndLine:   3,
		StartPos:  20,
		EndPos:    22,
	}

	posID := syntaxTree.NewNodeListTokenPosition([]NodeID{nsID, neID}, tkn)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewTokenNodeListPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
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

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 3,
		PE: 10,
		LS: 2,
		LE: 2,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 11,
		PE: 20,
		LS: 3,
		LE: 3,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

	posID := syntaxTree.NewTokenNodeListPosition(tkn, []NodeID{nsID, neID})
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeNodeListPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 26,
		LS: 1,
		LE: 3,
	}

	nPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nID := syntaxTree.Nodes.Create(Node{Pos: nPosID})

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 9,
		PE: 17,
		LS: 2,
		LE: 2,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 18,
		PE: 26,
		LS: 3,
		LE: 3,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

	posID := syntaxTree.NewNodeNodeListPosition(nID, []NodeID{nsID, neID})
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewNodeListNodePosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 26,
		LS: 1,
		LE: 3,
	}

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 9,
		PE: 17,
		LS: 2,
		LE: 2,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

	nPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 18,
		PE: 26,
		LS: 3,
		LE: 3,
	})
	nID := syntaxTree.Nodes.Create(Node{Pos: nPosID})

	posID := syntaxTree.NewNodeListNodePosition([]NodeID{nsID, neID}, nID)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewOptionalListTokensPosition(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
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

	posID := syntaxTree.NewOptionalListTokensPosition(nil, token1, token2)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNewOptionalListTokensPosition2(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := ast.Position{
		PS: 0,
		PE: 25,
		LS: 1,
		LE: 4,
	}

	nsPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 9,
		LS: 1,
		LE: 1,
	})
	nsID := syntaxTree.Nodes.Create(Node{Pos: nsPosID})

	nePosID := syntaxTree.Positions.Create(ast.Position{
		PS: 10,
		PE: 19,
		LS: 2,
		LE: 2,
	})
	neID := syntaxTree.Nodes.Create(Node{Pos: nePosID})

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

	posID := syntaxTree.NewOptionalListTokensPosition([]NodeID{nsID, neID}, tkn1, tkn2)
	actual := syntaxTree.Positions.Get(posID)

	assert.Equal(t, expected, actual)
}

func TestNilNodePos(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := PositionID(0)

	actual := syntaxTree.NewNodesPosition(NodeID(0), NodeID(0))
	assert.Equal(t, expected, actual)
}

func TestNilNodeListPos(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := PositionID(0)

	nPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nID := syntaxTree.Nodes.Create(Node{Pos: nPosID})

	actual := syntaxTree.NewNodeNodeListPosition(nID, nil)
	assert.Equal(t, expected, actual)
}

func TestNilNodeListTokenPos(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := PositionID(0)

	token := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	actual := syntaxTree.NewNodeListTokenPosition(nil, token)
	assert.Equal(t, expected, actual)
}

func TestEmptyNodeListPos(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := PositionID(0)

	nPosID := syntaxTree.Positions.Create(ast.Position{
		PS: 0,
		PE: 8,
		LS: 1,
		LE: 1,
	})
	nID := syntaxTree.Nodes.Create(Node{Pos: nPosID})

	actual := syntaxTree.NewNodeNodeListPosition(nID, []NodeID{})
	assert.Equal(t, expected, actual)
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	syntaxTree := AST{
		Positions: NewPositionStorage(nil),
		Nodes:     NewNodeStorage(nil),
	}

	expected := PositionID(0)

	token := &scanner.Token{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	actual := syntaxTree.NewNodeListTokenPosition([]NodeID{}, token)
	assert.Equal(t, expected, actual)
}
