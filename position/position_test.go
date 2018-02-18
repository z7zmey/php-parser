package position_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/token"
)

func TestNewTokenPosition(t *testing.T) {
	builder := position.Builder{}

	tkn := token.NewToken([]byte(`foo`), 1, 1, 0, 3)

	pos := builder.NewTokenPosition(tkn)

	if pos.String() != `Pos{Line: 1-1 Pos: 0-3}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewTokensPosition(t *testing.T) {
	builder := position.Builder{}

	token1 := token.NewToken([]byte(`foo`), 1, 1, 0, 3)
	token2 := token.NewToken([]byte(`foo`), 2, 2, 4, 6)

	pos := builder.NewTokensPosition(token1, token2)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-6}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodePosition(t *testing.T) {
	n := node.NewIdentifier("test node")

	p := &position.Positions{}
	p.AddPosition(n, &position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	})

	builder := position.Builder{
		Positions: p,
	}

	pos := builder.NewNodePosition(n)

	if pos.String() != `Pos{Line: 1-1 Pos: 0-3}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewTokenNodePosition(t *testing.T) {
	tkn := token.NewToken([]byte(`foo`), 1, 1, 0, 3)
	n := node.NewIdentifier("test node")

	p := &position.Positions{}
	p.AddPosition(n, &position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    12,
	})

	builder := position.Builder{
		Positions: p,
	}

	pos := builder.NewTokenNodePosition(tkn, n)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-12}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeTokenPosition(t *testing.T) {
	n := node.NewIdentifier("test node")
	tkn := token.NewToken([]byte(`foo`), 2, 2, 10, 12)

	p := &position.Positions{}
	p.AddPosition(n, &position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    9,
	})

	builder := position.Builder{
		Positions: p,
	}

	pos := builder.NewNodeTokenPosition(n, tkn)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-12}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeListPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n2 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
			n2: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  10,
				EndPos:    19,
			},
		},
	}

	pos := builder.NewNodeListPosition([]node.Node{n1, n2})

	if pos.String() != `Pos{Line: 1-2 Pos: 0-19}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodesPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n2 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
			n2: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  10,
				EndPos:    19,
			},
		},
	}

	pos := builder.NewNodesPosition(n1, n2)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-19}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeListTokenPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n2 := node.NewIdentifier("test node")
	tkn := token.NewToken([]byte(`foo`), 3, 3, 20, 22)

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    9,
			},
			n2: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  10,
				EndPos:    19,
			},
		},
	}

	pos := builder.NewNodeListTokenPosition([]node.Node{n1, n2}, tkn)

	if pos.String() != `Pos{Line: 1-3 Pos: 0-22}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewTokenNodeListPosition(t *testing.T) {
	tkn := token.NewToken([]byte(`foo`), 1, 1, 0, 2)
	n1 := node.NewIdentifier("test node")
	n2 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  3,
				EndPos:    10,
			},
			n2: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  11,
				EndPos:    20,
			},
		},
	}

	pos := builder.NewTokenNodeListPosition(tkn, []node.Node{n1, n2})

	if pos.String() != `Pos{Line: 1-3 Pos: 0-20}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeNodeListPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n2 := node.NewIdentifier("test node")
	n3 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
			n2: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  9,
				EndPos:    17,
			},
			n3: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  18,
				EndPos:    26,
			},
		},
	}

	pos := builder.NewNodeNodeListPosition(n1, []node.Node{n2, n3})

	if pos.String() != `Pos{Line: 1-3 Pos: 0-26}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewOptionalListTokensPosition(t *testing.T) {
	builder := position.Builder{}

	token1 := token.NewToken([]byte(`foo`), 1, 1, 0, 3)
	token2 := token.NewToken([]byte(`foo`), 2, 2, 4, 6)

	pos := builder.NewOptionalListTokensPosition(nil, token1, token2)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-6}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewOptionalListTokensPosition2(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n2 := node.NewIdentifier("test node")
	n3 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
			n2: &position.Position{
				StartLine: 2,
				EndLine:   2,
				StartPos:  9,
				EndPos:    17,
			},
			n3: &position.Position{
				StartLine: 3,
				EndLine:   3,
				StartPos:  18,
				EndPos:    26,
			},
		},
	}

	token1 := token.NewToken([]byte(`foo`), 4, 4, 27, 29)
	token2 := token.NewToken([]byte(`foo`), 5, 5, 30, 32)

	pos := builder.NewOptionalListTokensPosition([]node.Node{n2, n3}, token1, token2)

	if pos.String() != `Pos{Line: 2-5 Pos: 9-32}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNilNodePos(t *testing.T) {
	builder := position.Builder{}

	pos := builder.NewNodesPosition(nil, nil)

	if pos.String() != `Pos{Line: -1--1 Pos: -1--1}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNilNodeListPos(t *testing.T) {
	n1 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
		},
	}

	pos := builder.NewNodeNodeListPosition(n1, nil)

	if pos.String() != `Pos{Line: 1--1 Pos: 0--1}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNilNodeListTokenPos(t *testing.T) {
	token1 := token.NewToken([]byte(`foo`), 1, 1, 0, 3)

	builder := position.Builder{}

	pos := builder.NewNodeListTokenPosition(nil, token1)

	if pos.String() != `Pos{Line: -1-1 Pos: -1-3}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestEmptyNodeListPos(t *testing.T) {
	n1 := node.NewIdentifier("test node")

	builder := position.Builder{
		Positions: &position.Positions{
			n1: &position.Position{
				StartLine: 1,
				EndLine:   1,
				StartPos:  0,
				EndPos:    8,
			},
		},
	}

	pos := builder.NewNodeNodeListPosition(n1, []node.Node{})

	if pos.String() != `Pos{Line: 1--1 Pos: 0--1}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	token1 := token.NewToken([]byte(`foo`), 1, 1, 0, 3)

	builder := position.Builder{}

	pos := builder.NewNodeListTokenPosition([]node.Node{}, token1)

	if pos.String() != `Pos{Line: -1-1 Pos: -1-3}` {
		t.Errorf("token value is not equal\n")
	}
}
