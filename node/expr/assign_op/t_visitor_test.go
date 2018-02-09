package assign_op_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign_op"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&assign_op.AssignRef{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Assign{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.BitwiseAnd{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.BitwiseOr{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.BitwiseXor{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Concat{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Div{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Minus{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Mod{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Mul{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Plus{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.Pow{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.ShiftLeft{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign_op.ShiftRight{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
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
