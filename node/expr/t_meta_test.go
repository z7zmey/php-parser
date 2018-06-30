package expr_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&expr.ArrayDimFetch{},
	&expr.ArrayItem{},
	&expr.Array{},
	&expr.BitwiseNot{},
	&expr.BooleanNot{},
	&expr.ClassConstFetch{},
	&expr.Clone{},
	&expr.ClosureUse{},
	&expr.Closure{},
	&expr.ConstFetch{},
	&expr.Die{},
	&expr.Empty{},
	&expr.ErrorSuppress{},
	&expr.Eval{},
	&expr.Exit{},
	&expr.FunctionCall{},
	&expr.IncludeOnce{},
	&expr.Include{},
	&expr.InstanceOf{},
	&expr.Isset{},
	&expr.List{},
	&expr.MethodCall{},
	&expr.New{},
	&expr.PostDec{},
	&expr.PostInc{},
	&expr.PreDec{},
	&expr.PreInc{},
	&expr.Print{},
	&expr.PropertyFetch{},
	&expr.Reference{},
	&expr.RequireOnce{},
	&expr.Require{},
	&expr.ShellExec{},
	&expr.ShortArray{},
	&expr.ShortList{},
	&expr.StaticCall{},
	&expr.StaticPropertyFetch{},
	&expr.Ternary{},
	&expr.UnaryMinus{},
	&expr.UnaryPlus{},
	&expr.Variable{},
	&expr.YieldFrom{},
	&expr.Yield{},
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
