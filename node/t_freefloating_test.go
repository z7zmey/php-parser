package node_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
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
	&node.ArgumentList{
		FreeFloating: expected,
	},
	&node.Argument{
		FreeFloating: expected,
	},
	&node.Identifier{
		FreeFloating: expected,
	},
	&node.Nullable{
		FreeFloating: expected,
	},
	&node.Parameter{
		FreeFloating: expected,
	},
	&node.Root{
		FreeFloating: expected,
	},
}

func TestMeta(t *testing.T) {
	for _, n := range nodes {
		actual := *n.GetFreeFloating()
		assert.DeepEqual(t, expected, actual)
	}
}
