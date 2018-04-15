package parser_test

import (
	"testing"

	"github.com/z7zmey/php-parser/position"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/parser"
)

func TestPositions(t *testing.T) {
	n := node.NewIdentifier("test")

	expected := position.NewPosition(0, 0, 0, 0)

	positions := parser.Positions{}
	positions.AddPosition(n, expected)

	actual := positions[n]

	if actual != expected {
		t.Errorf("expected and actual are not equal\n")
	}
}
