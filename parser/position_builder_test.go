package parser_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/scanner"
)

func TestNewTokenPosition(t *testing.T) {
	builder := parser.PositionBuilder{}

	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	pos := builder.NewTokenPosition(tkn)

	if pos.String() != `Pos{Line: 1-1 Pos: 0-3}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewTokensPosition(t *testing.T) {
	builder := parser.PositionBuilder{}

	token1 := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	token2 := &scanner.Token{
		Value:     `foo`,
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    6,
	}

	pos := builder.NewTokensPosition(token1, token2)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-6}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodePosition(t *testing.T) {
	n := node.NewIdentifier("test node")
	n.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodePosition(n)

	if pos.String() != `Pos{Line: 1-1 Pos: 0-3}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewTokenNodePosition(t *testing.T) {
	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	n := node.NewIdentifier("test node")
	n.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    12,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewTokenNodePosition(tkn, n)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-12}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeTokenPosition(t *testing.T) {
	n := node.NewIdentifier("test node")
	n.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    9,
	})

	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    12,
	}

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeTokenPosition(n, tkn)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-12}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeListPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    9,
	})

	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    19,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeListPosition([]node.Node{n1, n2})

	if pos.String() != `Pos{Line: 1-2 Pos: 0-19}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodesPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    9,
	})

	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    19,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodesPosition(n1, n2)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-19}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeListTokenPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    9,
	})

	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  10,
		EndPos:    19,
	})

	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 3,
		EndLine:   3,
		StartPos:  20,
		EndPos:    22,
	}

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeListTokenPosition([]node.Node{n1, n2}, tkn)

	if pos.String() != `Pos{Line: 1-3 Pos: 0-22}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewTokenNodeListPosition(t *testing.T) {
	tkn := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    2,
	}

	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  3,
		EndPos:    10,
	})

	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 3,
		EndLine:   3,
		StartPos:  11,
		EndPos:    20,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewTokenNodeListPosition(tkn, []node.Node{n1, n2})

	if pos.String() != `Pos{Line: 1-3 Pos: 0-20}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeNodeListPosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    8,
	})

	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  9,
		EndPos:    17,
	})

	n3 := node.NewIdentifier("test node")
	n3.SetPosition(&position.Position{
		StartLine: 3,
		EndLine:   3,
		StartPos:  18,
		EndPos:    26,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeNodeListPosition(n1, []node.Node{n2, n3})

	if pos.String() != `Pos{Line: 1-3 Pos: 0-26}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewNodeListNodePosition(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    8,
	})
	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  9,
		EndPos:    17,
	})
	n3 := node.NewIdentifier("test node")
	n3.SetPosition(&position.Position{
		StartLine: 3,
		EndLine:   3,
		StartPos:  18,
		EndPos:    26,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeListNodePosition([]node.Node{n1, n2}, n3)

	if pos.String() != `Pos{Line: 1-3 Pos: 0-26}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewOptionalListTokensPosition(t *testing.T) {
	builder := parser.PositionBuilder{}

	token1 := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}
	token2 := &scanner.Token{
		Value:     `foo`,
		StartLine: 2,
		EndLine:   2,
		StartPos:  4,
		EndPos:    6,
	}

	pos := builder.NewOptionalListTokensPosition(nil, token1, token2)

	if pos.String() != `Pos{Line: 1-2 Pos: 0-6}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNewOptionalListTokensPosition2(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    8,
	})
	n2 := node.NewIdentifier("test node")
	n2.SetPosition(&position.Position{
		StartLine: 2,
		EndLine:   2,
		StartPos:  9,
		EndPos:    17,
	})
	n3 := node.NewIdentifier("test node")
	n3.SetPosition(&position.Position{
		StartLine: 3,
		EndLine:   3,
		StartPos:  18,
		EndPos:    26,
	})

	builder := parser.PositionBuilder{}

	token1 := &scanner.Token{
		Value:     `foo`,
		StartLine: 4,
		EndLine:   4,
		StartPos:  27,
		EndPos:    29,
	}
	token2 := &scanner.Token{
		Value:     `foo`,
		StartLine: 5,
		EndLine:   5,
		StartPos:  30,
		EndPos:    32,
	}

	pos := builder.NewOptionalListTokensPosition([]node.Node{n2, n3}, token1, token2)

	if pos.String() != `Pos{Line: 2-5 Pos: 9-32}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNilNodePos(t *testing.T) {
	builder := parser.PositionBuilder{}

	pos := builder.NewNodesPosition(nil, nil)

	if pos.String() != `Pos{Line: -1--1 Pos: -1--1}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNilNodeListPos(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    8,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeNodeListPosition(n1, nil)

	if pos.String() != `Pos{Line: 1--1 Pos: 0--1}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestNilNodeListTokenPos(t *testing.T) {
	token := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeListTokenPosition(nil, token)

	if pos.String() != `Pos{Line: -1-1 Pos: -1-3}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestEmptyNodeListPos(t *testing.T) {
	n1 := node.NewIdentifier("test node")
	n1.SetPosition(&position.Position{
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    8,
	})

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeNodeListPosition(n1, []node.Node{})

	if pos.String() != `Pos{Line: 1--1 Pos: 0--1}` {
		t.Errorf("token value is not equal\n")
	}
}

func TestEmptyNodeListTokenPos(t *testing.T) {
	token := &scanner.Token{
		Value:     `foo`,
		StartLine: 1,
		EndLine:   1,
		StartPos:  0,
		EndPos:    3,
	}

	builder := parser.PositionBuilder{}

	pos := builder.NewNodeListTokenPosition([]node.Node{}, token)

	if pos.String() != `Pos{Line: -1-1 Pos: -1-3}` {
		t.Errorf("token value is not equal\n")
	}
}
