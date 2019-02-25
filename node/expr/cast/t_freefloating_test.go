package cast_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr/cast"
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
	&cast.Array{
		FreeFloating: expected,
	},
	&cast.Bool{
		FreeFloating: expected,
	},
	&cast.Double{
		FreeFloating: expected,
	},
	&cast.Int{
		FreeFloating: expected,
	},
	&cast.Object{
		FreeFloating: expected,
	},
	&cast.String{
		FreeFloating: expected,
	},
	&cast.Unset{
		FreeFloating: expected,
	},
}

func TestMeta(t *testing.T) {
	for _, n := range nodes {
		actual := *n.GetFreeFloating()
		assert.DeepEqual(t, expected, actual)
	}
}
