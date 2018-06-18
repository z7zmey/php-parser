package node_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/kylelemons/godebug/pretty"
	"github.com/z7zmey/php-parser/node"
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
		map[string]interface{}{},
	},
	{
		&node.Root{
			Stmts: []node.Node{&stmt.Expression{}},
		},
		[]string{"Stmts"},
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

func TestNameVisitorDisableChildren(t *testing.T) {
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

func TestNameVisitor(t *testing.T) {
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
