package assign_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr/assign"
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
	&assign.Reference{
		FreeFloating: expected,
	},
	&assign.Assign{
		FreeFloating: expected,
	},
	&assign.BitwiseAnd{
		FreeFloating: expected,
	},
	&assign.BitwiseOr{
		FreeFloating: expected,
	},
	&assign.BitwiseXor{
		FreeFloating: expected,
	},
	&assign.Concat{
		FreeFloating: expected,
	},
	&assign.Div{
		FreeFloating: expected,
	},
	&assign.Minus{
		FreeFloating: expected,
	},
	&assign.Mod{
		FreeFloating: expected,
	},
	&assign.Mul{
		FreeFloating: expected,
	},
	&assign.Plus{
		FreeFloating: expected,
	},
	&assign.Pow{
		FreeFloating: expected,
	},
	&assign.ShiftLeft{
		FreeFloating: expected,
	},
	&assign.ShiftRight{
		FreeFloating: expected,
	},
	&assign.ShiftRight{
		FreeFloating: expected,
	},
}

func TestMeta(t *testing.T) {
	for _, n := range nodes {
		actual := *n.GetFreeFloating()
		assert.DeepEqual(t, expected, actual)
	}
}
