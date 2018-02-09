package binary_op_test

import (
	"reflect"
	"testing"
	
	"github.com/kylelemons/godebug/pretty"
	
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/binary_op"
	"github.com/z7zmey/php-parser/walker"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&binary_op.BitwiseAnd{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BitwiseOr{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BitwiseXor{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BooleanAnd{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BooleanOr{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Coalesce{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Concat{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Div{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Equal{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.GreaterOrEqual{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Greater{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Identical{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.LogicalAnd{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.LogicalOr{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.LogicalXor{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Minus{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Mod{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Mul{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.NotEqual{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.NotIdentical{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Plus{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Pow{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.ShiftLeft{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.ShiftRight{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.SmallerOrEqual{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Smaller{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Spaceship{
			BinaryOp: binary_op.BinaryOp{
				Left: &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
			},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
}

type visitorMock struct {
	visitChildren bool
	visitedKeys   []string
}

func (v *visitorMock) EnterNode(n walker.Walker) bool { return v.visitChildren }
func (v *visitorMock) GetChildrenVisitor(key string) walker.Visitor {
	v.visitedKeys = append(v.visitedKeys, key)
	return &visitorMock{v.visitChildren, nil}
}
func (v *visitorMock) LeaveNode(n walker.Walker) {}

func TestVisitorDisableChildren(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{false, nil}
		tt.node.Walk(v)

		expected := []string{}
		actual := v.visitedKeys

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}

func TestVisitor(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{true, nil}
		tt.node.Walk(v)

		expected := tt.expectedVisitedKeys
		actual := v.visitedKeys

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}

// test Attributes()

func TestNameAttributes(t *testing.T) {
	for _, tt := range nodesToTest {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		diff := pretty.Compare(expected, actual)
		if diff != "" {
			t.Errorf("%s diff: (-expected +actual)\n%s", reflect.TypeOf(tt.node), diff)
		}
	}
}
