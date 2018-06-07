package expr_test

import (
	"reflect"
	"testing"

	"github.com/z7zmey/php-parser/node/stmt"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node/scalar"

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
		&expr.ArrayDimFetch{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			Dim:      &scalar.Lnumber{Value: "1"},
		},
		[]string{"Variable", "Dim"},
		map[string]interface{}{},
	},
	{
		&expr.ArrayItem{
			Key: &scalar.String{Value: "key"},
			Val: &scalar.Lnumber{Value: "1"},
		},
		[]string{"Key", "Val"},
		map[string]interface{}{},
	},
	{
		&expr.Array{
			Items: []node.Node{
				&expr.ArrayItem{
					Key: &scalar.String{Value: "key"},
					Val: &scalar.Lnumber{Value: "1"},
				},
			},
		},
		[]string{"Items"},
		map[string]interface{}{},
	},
	{
		&expr.BitwiseNot{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.BooleanNot{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.ClassConstFetch{
			Class:        &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			ConstantName: &node.Identifier{Value: "foo"},
		},
		[]string{"Class", "ConstantName"},
		map[string]interface{}{},
	},
	{
		&expr.Clone{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.ClosureUse{
			Uses: []node.Node{
				&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
		},
		[]string{"Uses"},
		map[string]interface{}{},
	},
	{
		&expr.Closure{
			ReturnsRef:    true,
			Static:        false,
			PhpDocComment: "",
			Params:        []node.Node{&node.Parameter{}},
			ClosureUse:    &expr.ClosureUse{},
			ReturnType:    &name.Name{},
			Stmts:         []node.Node{&stmt.Nop{}},
		},
		[]string{"Params", "ClosureUse", "ReturnType", "Stmts"},
		map[string]interface{}{"ReturnsRef": true, "Static": false, "PhpDocComment": ""},
	},
	{
		&expr.ConstFetch{
			Constant: &node.Identifier{Value: "foo"},
		},
		[]string{"Constant"},
		map[string]interface{}{},
	},
	{
		&expr.Empty{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.ErrorSuppress{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Eval{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Exit{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Die{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.FunctionCall{
			Function:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			ArgumentList: &node.ArgumentList{},
		},
		[]string{"Function", "ArgumentList"},
		map[string]interface{}{},
	},
	{
		&expr.IncludeOnce{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Include{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.InstanceOf{
			Expr:  &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			Class: &name.Name{},
		},
		[]string{"Expr", "Class"},
		map[string]interface{}{},
	},
	{
		&expr.Isset{
			Variables: []node.Node{
				&expr.Variable{VarName: &node.Identifier{Value: "a"}},
			},
		},
		[]string{"Variables"},
		map[string]interface{}{},
	},
	{
		&expr.List{
			Items: []node.Node{
				&expr.ArrayItem{},
			},
		},
		[]string{"Items"},
		map[string]interface{}{},
	},
	{
		&expr.MethodCall{
			Variable:     &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			Method:       &node.Identifier{Value: "foo"},
			ArgumentList: &node.ArgumentList{},
		},
		[]string{"Variable", "Method", "ArgumentList"},
		map[string]interface{}{},
	},
	{
		&expr.New{
			Class:        &name.Name{},
			ArgumentList: &node.ArgumentList{},
		},
		[]string{"Class", "ArgumentList"},
		map[string]interface{}{},
	},
	{
		&expr.PostDec{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Variable"},
		map[string]interface{}{},
	},
	{
		&expr.PostInc{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Variable"},
		map[string]interface{}{},
	},
	{
		&expr.PreDec{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Variable"},
		map[string]interface{}{},
	},
	{
		&expr.PreInc{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Variable"},
		map[string]interface{}{},
	},
	{
		&expr.Print{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.PropertyFetch{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			Property: &node.Identifier{Value: "foo"},
		},
		[]string{"Variable", "Property"},
		map[string]interface{}{},
	},
	{
		&expr.Reference{
			Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Variable"},
		map[string]interface{}{},
	},
	{
		&expr.RequireOnce{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Require{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.ShellExec{
			Parts: []node.Node{
				&scalar.EncapsedStringPart{},
			},
		},
		[]string{"Parts"},
		map[string]interface{}{},
	},
	{
		&expr.ShortArray{
			Items: []node.Node{
				&expr.ArrayItem{},
			},
		},
		[]string{"Items"},
		map[string]interface{}{},
	},
	{
		&expr.ShortList{
			Items: []node.Node{
				&expr.ArrayItem{},
			},
		},
		[]string{"Items"},
		map[string]interface{}{},
	},
	{
		&expr.StaticCall{
			Class:        &name.Name{},
			Call:         &node.Identifier{Value: "foo"},
			ArgumentList: &node.ArgumentList{},
		},
		[]string{"Class", "Call", "ArgumentList"},
		map[string]interface{}{},
	},
	{
		&expr.StaticPropertyFetch{
			Class:    &name.Name{},
			Property: &node.Identifier{Value: "foo"},
		},
		[]string{"Class", "Property"},
		map[string]interface{}{},
	},
	{
		&expr.Ternary{
			Condition: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			IfTrue:    &expr.Variable{VarName: &node.Identifier{Value: "b"}},
			IfFalse:   &expr.Variable{VarName: &node.Identifier{Value: "c"}},
		},
		[]string{"Condition", "IfTrue", "IfFalse"},
		map[string]interface{}{},
	},
	{
		&expr.UnaryMinus{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.UnaryPlus{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Variable{VarName: &node.Identifier{Value: "a"}},
		[]string{"VarName"},
		map[string]interface{}{},
	},
	{
		&expr.YieldFrom{
			Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
		},
		[]string{"Expr"},
		map[string]interface{}{},
	},
	{
		&expr.Yield{
			Key:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
			Value: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
		},
		[]string{"Key", "Value"},
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
