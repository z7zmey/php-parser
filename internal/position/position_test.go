package position_test

import (
	"gotest.tools/assert"
	"testing"

	builder "github.com/z7zmey/php-parser/internal/position"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/token"
)

func TestNewTokenPosition(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewTokenPosition(tkn)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 1, EndPos: 3}, pos)
}

func TestNewTokensPosition(t *testing.T) {
	token1 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    6,
		},
	}

	pos := builder.NewTokensPosition(token1, token2)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 2, EndPos: 6}, pos)
}

func TestNewNodePosition(t *testing.T) {
	n := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewNodePosition(n)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 1, EndPos: 3}, pos)
}

func TestNewTokenNodePosition(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	n := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    12,
		},
	}

	pos := builder.NewTokenNodePosition(tkn, n)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 2, EndPos: 12}, pos)
}

func TestNewNodeTokenPosition(t *testing.T) {
	n := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    12,
		},
	}

	pos := builder.NewNodeTokenPosition(n, tkn)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 2, EndPos: 12}, pos)
}

func TestNewNodeListPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	pos := builder.NewNodeListPosition([]ast.Vertex{n1, n2})

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 2, EndPos: 19}, pos)
}

func TestNewNodesPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	pos := builder.NewNodesPosition(n1, n2)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 2, EndPos: 19}, pos)
}

func TestNewNodeListTokenPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    9,
		},
	}

	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  10,
			EndPos:    19,
		},
	}

	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  20,
			EndPos:    22,
		},
	}

	pos := builder.NewNodeListTokenPosition([]ast.Vertex{n1, n2}, tkn)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 3, EndPos: 22}, pos)
}

func TestNewTokenNodeListPosition(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    2,
		},
	}

	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  3,
			EndPos:    10,
		},
	}

	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  11,
			EndPos:    20,
		},
	}

	pos := builder.NewTokenNodeListPosition(tkn, []ast.Vertex{n1, n2})

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 3, EndPos: 20}, pos)
}

func TestNewNodeNodeListPosition(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}

	n3 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	pos := builder.NewNodeNodeListPosition(n1, []ast.Vertex{n2, n3})

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 3, EndPos: 26}, pos)
}

func TestNewNodeListNodePosition(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}
	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}
	n3 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	pos := builder.NewNodeListNodePosition([]ast.Vertex{n1, n2}, n3)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 3, EndPos: 26}, pos)
}

func TestNewOptionalListTokensPosition(t *testing.T) {
	token1 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  4,
			EndPos:    6,
		},
	}

	pos := builder.NewOptionalListTokensPosition(nil, token1, token2)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: 2, EndPos: 6}, pos)
}

func TestNewOptionalListTokensPosition2(t *testing.T) {
	n2 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 2,
			EndLine:   2,
			StartPos:  9,
			EndPos:    17,
		},
	}
	n3 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 3,
			EndLine:   3,
			StartPos:  18,
			EndPos:    26,
		},
	}

	token1 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 4,
			EndLine:   4,
			StartPos:  27,
			EndPos:    29,
		},
	}
	token2 := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 5,
			EndLine:   5,
			StartPos:  30,
			EndPos:    32,
		},
	}

	pos := builder.NewOptionalListTokensPosition([]ast.Vertex{n2, n3}, token1, token2)

	assert.DeepEqual(t, &position.Position{StartLine: 2, EndLine: 5, StartPos: 9, EndPos: 32}, pos)
}

func TestNilNodePos(t *testing.T) {
	pos := builder.NewNodesPosition(nil, nil)

	assert.DeepEqual(t, &position.Position{StartLine: -1, EndLine: -1, StartPos: -1, EndPos: -1}, pos)
}

func TestNilNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	pos := builder.NewNodeNodeListPosition(n1, nil)

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: -1, EndPos: -1}, pos)
}

func TestNilNodeListTokenPos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewNodeListTokenPosition(nil, tkn)

	assert.DeepEqual(t, &position.Position{StartLine: -1, EndLine: 1, StartPos: -1, EndPos: 3}, pos)
}

func TestEmptyNodeListPos(t *testing.T) {
	n1 := &ast.Identifier{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    8,
		},
	}

	pos := builder.NewNodeNodeListPosition(n1, []ast.Vertex{})

	assert.DeepEqual(t, &position.Position{StartLine: 1, EndLine: -1, EndPos: -1}, pos)
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	tkn := &token.Token{
		Value: []byte(`foo`),
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  0,
			EndPos:    3,
		},
	}

	pos := builder.NewNodeListTokenPosition([]ast.Vertex{}, tkn)

	assert.DeepEqual(t, &position.Position{StartLine: -1, EndLine: 1, StartPos: -1, EndPos: 3}, pos)
}
