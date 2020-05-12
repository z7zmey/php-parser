package positionbuilder_test

import (
	"gotest.tools/assert"
	"testing"

	"github.com/z7zmey/php-parser/internal/positionbuilder"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/position"
)

func TestNewTokenPosition(t *testing.T) {
	builder := positionbuilder.PositionBuilder{}

	tkn := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	pos := builder.NewTokenPosition(tkn)

	assert.DeepEqual(t, &position.Position{1, 1, 0, 3}, pos)

	assert.DeepEqual(t, &position.Position{1, 1, 0, 3}, pos)
}

func TestNewTokensPosition(t *testing.T) {
	builder := positionbuilder.PositionBuilder{}

	token1 := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	token2 := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    6,
	}

	pos := builder.NewTokensPosition(token1, token2)

	assert.DeepEqual(t, &position.Position{1, 2, 0, 6}, pos)
}

func TestNewNodePosition(t *testing.T) {
	n := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    3,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodePosition(n)

	assert.DeepEqual(t, &position.Position{1, 1, 0, 3}, pos)
}

func TestNewTokenNodePosition(t *testing.T) {
	tkn := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	n := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  4,
				EndPos:    12,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewTokenNodePosition(tkn, n)

	assert.DeepEqual(t, &position.Position{1, 2, 0, 12}, pos)
}

func TestNewNodeTokenPosition(t *testing.T) {
	n := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
		},
	}

	tkn := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    12,
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeTokenPosition(n, tkn)

	assert.DeepEqual(t, &position.Position{1, 2, 0, 12}, pos)
}

func TestNewNodeListPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
		},
	}

	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  10,
				EndPos:    19,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeListPosition([]ast.Vertex{n1, n2})

	assert.DeepEqual(t, &position.Position{1, 2, 0, 19}, pos)
}

func TestNewNodesPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
		},
	}

	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  10,
				EndPos:    19,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodesPosition(n1, n2)

	assert.DeepEqual(t, &position.Position{1, 2, 0, 19}, pos)
}

func TestNewNodeListTokenPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
		},
	}

	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  10,
				EndPos:    19,
			},
		},
	}

	tkn := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 3,
		EndLine:   3,
		StartPos:  20,
		EndPos:    22,
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeListTokenPosition([]ast.Vertex{n1, n2}, tkn)

	assert.DeepEqual(t, &position.Position{1, 3, 0, 22}, pos)
}

func TestNewTokenNodeListPosition(t *testing.T) {
	tkn := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    2,
	}

	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  3,
				EndPos:    10,
			},
		},
	}

	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  11,
				EndPos:    20,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewTokenNodeListPosition(tkn, []ast.Vertex{n1, n2})

	assert.DeepEqual(t, &position.Position{1, 3, 0, 20}, pos)
}

func TestNewNodeNodeListPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
		},
	}

	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  9,
				EndPos:    17,
			},
		},
	}

	n3 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  18,
				EndPos:    26,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeNodeListPosition(n1, []ast.Vertex{n2, n3})

	assert.DeepEqual(t, &position.Position{1, 3, 0, 26}, pos)
}

func TestNewNodeListNodePosition(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
		},
	}
	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  9,
				EndPos:    17,
			},
		},
	}
	n3 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  18,
				EndPos:    26,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeListNodePosition([]ast.Vertex{n1, n2}, n3)

	assert.DeepEqual(t, &position.Position{1, 3, 0, 26}, pos)
}

func TestNewOptionalListTokensPosition(t *testing.T) {
	builder := positionbuilder.PositionBuilder{}

	token1 := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	token2 := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    6,
	}

	pos := builder.NewOptionalListTokensPosition(nil, token1, token2)

	assert.DeepEqual(t, &position.Position{1, 2, 0, 6}, pos)
}

func TestNewOptionalListTokensPosition2(t *testing.T) {
	n2 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  9,
				EndPos:    17,
			},
		},
	}
	n3 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  18,
				EndPos:    26,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	token1 := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 4,
		EndLine:   4,
		StartPos:  27,
		EndPos:    29,
	}
	token2 := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 5,
		EndLine:   5,
		StartPos:  30,
		EndPos:    32,
	}

	pos := builder.NewOptionalListTokensPosition([]ast.Vertex{n2, n3}, token1, token2)

	assert.DeepEqual(t, &position.Position{2, 5, 9, 32}, pos)
}

func TestNilNodePos(t *testing.T) {
	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodesPosition(nil, nil)

	assert.DeepEqual(t, &position.Position{-1, -1, -1, -1}, pos)
}

func TestNilNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeNodeListPosition(n1, nil)

	assert.DeepEqual(t, &position.Position{1, -1, 0, -1}, pos)
}

func TestNilNodeListTokenPos(t *testing.T) {
	token := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeListTokenPosition(nil, token)

	assert.DeepEqual(t, &position.Position{-1, 1, -1, 3}, pos)
}

func TestEmptyNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Node: ast.Node{
			Position: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
		},
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeNodeListPosition(n1, []ast.Vertex{})

	assert.DeepEqual(t, &position.Position{1, -1, 0, -1}, pos)
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	token := &scanner.Token{
		Value:     []byte(`foo`),
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	builder := positionbuilder.PositionBuilder{}

	pos := builder.NewNodeListTokenPosition([]ast.Vertex{}, token)

	assert.DeepEqual(t, &position.Position{-1, 1, -1, 3}, pos)
}
