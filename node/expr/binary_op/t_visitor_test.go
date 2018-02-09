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
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BitwiseOr{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BitwiseXor{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BooleanAnd{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.BooleanOr{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Coalesce{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Concat{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Div{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Equal{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.GreaterOrEqual{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Greater{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Identical{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.LogicalAnd{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.LogicalOr{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.LogicalXor{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Minus{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Mod{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Mul{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.NotEqual{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.NotIdentical{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Plus{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Pow{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.ShiftLeft{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.ShiftRight{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.SmallerOrEqual{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Smaller{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary_op.Spaceship{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
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
