package cast_test

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/expr/cast"
	"github.com/z7zmey/php-parser/walker"
)

var nodesToTest = []struct {
	node                node.Node // node
	expectedVisitedKeys []string  // visited keys
	expectedAttributes  map[string]interface{}
}{
	{
		&cast.CastArray{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&cast.CastBool{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&cast.CastDouble{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&cast.CastInt{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&cast.CastObject{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&cast.CastString{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&cast.CastUnset{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
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
