package binary_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr/binary"
)

var expected freefloating.Collection = freefloating.Collection{
	freefloating.Start: []freefloating.String{
		{
			StringType: freefloating.WhiteSpaceType,
			Value:      "    ",
			Position:   nil,
		},
		{
			StringType: freefloating.CommentType,
			Value:      "//comment\n",
			Position:   nil,
		},
	},
}

var nodes = []node.Node{
	&binary.BitwiseAnd{
		FreeFloating: expected,
	},
	&binary.BitwiseOr{
		FreeFloating: expected,
	},
	&binary.BitwiseXor{
		FreeFloating: expected,
	},
	&binary.BooleanAnd{
		FreeFloating: expected,
	},
	&binary.BooleanOr{
		FreeFloating: expected,
	},
	&binary.Coalesce{
		FreeFloating: expected,
	},
	&binary.Concat{
		FreeFloating: expected,
	},
	&binary.Div{
		FreeFloating: expected,
	},
	&binary.Equal{
		FreeFloating: expected,
	},
	&binary.GreaterOrEqual{
		FreeFloating: expected,
	},
	&binary.Greater{
		FreeFloating: expected,
	},
	&binary.Identical{
		FreeFloating: expected,
	},
	&binary.LogicalAnd{
		FreeFloating: expected,
	},
	&binary.LogicalOr{
		FreeFloating: expected,
	},
	&binary.LogicalXor{
		FreeFloating: expected,
	},
	&binary.Minus{
		FreeFloating: expected,
	},
	&binary.Mod{
		FreeFloating: expected,
	},
	&binary.Mul{
		FreeFloating: expected,
	},
	&binary.NotEqual{
		FreeFloating: expected,
	},
	&binary.NotIdentical{
		FreeFloating: expected,
	},
	&binary.Plus{
		FreeFloating: expected,
	},
	&binary.Pow{
		FreeFloating: expected,
	},
	&binary.ShiftLeft{
		FreeFloating: expected,
	},
	&binary.ShiftRight{
		FreeFloating: expected,
	},
	&binary.SmallerOrEqual{
		FreeFloating: expected,
	},
	&binary.Smaller{
		FreeFloating: expected,
	},
	&binary.Spaceship{
		FreeFloating: expected,
	},
}

func TestMeta(t *testing.T) {
	for _, n := range nodes {
		actual := *n.GetFreeFloating()
		assert.DeepEqual(t, expected, actual)
	}
}
