package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
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
	&expr.ArrayDimFetch{
		FreeFloating: expected,
	},
	&expr.ArrayItem{
		FreeFloating: expected,
	},
	&expr.Array{
		FreeFloating: expected,
	},
	&expr.BitwiseNot{
		FreeFloating: expected,
	},
	&expr.BooleanNot{
		FreeFloating: expected,
	},
	&expr.ClassConstFetch{
		FreeFloating: expected,
	},
	&expr.Clone{
		FreeFloating: expected,
	},
	&expr.ClosureUse{
		FreeFloating: expected,
	},
	&expr.Closure{
		FreeFloating: expected,
	},
	&expr.ConstFetch{
		FreeFloating: expected,
	},
	&expr.Empty{
		FreeFloating: expected,
	},
	&expr.ErrorSuppress{
		FreeFloating: expected,
	},
	&expr.Eval{
		FreeFloating: expected,
	},
	&expr.Exit{
		FreeFloating: expected,
	},
	&expr.FunctionCall{
		FreeFloating: expected,
	},
	&expr.IncludeOnce{
		FreeFloating: expected,
	},
	&expr.Include{
		FreeFloating: expected,
	},
	&expr.InstanceOf{
		FreeFloating: expected,
	},
	&expr.Isset{
		FreeFloating: expected,
	},
	&expr.List{
		FreeFloating: expected,
	},
	&expr.MethodCall{
		FreeFloating: expected,
	},
	&expr.New{
		FreeFloating: expected,
	},
	&expr.PostDec{
		FreeFloating: expected,
	},
	&expr.PostInc{
		FreeFloating: expected,
	},
	&expr.PreDec{
		FreeFloating: expected,
	},
	&expr.PreInc{
		FreeFloating: expected,
	},
	&expr.Print{
		FreeFloating: expected,
	},
	&expr.PropertyFetch{
		FreeFloating: expected,
	},
	&expr.Reference{
		FreeFloating: expected,
	},
	&expr.RequireOnce{
		FreeFloating: expected,
	},
	&expr.Require{
		FreeFloating: expected,
	},
	&expr.ShellExec{
		FreeFloating: expected,
	},
	&expr.ShortArray{
		FreeFloating: expected,
	},
	&expr.ShortList{
		FreeFloating: expected,
	},
	&expr.StaticCall{
		FreeFloating: expected,
	},
	&expr.StaticPropertyFetch{
		FreeFloating: expected,
	},
	&expr.Ternary{
		FreeFloating: expected,
	},
	&expr.UnaryMinus{
		FreeFloating: expected,
	},
	&expr.UnaryPlus{
		FreeFloating: expected,
	},
	&expr.Variable{
		FreeFloating: expected,
	},
	&expr.YieldFrom{
		FreeFloating: expected,
	},
	&expr.Yield{
		FreeFloating: expected,
	},
}

func TestMeta(t *testing.T) {
	for _, n := range nodes {
		actual := *n.GetFreeFloating()
		assert.DeepEqual(t, expected, actual)
	}
}
