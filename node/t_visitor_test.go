package node_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/walker"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&node.Identifier{Value: "foo"},
		[]string{},
		map[string]interface{}{"Value": "foo"},
	},
	{
		&node.Nullable{Expr: &expr.Variable{}},
		[]string{"Expr"},
		nil,
	},
	{
		&node.Argument{Variadic: true, Expr: &expr.Variable{}},
		[]string{"Expr"},
		map[string]interface{}{"IsReference": false, "Variadic": true},
	},
	{
		&node.Parameter{
			ByRef:        false,
			Variadic:     true,
			VariableType: &node.Identifier{Value: "foo"},
			Variable:     &expr.Variable{},
			DefaultValue: &scalar.Lnumber{Value: "0"},
		},
		[]string{"VariableType", "Variable", "DefaultValue"},
		map[string]interface{}{"ByRef": false, "Variadic": true},
	},
	{
		&node.ArgumentList{
			Arguments: []node.Node{
				&node.Argument{},
			},
		},
		[]string{"Arguments"},
		nil,
	},
	{
		&node.Root{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
		nil,
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

func TestNameVisitorDisableChildren(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{false, []string{}}
		tt.node.Walk(v)

		expected := []string{}
		actual := v.visitedKeys

		assert.DeepEqual(t, expected, actual)
	}
}

func TestNameVisitor(t *testing.T) {
	for _, tt := range nodesToTest {
		v := &visitorMock{true, []string{}}
		tt.node.Walk(v)

		expected := tt.expectedVisitedKeys
		actual := v.visitedKeys

		assert.DeepEqual(t, expected, actual)
	}
}

// test Attributes()

func TestNameAttributes(t *testing.T) {
	for _, tt := range nodesToTest {
		expected := tt.expectedAttributes
		actual := tt.node.Attributes()

		assert.DeepEqual(t, expected, actual)
	}
}
