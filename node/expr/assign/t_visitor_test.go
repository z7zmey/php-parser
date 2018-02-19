package assign_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign"

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
		&assign.AssignRef{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Assign{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.BitwiseAnd{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.BitwiseOr{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.BitwiseXor{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Concat{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Div{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Minus{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Mod{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Mul{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Plus{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.Pow{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.ShiftLeft{
			Variable:   &expr.Variable{VarName: &node.Identifier{Value: "$a"}},
			Expression: &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
		},
		[]string{"Variable", "Expression"},
		map[string]interface{}{},
	},
	{
		&assign.ShiftRight{
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
