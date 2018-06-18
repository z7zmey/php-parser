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
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BitwiseOr{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BitwiseXor{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BooleanAnd{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.BooleanOr{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Coalesce{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Concat{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Div{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Equal{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.GreaterOrEqual{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Greater{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Identical{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.LogicalAnd{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.LogicalOr{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.LogicalXor{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Minus{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Mod{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Mul{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.NotEqual{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.NotIdentical{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Plus{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Pow{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.ShiftLeft{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.ShiftRight{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.SmallerOrEqual{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Smaller{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
		},
		[]string{"Left", "Right"},
		map[string]interface{}{},
	},
	{
		&binary.Spaceship{
			Left:  &expr.Variable{},
			Right: &expr.Variable{},
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
func (v *visitorMock) LeaveNode(n walker.Walkable)      {}
func (v *visitorMock) EnterChildNode(key string, w walker.Walkable) {
	v.visitedKeys = append(v.visitedKeys, key)
}
func (v *visitorMock) LeaveChildNode(key string, w walker.Walkable) {}
func (v *visitorMock) EnterChildList(key string, w walker.Walkable) {
	v.visitedKeys = append(v.visitedKeys, key)
}
func (v *visitorMock) LeaveChildList(key string, w walker.Walkable) {}

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
