package binary_test

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/binary"
	"github.com/z7zmey/php-parser/walker"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&binary.BitwiseAnd{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BitwiseOr{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BitwiseXor{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BooleanAnd{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BooleanOr{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Coalesce{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Concat{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Div{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Equal{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.GreaterOrEqual{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Greater{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Identical{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.LogicalAnd{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.LogicalOr{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.LogicalXor{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Minus{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Mod{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Mul{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.NotEqual{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.NotIdentical{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Plus{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Pow{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.ShiftLeft{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.ShiftRight{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.SmallerOrEqual{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Smaller{
			Left:  &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Right: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Spaceship{
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

func (v *visitorMock) EnterNode(n walker.Walkable) bool { return v.visitChildren }
func (v *visitorMock) GetChildrenVisitor(key string) walker.Visitor {
	v.visitedKeys = append(v.visitedKeys, key)
	return &visitorMock{v.visitChildren, nil}
}
func (v *visitorMock) LeaveNode(n walker.Walkable) {}

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
