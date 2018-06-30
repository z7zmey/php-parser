package binary_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/expr/binary"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&binary.BitwiseAnd{},
	&binary.BitwiseOr{},
	&binary.BitwiseXor{},
	&binary.BooleanAnd{},
	&binary.BooleanOr{},
	&binary.Coalesce{},
	&binary.Concat{},
	&binary.Div{},
	&binary.Equal{},
	&binary.GreaterOrEqual{},
	&binary.Greater{},
	&binary.Identical{},
	&binary.LogicalAnd{},
	&binary.LogicalOr{},
	&binary.LogicalXor{},
	&binary.Minus{},
	&binary.Mod{},
	&binary.Mul{},
	&binary.NotEqual{},
	&binary.NotIdentical{},
	&binary.Plus{},
	&binary.Pow{},
	&binary.ShiftLeft{},
	&binary.ShiftRight{},
	&binary.SmallerOrEqual{},
	&binary.Smaller{},
	&binary.Spaceship{},
}

func TestMeta(t *testing.T) {
	expected := []meta.Meta{
		meta.NewComment("//comment\n", nil),
		meta.NewWhiteSpace("    ", nil),
	}
	for _, n := range nodes {
		n.AddMeta(expected)
		actual := n.GetMeta()
		assertEqual(t, expected, actual)
	}
}
